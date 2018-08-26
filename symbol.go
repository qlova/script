package script

//import "math/big"
import "errors"

type Symbol string

func (n Symbol) Name() string {
	return "symbol"
}

func (n Symbol) Equals(t interface{}) bool {
	if _, ok := t.(Symbol); ok {
		return true
	}
	return false
}

func (n Symbol) String() string {
	return string(n)
}

func (q *Script) LiteralSymbol(symbol string) Symbol {
	return Symbol(q.Language.LiteralSymbol(symbol))
}

func (q *Script) NewSymbol(name string, value Symbol) {
	q.IndentBody()
	q.Language.NewSymbol(q, name, value)
}

func (q *Script) SetSymbol(name string, value Symbol) {
	q.IndentBody()
	q.Language.SetSymbol(q, name, value)
}

func (q *Script) ReadSymbol(symbol Symbol) String {
	return String(q.Language.ReadSymbol(symbol.String()))
}

func (q *Script) ToSymbol(T Type) (Symbol, error) {
	switch T.(type) {
		
		case Symbol:
			return T.(Symbol), nil
		
		case Number:
			return Symbol(q.Language.NumberToSymbol(string(T.(Number)))), nil
		
		default:
			return "", errors.New("Cannot convert "+T.Name()+" to symbol!")
	}
}
