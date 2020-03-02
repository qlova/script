package script

import (
	"fmt"
	"reflect"
)

type Function struct {
	Name    string
	Args    []Argument
	Returns Value
	Block   func()
}

type Argument struct {
	Name string
	Value
}

type Variable string

func (q Ctx) dummyFunc(function interface{}, strings ...string) (func(), []reflect.Value) {
	var FunctionType = reflect.TypeOf(function)
	var Original = reflect.ValueOf(function)

	var NumberOfArguments = FunctionType.NumIn()

	//Define the function body by passing dummy arguments with their expected locations.
	var DummyArgs = make([]reflect.Value, NumberOfArguments)
	for i := 0; i < NumberOfArguments; i++ {
		var ArgumentName = Variable(fmt.Sprintf("arg_%v", i))
		if len(strings)-1 > i {
			ArgumentName = Variable(strings[i+1])
		}

		var Pointer = reflect.New(FunctionType.In(0))

		DummyArgs[i] = Pointer.Elem()
		q.New(Pointer.Interface().(EmptyType))
		DummyArgs[i].FieldByName("Type").Set(reflect.ValueOf(Type{
			Ctx:     q,
			Runtime: q.Language.Argument(string(ArgumentName), i),
		}))
	}

	return func() {
		Original.Call(DummyArgs)
	}, DummyArgs
}

//DefineFunc defines a new function from a *func(...) ...
func (q Ctx) DefineFunc(function interface{}, strings ...string) {
	var FunctionType = reflect.TypeOf(function).Elem()
	var FunctionValue = reflect.ValueOf(function).Elem()

	var NumberOfArguments = FunctionType.NumIn()

	var FunctionName = q.ID("function_")
	if len(strings) > 0 {
		FunctionName = strings[0]
	}

	var f, DummyArgs = q.dummyFunc(reflect.ValueOf(function).Elem().Interface(), strings...)

	var Args = make([]Argument, NumberOfArguments)

	for i, arg := range DummyArgs {
		Args[i].Name = fmt.Sprintf("arg_%v", i)
		if len(strings)-1 > i {
			Args[i].Name = strings[i+1]
		}
		Args[i].Value = arg.Interface().(Value)
	}

	var Returns Value = nil
	if FunctionType.NumOut() > 0 {
		Returns = reflect.New(FunctionType.Out(0)).Elem().Interface().(Value)
	}

	q.Language.DefineFunction(Function{
		Name:    FunctionName,
		Args:    Args,
		Returns: Returns,
		Block:   f,
	})

	FunctionValue.Set(reflect.MakeFunc(FunctionType,
		func(ActualArgs []reflect.Value) (returns []reflect.Value) {

			var Args = make([]Value, NumberOfArguments)

			for i, arg := range ActualArgs {
				Args[i] = arg.Interface().(Value)
			}

			if Returns == nil {
				q.Language.RunFunction(FunctionName, Args)
				return
			}

			var NewReturns = reflect.New(FunctionType.Out(0)).Elem()
			NewReturns.FieldByName("Type").Set(reflect.ValueOf(
				Type{
					Ctx:     q,
					Runtime: q.Language.CallFunction(FunctionName, Args),
				},
			))

			return []reflect.Value{
				NewReturns,
			}
		}))
}

func (q Ctx) Return(v ...Value) {
	if len(v) == 0 {
		q.Language.Return(nil)
		return
	}
	q.Language.Return(v[0])
}
