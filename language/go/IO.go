package Go

import "github.com/qlova/script/language"

func (implementation Implementation) Print(values ...language.Type) language.Statement {
	implementation.Import("fmt")
	
	var statement language.Statement = "fmt.Println("
	for i, value := range values {
		statement += implementation.ExpressionOf(value)
		
		if i < len(values)-1 {
			statement += ","
		}
	}
	
	statement += ")\n"
	
	return statement
}

func (implementation Implementation) Write(stream language.Stream, values ...language.Type) language.Statement {
	panic(implementation.Name()+".Write() Unimplemented")
	return language.Statement("")
}

