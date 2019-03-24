package Go

import "github.com/qlova/script/language"

func (implementation Implementation) Loop() language.Statement {
	panic(implementation.Name()+".Loop() Unimplemented")
	return language.Statement("")
}

func (implementation Implementation) EndLoop() language.Statement {
	panic(implementation.Name()+".EndLoop() Unimplemented")
	return language.Statement("")
}

func (implementation Implementation) Break() language.Statement {
	panic(implementation.Name()+".Break() Unimplemented")
	return language.Statement("")
}

func (implementation Implementation) While(condition language.Bit) language.Statement {
	panic(implementation.Name()+".While() Unimplemented")
	return language.Statement("")
}

func (implementation Implementation) EndWhile() language.Statement {
	panic(implementation.Name()+".EndWhile() Unimplemented")
	return language.Statement("")
}

func (implementation Implementation) ForRange(i string, a, b language.Number) (language.Statement, language.Type) {
	panic(implementation.Name()+".ForRange() Unimplemented")
	return language.Statement(""), nil
}

func (implementation Implementation) EndForRange() language.Statement {
	panic(implementation.Name()+".EndForRange() Unimplemented")
	return language.Statement("")
}

func (implementation Implementation) ForEach(i, v string, list language.Type) (language.Statement, language.Type, language.Type) {
	
	if implementation.Flag("use") {
		implementation.neck.WriteString("func use(variable interface{}) {}\n\n")
	}
	
	return language.Statement("for "+i+", "+v+" := range "+string(list.Raw())+" { use("+i+"); use("+v+");\n"), Integer{Expression:language.Statement(i)}, list.(List).Subtype.Register(v)
}

func (implementation Implementation) EndForEach() language.Statement {
	return language.Statement("}\n")
}

func (implementation Implementation) For(i string, condition language.Bit, action language.Statement) (language.Statement, language.Type) {
	panic(implementation.Name()+".For() Unimplemented")
	return language.Statement(""), nil
}

func (implementation Implementation) EndFor() language.Statement {
	panic(implementation.Name()+".EndFor() Unimplemented")
	return language.Statement("")
}

