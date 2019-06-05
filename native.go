package script

import "github.com/qlova/script/language"

//A native value is a custom type defined in the target language.
type Native struct {
	script   Script
	internal language.Native
}

//Cast an NativeObject to a language.Type ready to be passed to the method of a Language.
func (n Native) LanguageType() language.Type {
	return n.internal
}

func (n Native) Value() Value {
	return Value{
		script:   n.script,
		internal: n.LanguageType(),
	}
}

//Return this Native as a variable (optionally named).
func (n Native) Var(name ...string) Native {
	var register = unique(name)

	n.script.indent()
	statement, variable := n.script.lang.Register(register, n.LanguageType())
	n.script.write(statement)

	return Native{
		script:   n.script,
		internal: variable.(language.Native),
	}
}

//Wrap a language.Type to a Native.
func (q Script) NativeFromLanguageType(T language.Type) Native {
	if internal, ok := T.(language.Native); ok {
		return Native{
			internal: internal,
			script:   q,
		}
	}
	panic("Invalid wrap!")
	return Native{}
}
