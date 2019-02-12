package example

import "github.com/qlova/script/language"

func (implementation Implementation) Main() language.Statement {
	panic(implementation.Name()+".Main() Unimplemented")
	return language.Statement("")
}

func (implementation Implementation) EndMain() language.Statement {
	panic(implementation.Name()+".EndMain() Unimplemented")
	return language.Statement("")
}

func (implementation Implementation) Exit() language.Statement {
	panic(implementation.Name()+".Exit() Unimplemented")
	return language.Statement("")
}

