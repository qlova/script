package example

import "github.com/qlova/script/language"

func (implementation Implementation) Error() language.Error {
	panic(implementation.Name()+".Error() Unimplemented")
	return nil
}

func (implementation Implementation) Trace(line int, file string) language.Statement {
	panic(implementation.Name()+".Trace() Unimplemented")
	return language.Statement("")
}

func (implementation Implementation) Throw(code language.Number, message language.String) language.Statement {
	panic(implementation.Name()+".Throw() Unimplemented")
	return language.Statement("")
}

func (implementation Implementation) Catch() language.Bit {
	panic(implementation.Name()+".Catch() Unimplemented")
	return nil
}

