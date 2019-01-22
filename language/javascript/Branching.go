package Javascript

import "github.com/qlova/script/language"

func (implementation Implementation) If(condition language.Bit) language.Statement {
	return language.Statement("if ("+implementation.ExpressionOf(condition)+") {")
}

func (implementation Implementation) ElseIf(condition language.Bit) language.Statement {
	return language.Statement("} else if ("+implementation.ExpressionOf(condition)+") {")
}

func (implementation Implementation) Else() language.Statement {
	return language.Statement("} else {")
}

func (implementation Implementation) EndIf() language.Statement {
	return language.Statement("}")
}

