package example

import "github.com/qlova/script/language"

func (implementation Implementation) If(condition language.Bit) language.Statement {
	panic(implementation.Name()+".If() Unimplemented")
	return language.Statement("")
}

func (implementation Implementation) ElseIf(condition language.Bit) language.Statement {
	panic(implementation.Name()+".ElseIf() Unimplemented")
	return language.Statement("")
}

func (implementation Implementation) Else() language.Statement {
	panic(implementation.Name()+".Else() Unimplemented")
	return language.Statement("")
}

func (implementation Implementation) EndIf() language.Statement {
	panic(implementation.Name()+".EndIf() Unimplemented")
	return language.Statement("")
}

