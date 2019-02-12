package Go

import "github.com/qlova/script/language"

func (implementation Implementation) If(condition language.Bit) language.Statement {
	return language.Statement(`if (`+condition.Raw()+`) {\n`)
}

func (implementation Implementation) ElseIf(condition language.Bit) language.Statement {
	return language.Statement(`} else if (`+condition.Raw()+`) {\n`)
}

func (implementation Implementation) Else() language.Statement {
	return language.Statement(`} else {`)
}

func (implementation Implementation) EndIf() language.Statement {
	return language.Statement(`}`)
}

