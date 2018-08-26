package script

//import "math/big"
import "errors"

type Number string

func (n Number) Name() string {
	return "number"
}

func (n Number) Equals(t interface{}) bool {
	if _, ok := t.(Number); ok {
		return true
	}
	return false
}

func (n Number) String() string {
	return string(n)
}

func (q Script) Add(a, b Number) Number {
	return Number(q.Language.Add(a.String(), b.String()))
}

func (q Script) Subtract(a, b Number) Number {
	return Number(q.Language.Subtract(a.String(), b.String()))
}

func (q Script) Multiply(a, b Number) Number {
	return Number(q.Language.Multiply(a.String(), b.String()))
}

func (q Script) Divide(a, b Number) Number {
	return Number(q.Language.Divide(a.String(), b.String()))
}

func (q Script) Modulo(a, b Number) Number {
	return Number(q.Language.Modulo(a.String(), b.String()))
}

func (q Script) Power(a, b Number) Number {
	return Number(q.Language.Power(a.String(), b.String()))
}

func (q *Script) LiteralNumber(number string) Number {
	return Number(q.Language.LiteralNumber(number))
}

func (q *Script) NewNumber(name string, value Number) {
	q.IndentBody()
	q.Language.NewNumber(q, name, value)
}

func (q *Script) SetNumber(name string, value Number) {
	q.IndentBody()
	q.Language.SetNumber(q, name, value)
}

func (q *Script) ToNumber(T Type) (Number, error) {
	switch T.(type) {
		
		case Number:
			return T.(Number), nil
		
		case String:
			return Number(q.Language.StringToNumber(string(T.(String)))), nil
		
		case Symbol:
			return Number(q.Language.SymbolToNumber(string(T.(Symbol)))), nil
		
		default:
			return "", errors.New("Cannot convert "+T.Name()+" to number!")
	}
}
