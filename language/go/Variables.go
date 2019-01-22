package Go

import "github.com/qlova/script/language"

func (implementation Implementation) Register(register string, value language.Type) (language.Statement, language.Type) {
	return language.Statement("var "+register+" = ")+implementation.ExpressionOf(value)+"\n", value.Register(register)
}

func (implementation Implementation) Set(variable, value language.Type) language.Statement {
	return language.Statement(implementation.ExpressionOf(variable)+" = "+implementation.ExpressionOf(value)+"\n")
}
