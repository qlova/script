package Go

import "github.com/qlova/script/language"

const Name = "Go"

type implementation struct {
	imports []string
	helpers []string
}

func (l *implementation) Import(pkg string) {
	for i := range l.imports {
		if l.imports[i] == pkg {
			return
		}
	}
	l.imports = append(l.imports, pkg)
}


func (l *implementation) AddHelper(helper string) {
	for i := range l.helpers {
		if l.helpers[i] == helper {
			return
		}
	}
	l.helpers = append(l.helpers, helper)
}

func Language() *implementation {
	return new(implementation)
}

const NumberHelper = `func Number(representation string) *big.Int {
	var z big.Int
	z.SetString(representation, 10)
	return &z
}`


func (l *implementation) GetExpression(T language.Type) string {
	var PanicName = "Error in "+Name+".GetExpression(Type)"
	
	switch T.(type) {
		case language.String:
			return string(T.(String))
			
		case language.Number:
			n := T.(Number)
			
			if n.Literal == nil {
				return n.Expression
			} else {
				
				l.Import("math/big")
				
				if n.Literal.IsInt64() {
					return "big.NewInt("+n.Literal.String()+")"
				}
				
				l.AddHelper(NumberHelper)
				
				return "Number("+n.Literal.String()+")"
			}
		
		case language.Switch, language.Symbol, 
			language.Custom, language.Stream, language.List, language.Array, 
			language.Table, language.Error, language.Float, language.Pointer, 
			language.Dynamic, language.Function, language.Metatype, language.FunctionType:
		
		panic(PanicName+": Unimplented")
			
		default:
			panic(PanicName+": Invalid Type")
	}
	
	return ""
}

//TODO remove.
func init() {
	language.Default = language.Interface(Language())
}


func (l *implementation) Init() {}
func (l *implementation) Head() language.Statement {
	var result = "package main\n\n"

	for _, pkg := range l.imports {
		result += `import "`+pkg+`"`+"\n"
	}
	
	result += "\n"
	
	for _, helper := range l.helpers {
		result += helper
	}

	return language.Statement(result)
}
func (l *implementation) Neck() language.Statement { return "" }
func (l *implementation) Body() language.Statement { return "" }
func (l *implementation) Tail() language.Statement { return "" }
func (l *implementation) Last() language.Statement { return "" }

//Returns a Statement that begins the main entry point to the program.
func (l *implementation) Main() language.Statement {
	return "func main() {\n"
}

//Returns a Statement that exits the program.
func (l *implementation) Exit() language.Statement {
	l.Import("os")
	l.imports = append(l.imports, "os")
	
	return "os.Exit(0)"
}

//Returns a Statement that ends the main entry point to the program.
func (l *implementation) EndMain() language.Statement {
	return "}"
}
