package Example

import "github.com/qlova/script/language"

func (implementation Implementation) Function(name string, registers []string, arguments []language.Type, returns language.Type) language.Statement {
	panic(implementation.Name()+".Function() Unimplemented")
	return language.Statement("")
}

func (implementation Implementation) EndFunction() language.Statement {
	panic(implementation.Name()+".EndFunction() Unimplemented")
	return language.Statement("")
}

func (implementation Implementation) Call(name string, arguments []language.Type) language.Type {
	panic(implementation.Name()+".Call() Unimplemented")
	return nil
}

func (implementation Implementation) Run(name string, arguments []language.Type) language.Statement {
	panic(implementation.Name()+".Run() Unimplemented")
	return language.Statement("")
}

func (implementation Implementation) Return(value language.Type) language.Statement {
	panic(implementation.Name()+".Return() Unimplemented")
	return language.Statement("")
}

