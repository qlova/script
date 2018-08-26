package script

import "errors"

type List Array

func (n List) Name() string {
	return "List"
}

func (n List) Equals(t interface{}) bool {
	if _, ok := t.(List); ok {
		return true
	}
	return false
}

func (n List) String() string {
	return n.data
}

func (n List) Subtype() Type {
	return n.subtype
}

func (q *Script) LiteralList(subtype Type, elements ...Type) List {
	
	//TODO type validation!
	
	return List{
		subtype: subtype, 
		data: q.Language.LiteralList(subtype, elements...),
	}
}

func (q *Script) NewList(subtype Type, name string, value List) {
	q.IndentBody()
	q.Language.NewList(subtype, q, name, value)
}

func (q *Script) LengthList(list List) Number {
	return Number(q.Language.LengthList( list.String() ))
}

func (q *Script) IndexList(list List, index Number) Type {
	return q.Raw(list.subtype, q.Language.IndexList(list.String(), index.String()))
}

func (q *Script) IndexListRaw(list string, index string) string {
	return q.Language.IndexList(list, index)
}

func (q *Script) SetList(subtype Type, name string, value List) {
	q.IndentBody()
	q.Language.SetList(subtype, q, name, value)
}

func (q *Script) ToList(subtype Type, T Type) (List, error) {
	if list, ok := T.(List); ok {
		if list.subtype.Equals(subtype) {
			return List{subtype: subtype, data: T.String()}, nil
		}
	}
	
	switch T.(type) {
		
		case Number:
			if (subtype.Equals(Number(""))) {
				return List{
					subtype: Number(""), 
					data:q.Language.NumberToList(subtype, string(T.(Number))),
				}, nil
			}

	}
	return List{}, errors.New("Cannot convert "+T.Name()+" to List!")
}
