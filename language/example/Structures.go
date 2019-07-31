package example

import "github.com/qlova/script/language"

func (implementation Implementation) Index(structure, index language.Type) language.Type {
	panic(implementation.Name() + ".Index() Unimplemented")
	return nil
}

func (implementation Implementation) Modify(structure, index, value language.Type) language.Statement {
	panic(implementation.Name() + ".Modify() Unimplemented")
	return language.Statement("")
}
