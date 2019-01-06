package interpreter

import "github.com/qlova/script/language"

func (implementation Implementation) Type(name string, registers []string, elements []language.Type) language.Statement {
	panic(implementation.Name()+".Type() Unimplemented")
	return language.Statement("")
}

func (implementation Implementation) Method(t string, name string, registers []string, arguments []language.Type, returns language.Type) language.Statement {
	panic(implementation.Name()+".Method() Unimplemented")
	return language.Statement("")
}

func (implementation Implementation) This() language.Type {
	panic(implementation.Name()+".This() Unimplemented")
	return nil
}

func (implementation Implementation) New(name string) language.Type {
	panic(implementation.Name()+".New() Unimplemented")
	return nil
}

func (implementation Implementation) Invoke(t language.Type, method string, arguments []language.Type) language.Type {
	panic(implementation.Name()+".Invoke() Unimplemented")
	return nil
}

func (implementation Implementation) Execute(t language.Type, method string, arguments []language.Type) language.Statement {
	panic(implementation.Name()+".Execute() Unimplemented")
	return language.Statement("")
}

func (implementation Implementation) EndMethod() language.Statement {
	panic(implementation.Name()+".EndMethod() Unimplemented")
	return language.Statement("")
}

