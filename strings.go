package script

import "github.com/qlova/script/language"

//A String is an ordered sequence of symbols, often used to denote words, sentences and text.
type String struct {
	script Script
	internal language.String

	literal *string
}

//Return a new String type with the value s.
func (q Script) String(s ...string) String {
	var literal = ""

	if len(s) > 0 {
		literal = s[0]
	}

	return String{
		script: q,
		literal: &literal,
	}
}

//Cast a String to a language.Type ready to be passed to the method of a Language.
func (s String) LanguageType() language.Type {
	if s.literal != nil {
		return s.script.lang.String(*s.literal)
	} else {
		return s.internal
	}
}

//Wrap a language.Type to an Integer.
func (q Script) StringFromLanguageType(T language.Type) String {
	if internal, ok := T.(language.String); ok {
		return String{
			internal: internal,
			script: q,
		}
	}
	panic("Invalid wrap!")
	return String{}
}

func (a String) Equals(b String) Bool {
	return a.script.BoolFromLanguageType(a.script.lang.Equals(a.LanguageType(), b.LanguageType()))
}
