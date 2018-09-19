package script

import "math/big"
import "github.com/qlova/script/language"

type Number struct {
	language.Number
	EmbeddedScript
	
	Literal *big.Int
}

func (Number) SameAs(i interface{}) bool { _, ok := i.(Number); return ok }

//Converts a Go integer to a language.Number.
func (q *Script) Number(n ...int) Number {
	if len(n) > 0 {
		num := n[0]
		return Number{Literal: big.NewInt(int64(num)), EmbeddedScript: EmbeddedScript{ q: q }}
	}
	return Number{Literal: new(big.Int), EmbeddedScript: EmbeddedScript{ q: q }}
}

//Converts a Go integer to a language.Number.
func (q *Script) BigNumber(i *big.Int) Number {
	return Number{Literal: i, EmbeddedScript: EmbeddedScript{ q: q }}
}

func (q *Script) Add(a, b Number) Number {
	return q.wrap(q.lang.Add(convert(a).(language.Number), convert(b).(language.Number))).(Number)
}

func (q *Script) Sub(a, b Number) Number {
	return q.wrap(q.lang.Sub(convert(a).(language.Number), convert(b).(language.Number))).(Number)
}

func (q *Script) Pow(a, b Number) Number {
	return q.wrap(q.lang.Pow(convert(a).(language.Number), convert(b).(language.Number))).(Number)
}

func (q *Script) Mul(a, b Number) Number {
	return q.wrap(q.lang.Mul(convert(a).(language.Number), convert(b).(language.Number))).(Number)
}

func (q *Script) Div(a, b Number) Number {
	return q.wrap(q.lang.Div(convert(a).(language.Number), convert(b).(language.Number))).(Number)
}

func (q *Script) Mod(a, b Number) Number {
	return q.wrap(q.lang.Mod(convert(a).(language.Number), convert(b).(language.Number))).(Number)
}
