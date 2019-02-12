package example

import "github.com/qlova/script/language"

func (implementation Implementation) Print(values ...language.Type) language.Statement {
	panic(implementation.Name()+".Print() Unimplemented")
	return language.Statement("")
}

func (implementation Implementation) Write(stream language.Stream, values ...language.Type) language.Statement {
	panic(implementation.Name()+".Write() Unimplemented")
	return language.Statement("")
}

