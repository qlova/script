package script

import "github.com/qlova/script/language"

type Value struct {
	script Script
	internal language.Type
}

func (v Value) LanguageType() language.Type {
	return v.internal
}

func (v Value) Value() Value {
	return v
}

func (v Value) Set(value Type) {
	var q = v.script
	q.indent()
	q.write(q.lang.Set(v.internal, value.LanguageType()))
}

func (v Value) Var(name ...string) Type {
	var register string
	if len(name) > 0 {
		register = name[0]
	} else {
		register = Unique()
	}
	
	v.script.indent()
	statement, variable := v.script.lang.Register(register, v.LanguageType())
	v.script.write(statement)
	
	return Value{
		script: v.script,
		internal: variable,
	}
}
