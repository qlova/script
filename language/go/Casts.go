package Go

import "github.com/qlova/script/language"

func (implementation Implementation) Cast(a, b language.Type) language.Type {
	
	
	switch a.(type) {
		case Integer:
			switch b.(type) {
				case String:
					implementation.Import("fmt")
					return String{Expression: "fmt.Sprint("+implementation.ExpressionOf(a)+")"}
					
				case Symbol:
					return Symbol{Expression: "rune("+implementation.ExpressionOf(a)+")"}
					
				case Bit:
					return Bit{Expression: "("+implementation.ExpressionOf(a)+" != 0)"}
			}

		case Symbol:
			switch b.(type) {
				case String:
					return String{Expression: "string("+implementation.ExpressionOf(a)+")"}
					
				case Integer:
					return Integer{Expression: "int("+implementation.ExpressionOf(a)+")"}
			}
			
		case String:
			switch b.(type) {
				case Integer:
					implementation.Import("strconv")
					if implementation.Flag("atoi") {
						implementation.neck.WriteString("func atoi(text string) int {\n\ti, _ := strconv.Atoi(text)\n\treturn i\n}\n\n")
					}
					return Integer{Expression: "atoi("+implementation.ExpressionOf(a)+")"}
			}
			
		case Bit:
			switch b.(type) {
				case Integer:
					if implementation.Flag("btoi") {
						implementation.neck.WriteString("func btoi(bit bool) int {if bit{return 1}else{return 0}}\n\n")
					}
					return Integer{Expression: "btoi("+implementation.ExpressionOf(a)+")"}
				case String:
					implementation.Import("fmt")
					return String{Expression: "fmt.Sprint("+implementation.ExpressionOf(a)+")"}
			}
	}
	
	panic(implementation.Name()+".Cast("+a.Name()+", "+b.Name()+") Unimplemented")
	return nil
}

