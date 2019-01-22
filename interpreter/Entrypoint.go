package interpreter

import "github.com/qlova/script/language"

func (implementation Implementation) Main() language.Statement {
	*implementation.active = implementation.program.CreateBlock()
	return language.Statement("")
}

func (implementation Implementation) EndMain() language.Statement {
	return language.Statement("")
}

func (implementation Implementation) Exit() language.Statement {
	panic(implementation.Name()+".Exit() Unimplemented")
	return language.Statement("")
}

