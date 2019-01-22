package script

import "github.com/qlova/script/language"

type Bool struct {
	internal language.Bit
	script Script
	literal *bool
}

func (q Script) Not(b Bool) Bool {
	return q.BoolFromLanguageType(q.lang.Not(b.LanguageType().(language.Bit)))
}

func (*Bool) wrap(T interface{}) Bool {
	if i, ok := T.(language.Bit); ok {
		return Bool{internal: i}
	}
	panic("Invalid wrap!")
	return Bool{}
}

//Cast an Int to a language.Type ready to be passed to the method of a Language.
func (b Bool) LanguageType() language.Type {
	if b.literal != nil {
		return b.script.lang.Bit(*b.literal)
	} else {
		return b.internal
	}
}

//Wrap a language.Type to an Integer.
func (q Script) BoolFromLanguageType(T language.Type) Bool {
	if internal, ok := T.(language.Bit); ok {
		return Bool{
			internal: internal,
			script: q,
		}
	}
	panic("Invalid wrap!")
	return Bool{}
}
