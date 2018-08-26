package Go
 
import qlova "github.com/qlova/script" 
 
func (l *language) NewString(q *qlova.Script, name string, value ...qlova.String) {
	q.Body.WriteString("var ")
	q.Body.WriteString(name)
	
	if len(value) == 0 {
		q.Body.WriteString(" string")
	} else {
		q.Body.WriteString(" = ")
		q.Body.WriteString(value[0].String())
		q.Body.WriteString("\n")
	}
}

func (l *language) AddStrings(a, b string) string {
	return a+" + "+b
}

func (l *language) SetString(q *qlova.Script, name string, value qlova.String) {
	q.Body.WriteString(name)
	q.Body.WriteString(" = ")
	q.Body.WriteString(value.String())
	q.Body.WriteString("\n")
}

func (l *language) StringToNumber(number string) string {
	l.importsList = append(l.importsList, "math/big")
	l.helpersList = append(l.helpersList, `func StringToNumber(number string) *big.Int {
	var z big.Int
	z.SetString(number, 10)
	return &z
}
`)
	
	return "StringToNumber("+number+")"
}
