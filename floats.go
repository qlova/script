package script


import "github.com/qlova/script/language"

//An Int is a numeric integer value consisting of a magnitude and a sign.
type Float struct {
	script Script
	internal language.Integer

	literal *float64
}

//Cast an Int to a language.Type ready to be passed to the method of a Language.
func (f Float) LanguageType() language.Type {
	if f.literal != nil {
		return f.script.lang.Real(*f.literal)
	} else {
		return f.internal
	}
}

func (f Float) Value() Value {
	return Value{
		script: f.script,
		internal: f.LanguageType(),
	}
}

//Return a new String type with the value s.
func (q Script) Float(f ...float64) Float {
	var literal float64 = 0

	if len(f) > 0 {
		literal = f[0]
	}

	return Float{
		script: q,
		literal: &literal,
	}
}
