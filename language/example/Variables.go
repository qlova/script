package Example

import "github.com/qlova/script/language"

func (implementation Implementation) Register(register string, value language.Type) (language.Statement, language.Type) {
	panic(implementation.Name()+".Register() Unimplemented")
	return nil
}

func (implementation Implementation) Set(register string, value language.Type) language.Statement {
	panic(implementation.Name()+".Set() Unimplemented")
	return language.Statement("")
}

func (implementation Implementation) Get(register string, value language.Type) language.Type {
	panic(implementation.Name()+".Get() Unimplemented")
	return nil
}

