package Javascript

import "github.com/qlova/script/language"

func (implementation Implementation) Register(register string, value language.Type) (language.Statement, language.Type) {
	panic(implementation.Name()+".Register() Unimplemented")
	return language.Statement(""), nil
}

func (implementation Implementation) Set(variable, value language.Type) language.Statement {
	return language.Statement(variable.Raw()+" = "+value.Raw()+";")
}
