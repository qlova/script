package script

import "errors"

type Array struct {
	data string
	subtype Type
}

func (n Array) Name() string {
	return "Array"
}

func (n Array) Equals(t interface{}) bool {
	if _, ok := t.(Array); ok {
		return true
	}
	return false
}

func (n Array) String() string {
	return n.data
}

func (n Array) Subtype() Type {
	return n.subtype
}

func (q *Script) LiteralArray(subtype Type, elements ...Type) Array {
	
	//TODO type validation!
	
	return Array{
		subtype: subtype, 
		data: q.Language.LiteralArray(subtype, elements...),
	}
}

func (q *Script) NewArray(subtype Type, name string, value Array) {
	q.IndentBody()
	q.Language.NewArray(subtype, q, name, value)
}

func (q *Script) LengthArray(array Array) Number {
	return Number(q.Language.LengthArray( array.String() ))
}

func (q *Script) IndexArray(array Array, index Number) Type {
	return q.Raw(array.subtype, q.Language.IndexArray(array.String(), index.String()))
}

func (q *Script) IndexArrayRaw(array string, index string) string {
	return q.Language.IndexArray(array, index)
}

func (q *Script) SetArray(subtype Type, name string, value Array) {
	q.IndentBody()
	q.Language.SetArray(subtype, q, name, value)
}

func (q *Script) ToArray(subtype Type, T Type) (Array, error) {
	if array, ok := T.(Array); ok {
		if array.subtype.Equals(subtype) {
			return Array{subtype: subtype, data: T.String()}, nil
		}
	}
	
	switch T.(type) {
		
		case Number:
			if (subtype.Equals(Number(""))) {
				return Array{
					subtype: Number(""), 
					data:q.Language.NumberToArray(subtype, string(T.(Number))),
				}, nil
			}

	}
	return Array{}, errors.New("Cannot convert "+T.Name()+" to Array!")
}
