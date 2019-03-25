package script

import "github.com/qlova/script/language"

type Value struct {
	script Script
	internal language.Type
	
	//Array info
	subtype Type
	length int
	
	arguments []Type
}

func (v Value) LanguageType() language.Type {
	return v.internal
}


//Wrap a language.Type to an Integer.
func (q Script) ValueFromLanguageType(T language.Type) Value {
	return Value{
		internal: T,
		script: q,
	}
}

func (v Value) Is(t Type) bool {
	return v.LanguageType().Is(t.LanguageType())
}

func (v Value) IsArray() bool {
	_, ok := v.LanguageType().(language.Array)
	return ok
}

func (v Value) IsBool() bool {
	_, ok := v.LanguageType().(language.Bit)
	return ok
}

func (v Value) IsString() bool {
	_, ok := v.LanguageType().(language.String)
	return ok
}

func (v Value) IsInt() bool {
	_, ok := v.LanguageType().(language.Integer)
	return ok
}

func (v Value) IsList() bool {
	_, ok := v.LanguageType().(language.List)
	return ok
}

func (v Value) Value() Value {
	return v
}

func (v Value) Set(value Type) {
	var q = v.script
	q.indent()
	q.write(q.lang.Set(v.internal, value.LanguageType()))
}

func (v Value) Var(name ...string) Value {
	var register string
	if len(name) > 0 {
		register = name[0]
	} else {
		register = Unique()
	}
	
	v.script.indent()
	statement, variable := v.script.lang.Register(register, v.LanguageType())
	v.script.write(statement)
	
	var result = v
	result.internal = variable
	return result
}
