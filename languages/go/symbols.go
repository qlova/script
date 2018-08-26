package Go

import qlova "github.com/qlova/script"

func (l *language) NewSymbol(q *qlova.Script, name string, value ...qlova.Symbol) {	
	New("rune", q, name, value[0].String())
}

func (l *language) SetSymbol(q *qlova.Script, name string, value qlova.Symbol) {	
	Set(q, name, value.String())
}

func (l *language) LiteralSymbol(symbol string) string {
	return symbol
}

func (l *language) NumberToSymbol(number string) string {
	return "rune("+number+".Int64())"
}

func (l *language) SymbolToNumber(number string) string {
	l.importsList = append(l.importsList, "math/big")

	return "big.NewInt(int64("+number+"))"
} 

func (l *language) SymbolToString(symbol string) string {
	l.importsList = append(l.importsList, "math/big")

	return "string("+symbol+")"
} 

//TODO deal with errors!
func (l *language) ReadSymbol(symbol string) string {
	l.importsList = append(l.importsList, "bufio")
	l.helpersList = append(l.helpersList, `var BStdin = bufio.NewReader(os.Stdin) 
	func ReadSymbol(symbol rune) string {
	result, _ := BStdin.ReadString(byte(symbol))
	return result
}
`)

	return "ReadSymbol("+symbol+")"
}
