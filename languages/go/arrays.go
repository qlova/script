package Go

import qlova "github.com/qlova/script"

func (l *language) NewArray(T qlova.Type,q *qlova.Script, name string, value ...qlova.Array) {
	New("[]"+l.GoTypeOf(T), q, name, value[0].String())
}

func (l *language) SetArray(T qlova.Type, q *qlova.Script, name string, value qlova.Array) {
	Set(q, name, value.String())
}

func (l *language) IndexArray(array string, index string) string {	
	return array+"["+index+".Int64() % int64(len("+array+"))]"
}

func (l *language) LengthArray(array string) string {	
	l.importsList = append(l.importsList, "math/big")
	
	return "big.NewInt(int64(len("+array+")))"
}

func (l *language) LiteralArray(T qlova.Type, elements ...qlova.Type) string {
	l.importsList = append(l.importsList, "math/big")
	
	var literal = "[...]"+l.GoTypeOf(T)+"{"

	for i := range elements {
		literal += elements[i].String()
		
		if i < len(elements)-1 {
			literal += ","
		}
	}
	literal += "}"
	return literal
}

func (l *language) NumberToArray(T qlova.Type, number string) string {
	l.importsList = append(l.importsList, "math/big")
	
	return "make([]"+l.GoTypeOf(T)+", "+number+".Int64())"
}
