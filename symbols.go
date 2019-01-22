package script

import "github.com/qlova/script/language"

type Rune struct {
	script Script
	internal language.Symbol

	literal *rune
}

//Return a new String type with the value s.
func (q Script) Rune(s ...rune) Rune {
	var literal rune

	if len(s) > 0 {
		literal = s[0]
	}

	return Rune{
		script: q,
		literal: &literal,
	}
}

func (r Rune) LanguageType() language.Type {
	if r.literal != nil {
		return r.script.lang.Symbol(*r.literal)
	} else {
		return r.internal
	}
}

func (r Rune) Value() Value {
	return Value{
		script: r.script,
		internal: r.LanguageType(),
	}
}

func (v Value) Rune() Rune {
	if r, ok := v.internal.(language.Symbol); ok {
		return Rune{
			script: v.script,
			internal: r,
		}
	}

	return v.script.RuneFromLanguageType(v.script.lang.Cast(v.internal, v.script.Rune().LanguageType()))
}

//Wrap a language.Type to an Integer.
func (q Script) RuneFromLanguageType(T language.Type) Rune {
	if internal, ok := T.(language.Symbol); ok {
		return Rune{
			internal: internal,
			script: q,
		}
	}
	panic("Invalid wrap!")
	return Rune{}
}

func (a Rune) Equals(b Rune) Bool {
	return a.script.BoolFromLanguageType(a.script.lang.Equals(a.LanguageType(), b.LanguageType()))
}
