package script

import "reflect"
import "bytes"
import "github.com/qlova/script/language"
import "github.com/qlova/script/go"

type Function struct {
	language.Function
	EmbeddedScript

	Literal interface{}
	Names []Go.String
}

func (q Script) External(name string) Function {
	return q.wrap(q.lang.External(convert(name).(language.String))).(Function)
}

func (q Script) Function(i interface{}) Function {
	return Function{Literal: i, EmbeddedScript: EmbeddedScript{q:q}}
}

func (q Script) Return(T Type) {
	q.indent()
	q.write(q.lang.Return(convert(T)))

	q.returns[len(q.returns)-1] = T
}

func (f Function) Call(arguments ...Type) Type {
	
	var converted = make([]language.Type, len(arguments))
	for i := range arguments {
		converted[i] = convert(arguments[i])
	}
	
	if f.Returns() != nil {
		return f.q.wrap(f.q.lang.Call(convert(f).(language.Function), converted))
	}
	
	f.q.indent()
	f.q.write(f.q.lang.Run(convert(f).(language.Function), converted))
	return nil
}

func (f Function) Run(arguments ...Type) {
	
	var converted = make([]language.Type, len(arguments))
	for i := range arguments {
		converted[i] = convert(arguments[i])
	}
	
	f.q.indent()
	f.q.write(f.q.lang.Run(convert(f).(language.Function), converted))
	return 
}


func (function *Function) NameArguments(names []Go.String) {
	function.Names = names
}

//Like Define but instead of being defined now, it is promoted to the Top-level of the script.
func (function *Function) Promote(name Go.String) {
	if function.Literal == nil {
		panic("Function.Promote(): Cannot promote dynamic functions!")
	}
	
	var q = function.q
	
	var old = q.body
	var old_depth = q.depth
	q.body = bytes.Buffer{}
	q.depth = 1

	var statement language.Statement

	if f, ok := function.Literal.(func()); ok {

		function.Function, statement = q.lang.Function(name, function.Names, nil, nil)
		f()

	} else {
		
		//Let the function know what arguments it has.
		var ReflectedFunctionType = reflect.TypeOf(function.Literal)
		var Arguments = make([]language.Type, ReflectedFunctionType.NumIn())
		for i := 0; i < ReflectedFunctionType.NumIn(); i++ {
			var Argument = ReflectedFunctionType.In(i)
			
			//Check for fancy Go stuff like slices, subtypes etc.
			
			switch Argument {
				case reflect.TypeOf(Function{}):
					Arguments[i] = language.FunctionType{}
				
				case reflect.TypeOf(Number{}):
					Arguments[i] = language.NumberType{}
				
				default:
					panic("Function.Promote(): Invalid argument: "+Argument.Name())
			}
		}
		
		function.Function, statement = q.lang.Function(name, function.Names, Arguments, nil)
		
		var Args = function.Arguments()
		var Reflected = make([]reflect.Value, len(Args))
		for i := range Arguments {
			Reflected[i] = reflect.ValueOf(q.wrap(Args[i]))
		}
		
		q.returns = append(q.returns, nil)
		reflect.ValueOf(function.Literal).Call(Reflected)
		
		if q.returns[len(q.returns)-1] != nil {
			//Damn, we have returns, how we gonna flip this around.
			//Hmmm... Time to invent a new protocol >:)
			function.Function, statement = q.lang.UpdateFunction(function.Function, name, function.Names, Arguments, convert(q.returns[len(q.returns)-1]))
			q.returns = q.returns[:len(q.returns)-1]
		}
	}
	
	q.neck.WriteString(Go.String(statement))
	q.neck.Write(q.body.Bytes())
	q.neck.WriteString(Go.String(q.lang.EndFunction()))

	q.depth = 0
	
	q.body = old
	q.depth = old_depth

	function.Literal = nil
}
