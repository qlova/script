package Go
 
import qlova "github.com/qlova/script"
 
//TODO allow using native int64, int32 types.
 
func (l *language) NewNumber(q *qlova.Script, name string, value ...qlova.Number) {
	l.Import(q, "math/big")
	
	New("*big.Int", q, name, value[0])
}

func (l *language) SetNumber(q *qlova.Script, name string, value qlova.Number) {
	Set(q, name, value)
}

func (l *language) LiteralNumber(number string) string {
	l.importsList = append(l.importsList, "math/big")
	
	return "Number("+strconv.Quote(number)+")"
}

 
 func (l *language) NumberToString(number string) string {
	l.helpersList = append(l.helpersList, `func NumberToString(number *big.Int) string {
	if number == nil {
		return "0"
	} else {
		return number.String()
	}
}
`)
	
	return "NumberToString("+number+")"
}

func (l *language) Add(a, b string) string {
	l.helpersList = append(l.helpersList, `func Add(a, b *big.Int) *big.Int {
	var z big.Int
	z.Add(a, b)
	return &z
}
`)
	
	return "Add("+a+","+b+")"
}

func (l *language) Subtract(a, b string) string {
	l.helpersList = append(l.helpersList, `func Subtract(a, b *big.Int) *big.Int {
	var z big.Int
	z.Sub(a, b)
	return &z
}
`)
	
	return "Subtract("+a+","+b+")"
}

func (l *language) Multiply(a, b string) string {
	l.helpersList = append(l.helpersList, `func Mutiply(a, b *big.Int) *big.Int {
	var z big.Int
	z.Mul(a, b)
	return &z
}
`)
	
	return "Mutiply("+a+","+b+")"
}

func (l *language) Divide(a, b string) string {
	l.helpersList = append(l.helpersList, `func Divide(a, b *big.Int) *big.Int {
	var z big.Int
	z.Div(a, b)
	return &z
}
`)
	
	return "Divide("+a+","+b+")"
}

func (l *language) Modulo(a, b string) string {
	l.helpersList = append(l.helpersList, `func Modulo(a, b *big.Int) *big.Int {
	var z big.Int
	z.Mod(a, b)
	return &z
}
`)
	
	return "Modulo("+a+","+b+")"
}

func (l *language) Power(a, b string) string {
	l.helpersList = append(l.helpersList, `func Power(a, b *big.Int) *big.Int {
	var z big.Int
	z.Exp(a, b, nil)
	return &z
}
`)
	
	return "Power("+a+","+b+")"
}
