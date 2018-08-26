package Go

import qlova "github.com/qlova/script"

func (l *language) NewList(T qlova.Type,q *qlova.Script, name string, value ...qlova.List) {
	New("[]"+l.GoTypeOf(T), q, name, value[0].String())
}

func (l *language) SetList(T qlova.Type, q *qlova.Script, name string, value qlova.List) {
	Set(q, name, value.String())
}

func (l *language) IndexList(list string, index string) string {	
	return list+"["+index+".Int64() % int64(len("+list+"))]"
}

func (l *language) LengthList(list string) string {	
	l.importsList = append(l.importsList, "math/big")
	
	return "big.NewInt(int64(len("+list+")))"
}

func (l *language) LiteralList(T qlova.Type, elements ...qlova.Type) string {
	l.importsList = append(l.importsList, "math/big")
	
	var literal = "[]"+l.GoTypeOf(T)+"{"

	for i := range elements {
		literal += elements[i].String()
		
		if i < len(elements)-1 {
			literal += ","
		}
	}
	literal += "}"
	return literal
}

func (l *language) NumberToList(T qlova.Type, number string) string {
	l.importsList = append(l.importsList, "math/big")
	
	return "make([]"+l.GoTypeOf(T)+", "+number+".Int64())"
}
