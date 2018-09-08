package Go

import "github.com/qlova/script/language"
import "fmt"
import "os/exec"
import "strconv"

const Name = "Go"

const NumberType = "Number"
const NumberImport = "math/big"
const NumberTypeDefinition = `type Number struct { 
	Small int64
	Large *big.Int
}

`
const NumberToInt64 = `func (number Number) Int64() int64 { 
	if number.Large == nil { 
		return number.Small
	} 
	return number.Large.Int64() 
}

`
const NumberToString = `func (number Number) String() string { 
	if number.Large == nil { 
		return fmt.Sprint(number.Small) 
	} 
	return number.Large.String() 
}

`

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

func (l *implementation) Name() string { return Name }

func (l *implementation) GoTypeOf(T language.Type) string {
	switch T.(type) {
		case language.Number:
			l.Import(NumberImport)
			l.AddHelper(NumberTypeDefinition)
			return "Number"

		case language.Symbol:
			return "rune"
			
		case language.String:
			return "string"
			
		case language.FunctionType:
			
			var result = "func("
			
			var args = T.(language.FunctionType).Arguments()
			for i, t := range args {
				result += l.GoTypeOf(t)
				
				if i < len(args)-1 {
					result += ","
				}
			}
			
			//TODO return types.
			
			return result + ")"
			
		default:
			panic("Not corresponding Go type for "+T.Name())
	}
}


//Returns a Commands that builds go code sitting at path.
func (l *implementation) Build(path string) *exec.Cmd {
	return exec.Command("go", "build", path)
}

func (l *implementation) GetExpression(T language.Type) string {
	if T == nil {
		return ""
	}
	
	var PanicName = "Error in "+Name+".GetExpression("+T.Name()+")"
	
	switch T.(type) {
		case language.String:
			return string(T.(String))
			
		case language.Symbol:
			return string(T.(Symbol))
			
		case language.Boolean:
			return string(fmt.Sprint(T))
			
		case language.Number:
			n := T.(Number)
			
			if n.Literal == nil {
				return n.Expression
			} else {
				
				l.Import(NumberImport)
				l.AddHelper(NumberTypeDefinition)
				
				if n.Literal.IsInt64() {
					return "Number{Small:"+n.Literal.String()+"}"
				}
				
				l.AddHelper(NumberTypeDefinition)
				return "Number{Large: new(big.Int).SetBytes("+strconv.Quote(string(n.Literal.Bytes()))+")}"
			}
			
		case language.Array:
			a := T.(Array)
			return a.Expression
		
		case language.Function:
			return T.(Function).Expression
		
		case language.Custom, language.Stream, language.List,
			language.Table, language.Error, language.Float, language.Pointer, 
			language.Dynamic, language.Metatype, language.FunctionType:
		
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
