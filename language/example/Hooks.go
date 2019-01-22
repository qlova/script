package Example

import "github.com/qlova/script/language"

func (implementation Implementation) Head() language.Statement {
	panic(implementation.Name()+".Head() Unimplemented")
	return language.Statement("")
}

func (implementation Implementation) Neck() language.Statement {
	panic(implementation.Name()+".Neck() Unimplemented")
	return language.Statement("")
}

func (implementation Implementation) Body() language.Statement {
	panic(implementation.Name()+".Body() Unimplemented")
	return language.Statement("")
}

func (implementation Implementation) Tail() language.Statement {
	panic(implementation.Name()+".Tail() Unimplemented")
	return language.Statement("")
}

func (implementation Implementation) Last() language.Statement {
	panic(implementation.Name()+".Last() Unimplemented")
	return language.Statement("")
}

