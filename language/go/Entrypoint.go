package Go

import "github.com/qlova/script/language"

func (implementation Implementation) Main() language.Statement {
	return language.Statement("func main() {\n")
}

func (implementation Implementation) EndMain() language.Statement {
	return language.Statement("}\n")
}

func (implementation Implementation) Exit() language.Statement {
	panic(implementation.Name()+".Exit() Unimplemented")
	return language.Statement("")
}

