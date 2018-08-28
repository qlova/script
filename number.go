package script

import "math/big"
import "github.com/qlova/script/language"

//Converts a Go integer to a language.Number.
func (q Script) Number(i int) language.Number {
	return q.lang.LiteralNumber(big.NewInt(int64(i)))
}

/*	
	Define 'name' to be a Number with an optional 'value' of type Number.
	Returns a number.
	
	Example:
		script.DefineNumber(name) -> var name int
		script.DefineNumber(name, value) -> var name int = value
*/
func (q *Script) DefineNumber(name string, value ...language.Number) language.Number {
	q.indent()
	
	var initial language.Number
	if len(value) == 0 {
		initial = q.Number(0)
	} else {
		initial = value[0]
	}
	
	number, statement := q.lang.Define(name, initial)

	q.write(statement)

	return number.(language.Number)
}

func (q *Script) Add(a, b language.Number) language.Number {
	return q.lang.Add(a, b)
}

func (q *Script) Sub(a, b language.Number) language.Number {
	return q.lang.Sub(a, b)
}

func (q *Script) Pow(a, b language.Number) language.Number {
	return q.lang.Pow(a, b)
}

func (q *Script) Mul(a, b language.Number) language.Number {
	return q.lang.Mul(a, b)
}

func (q *Script) Div(a, b language.Number) language.Number {
	return q.lang.Div(a, b)
}

func (q *Script) Mod(a, b language.Number) language.Number {
	return q.lang.Mod(a, b)
}
