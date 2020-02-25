package script

import "reflect"

//Collection is a collection value.
type Collection interface {
	Value
	Make(Ctx, Collection, ...int)
}

//Make makes a structured type.
func (q Ctx) Make(collection Collection, sizes ...int) {
	collection.Make(q, collection, sizes...)

	if q.defining {
		var variable = q.getVar()

		reflect.ValueOf(collection).Elem().FieldByName("Type").Set(reflect.ValueOf(Type{
			Ctx:     q,
			Runtime: q.Language.DefineVariable(variable, collection),
		}))
	}
}
