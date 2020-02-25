package script

import (
	"reflect"
)

//List type.
type List struct {
	Type
}

//Make makes a list.
func (*List) Make(q Ctx, collection Collection, sizes ...int) {
	var T = reflect.TypeOf(collection).Elem()
	var V = reflect.ValueOf(collection).Elem()

	var L, ok = T.FieldByName("L")
	if !ok {
		panic("array type must have literal `L` field")
	}
	var ElementType = L.Type.Elem()

	var ZeroType = reflect.SliceOf(
		GoTypeOf(reflect.Zero(ElementType).Interface().(Value)))

	var Zero = reflect.Zero(ZeroType).Interface()

	if len(sizes) > 0 {
		Zero = reflect.MakeSlice(ZeroType, sizes[0], sizes[0]).Interface()
	}

	//Create a runtime representation of the array.
	V.FieldByName("List").Set(reflect.ValueOf(List{
		NewType(q, func() interface{} {
			return Zero
		}),
	}))

	//Create Mutate method.
	if Mutate, ok := T.FieldByName("Mutate"); ok {
		V.FieldByName("Mutate").Set(reflect.MakeFunc(Mutate.Type,
			func(args []reflect.Value) []reflect.Value {
				q.Mutate(collection, args[0].Interface().(Int), args[1].Interface().(Value))
				return nil
			}))
	}

	//Create Index method.
	if Index, ok := T.FieldByName("Index"); ok {
		V.FieldByName("Index").Set(reflect.MakeFunc(Index.Type,
			func(args []reflect.Value) (returns []reflect.Value) {
				var result = reflect.New(ElementType).Elem()
				result.FieldByName("Type").Set(reflect.ValueOf(Type{
					Ctx:     q,
					Runtime: q.Index(collection, args[0].Interface().(Int)),
				}))
				returns = append(returns, result)
				return
			}))
	}
}
