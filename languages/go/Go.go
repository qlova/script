package Go

import qlova "github.com/qlova/script"

import "strconv"

type language struct {
	imports map[string]bool
	importsList []string
	
	helpers map[string]bool
	helpersList []string
}

func Language() *language {
	return &language{
		imports: make(map[string]bool),
		helpers: make(map[string]bool),
	}
}

func (l *language) GoTypeOf(T qlova.Type) string {
	switch T.(type) {
		case qlova.Number:
			return "*big.Int"

		case qlova.Symbol:
			return "rune"
			
		case qlova.String:
			return "string"
			
		case qlova.FunctionType:
			
			var result = "func("
			
			var args = T.(qlova.FunctionType).Arguments()
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

func (l *language) Init(q *qlova.Script) {
	q.Head.WriteString("package main\n")
}

func (l *language) Last(q *qlova.Script) {
	for i := range l.importsList {
		l.Import(q, l.importsList[i])
	}
	for i := range l.helpersList {
		l.Helper(q, l.helpersList[i])
	}
}

func New(T string, q *qlova.Script, name, value string) {
	q.Body.WriteString("var ")
	q.Body.WriteString(name)
	
	if len(value) == 0 {
		q.Body.WriteString(" ")
		q.Body.WriteString(T)
	} else {
		q.Body.WriteString(" = ")
		q.Body.WriteString(value)
		q.Body.WriteString("\n")
	}
}

func Set(q *qlova.Script, name, value string) {
	q.Body.WriteString(name)
	q.Body.WriteString(" = ")
	q.Body.WriteString(value)
	q.Body.WriteString("\n")
}

func (l *language) Import(q *qlova.Script, pkg string) {
	if !l.imports[pkg] && !q.FakeScript {
		q.Head.WriteString(`import "`)
		q.Head.WriteString(pkg)
		q.Head.WriteString("\"\n")
		l.imports[pkg] = true
		
		if pkg == "math/big" {
			q.Neck.WriteString(`func Number(data string) *big.Int {
	i, _ := big.NewInt(0).SetString(data, 10)
	return i
}
`)
		}
	}
}

func (l *language) Helper(q *qlova.Script, helper string) {
	if !l.helpers[helper] {
		q.Neck.WriteString(helper)
		l.helpers[helper] = true
	}
}


func (l *language) Main(q *qlova.Script) {
	q.Body.WriteString("func main() {\n")
}

func (l *language) EndMain(q *qlova.Script) {
	q.Body.WriteString("}\n")
}

