package Go

import "github.com/qlova/script/language"

func (implementation Implementation) Cast(a, b language.Type) language.Type {
	
	
	switch a.(type) {
		case Integer:
			switch b.(type) {
				case String:
					return String{Expression: "fmt.Sprint("+implementation.ExpressionOf(a)+")"}
					
				case Symbol:
					return Symbol{Expression: "rune("+implementation.ExpressionOf(a)+")"}
			}

		case Symbol:
			switch b.(type) {
				case String:
					return String{Expression: "string("+implementation.ExpressionOf(a)+")"}
					
				case Integer:
					return Integer{Expression: "int("+implementation.ExpressionOf(a)+")"}
			}
	}
	
	panic(implementation.Name()+".Cast("+a.Name()+", "+b.Name()+") Unimplemented")
	return nil
}

