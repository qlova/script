package interpreter

import "github.com/qlova/script/language"

func (implementation Implementation) Main() language.Statement {
	implementation.CreateBlock()
	var block = implementation.Active()
		block.Main = true
	return language.Statement("")
}

func (implementation Implementation) EndMain() language.Statement {
	return language.Statement("")
}

func (implementation Implementation) Exit() language.Statement {
	panic(implementation.Name()+".Exit() Unimplemented")
	return language.Statement("")
}

