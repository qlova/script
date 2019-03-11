package script

import "github.com/qlova/script/language"

type Bool struct {
	internal language.Bit
	script Script
	literal *bool
}

func (b Bool) Value() Value {
	return Value{
		script: b.script,
		internal: b.LanguageType(),
	}
}

func (q Script) Bool(b ...bool) Bool {
	if len(b) > 0 {
		return q.BoolFromLanguageType(q.lang.Bit(b[0]))
	}
	return q.BoolFromLanguageType(q.lang.Bit(false))
}

func (b Bool) True() {
	b.script.indent()
	b.script.write(b.script.lang.Set(b.LanguageType(), b.script.lang.Bit(true)))
}

func (b Bool) False() {
	b.script.indent()
	b.script.write(b.script.lang.Set(b.LanguageType(), b.script.lang.Bit(false)))
}

func (q Script) Not(b Bool) Bool {
	return q.BoolFromLanguageType(q.lang.Not(b.LanguageType().(language.Bit)))
}

//Get this value as a string or cast to a string.
func (v Value) Bool() Bool {
	if b, ok := v.internal.(language.Bit); ok {
		return Bool{
			script: v.script,
			internal: b,
		}
	}

	return v.script.BoolFromLanguageType(v.script.lang.Cast(v.internal, v.script.Bool().LanguageType()))
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
