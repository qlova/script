package script

import (
	"reflect"
)

//New initialises an empty script.Type derived value.
func (q Ctx) New(structures ...EmptyType) {
	for _, structure := range structures {
		var T = reflect.TypeOf(structure).Elem()
		var V = reflect.ValueOf(structure).Elem()

		var fields = q.getFields(T)

		//Create a runtime representation of the structure.
		if structure.T().Runtime == nil {
			V.FieldByName("Type").Set(reflect.ValueOf(
				NewType(q, func() interface{} {
					var structure = make(map[string]interface{})
					for _, field := range fields {
						structure[field.Name] = GoValueOf(field.Type)
					}
					return structure
				}),
			))
		}

		//Fill the fields.
		for i := 0; i < T.NumField(); i++ {
			var field = T.Field(i)
			if field.Name == "Type" {
				continue
			}
			if field.Name == "Set" {
				continue
			}

			//Initialise the fields.
			if _, ok := reflect.Zero(field.Type).Interface().(Value); ok {
				V.Field(i).FieldByName("Type").Set(reflect.ValueOf(
					Type{
						Ctx:     q,
						Runtime: q.Language.Field(structure, field.Name),
					},
				))
			}

			//Initialise the methods.
			if field.Type.Kind() == reflect.Func {
				var NumberOfArguments = field.Type.NumIn()
				var MethodType = field.Type
				var MethodName = field.Name
				V.Field(i).Set(reflect.MakeFunc(MethodType,
					func(ActualArgs []reflect.Value) (returns []reflect.Value) {

						var Args = make([]Value, NumberOfArguments)

						for i, arg := range ActualArgs {
							Args[i] = arg.Interface().(Value)
						}

						if MethodType.NumOut() == 0 {
							q.Language.RunMethod(structure, MethodName, Args)
							return
						}

						var Pointer = reflect.New(MethodType.Out(0))
						var NewReturns = Pointer.Elem()
						NewReturns.FieldByName("Type").Set(reflect.ValueOf(
							Type{
								Ctx:     q,
								Runtime: q.Language.CallMethod(structure, MethodName, Args),
							},
						))

						q.New(Pointer.Interface().(EmptyType))

						return []reflect.Value{
							NewReturns,
						}
					}))
			}
		}

		//Create Set method.
		if Set, ok := T.FieldByName("Set"); ok {
			V.FieldByName("Set").Set(reflect.MakeFunc(Set.Type,
				func(args []reflect.Value) []reflect.Value {
					q.Set(structure, args[1].Interface().(Value))
					return nil
				}))
		}

		if q.defining {
			var variable = q.getVar()

			reflect.ValueOf(structure).Elem().FieldByName("Type").Set(reflect.ValueOf(Type{
				Ctx:     q,
				Runtime: q.Language.DefineVariable(variable, structure),
			}))
		}
	}
}

type Field struct {
	Name string
	Type Value
}

type Struct struct {
	Name   string
	Fields []Field

	Reciever string
	Methods  []Function
}

func (q Ctx) getFields(T reflect.Type) (fields []Field) {
	for i := 0; i < T.NumField(); i++ {
		var field = T.Field(i)
		if field.Name == "Type" {
			continue
		}
		if field.Name == "Set" {
			continue
		}
		if zero, ok := reflect.Zero(field.Type).Interface().(Value); ok {
			fields = append(fields, Field{
				Name: field.Name,
				Type: zero,
			})
		}
	}
	return
}

func (q Ctx) getMethods(metatype Value) (methods []Function) {
	var T = reflect.TypeOf(metatype).Elem()

	for i := 0; i < T.NumField(); i++ {
		var field = T.Field(i)
		if field.Name == "Type" {
			continue
		}
		if field.Name == "Set" {
			continue
		}
		if field.Type.Kind() == reflect.Func {
			var args []Argument
			var returns Value

			for i := 0; i < field.Type.NumIn(); i++ {
				args = append(args, Argument{
					Name:  q.ID("arg_"),
					Value: reflect.Zero(field.Type.In(i)).Interface().(Value),
				})
			}

			if field.Type.NumOut() > 0 {
				returns = reflect.Zero(field.Type.Out(0)).Interface().(Value)
			}

			var strings = make([]string, len(args))
			for i := range args {
				strings[i] = args[i].Name
			}
			var block, _ = q.dummyFunc(reflect.ValueOf(metatype).Elem().Field(i).Interface(), strings...)

			methods = append(methods, Function{
				Name:    field.Name,
				Args:    args,
				Returns: returns,
				Block:   block,
			})
		}
	}
	return
}

//DefineStruct defines a type with the included methods given a metatype.
//Takes an optional reciever name.
func (q Ctx) defineStruct(metatype EmptyType, reciever ...string) {
	var T = reflect.TypeOf(metatype).Elem()

	var variable string
	if len(reciever) > 0 {
		variable = reciever[0]
	} else {
		variable = q.ID("recv")
	}

	var fields = q.getFields(T)
	var methods = q.getMethods(metatype)

	//Remember to initialise the reciever.
	reflect.ValueOf(metatype).Elem().FieldByName("Type").Set(reflect.ValueOf(Type{
		Ctx:     q,
		Runtime: q.Language.Argument(string(variable), -1),
	}))
	q.New(metatype)

	q.Language.DefineStruct(Struct{
		Name:   T.Name(),
		Fields: fields,

		Reciever: variable,
		Methods:  methods,
	})
}

//DefineStruct defines a type with the included methods given a metatype.
//Takes an optional reciever name.
func (q Ctx) DefineType(metatype EmptyType, reciever ...string) {
	//var T = reflect.TypeOf(metatype)
	q.defineStruct(metatype, reciever...)
}
