package script

import "github.com/qlova/script/language"

type Symbol struct {
	language.Symbol
	EmbeddedScript

	Literal *rune
}

//Converts a Go rune to a Symbol.
func (q *Script) Symbol(s ...rune) Symbol {
	var r rune = 0
	if len(s) > 0 {
		r = s[0]
	}
	return Symbol{Literal: &r, EmbeddedScript: EmbeddedScript{ q: q }}
}
