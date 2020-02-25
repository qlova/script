package script

import "reflect"

//Table type.
type Table struct {
	Type
}

//Make makes a table.
func (*Table) Make(q Ctx, collection Collection, sizes ...int) {
	var T = reflect.TypeOf(collection).Elem()
	var V = reflect.ValueOf(collection).Elem()

	var L, ok = T.FieldByName("L")
	if !ok {
		panic("table type must have literal `L` field")
	}
	var ElementType = L.Type.Elem()

	var ZeroType = reflect.MapOf(reflect.TypeOf(""),
		GoTypeOf(reflect.Zero(ElementType).Interface().(Value)))

	var Zero = reflect.MakeMap(ZeroType).Interface()

	//Create a runtime representation of the array.
	V.FieldByName("Table").Set(reflect.ValueOf(Table{
		NewType(q, func() interface{} {
			return Zero
		}),
	}))

	//Create Mutate method.
	if Mutate, ok := T.FieldByName("Insert"); ok {
		V.FieldByName("Insert").Set(reflect.MakeFunc(Mutate.Type,
			func(args []reflect.Value) []reflect.Value {
				q.Insert(collection, args[0].Interface().(String), args[1].Interface().(Value))
				return nil
			}))
	}

	//Create Index method.
	if Index, ok := T.FieldByName("Lookup"); ok {
		V.FieldByName("Lookup").Set(reflect.MakeFunc(Index.Type,
			func(args []reflect.Value) (returns []reflect.Value) {
				var result = reflect.New(ElementType).Elem()
				result.FieldByName("Type").Set(reflect.ValueOf(Type{
					Ctx:     q,
					Runtime: q.Lookup(collection, args[0].Interface().(String)),
				}))
				returns = append(returns, result)
				return
			}))
	}
}
