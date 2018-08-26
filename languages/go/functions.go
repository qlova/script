package Go
 
import qlova "github.com/qlova/script" 

func (l *language) Call(q *qlova.Script, name string, arguments []qlova.Type) string {
	var result = name + "("
	for i := range arguments {
		result += arguments[i].String()
		
		if i < len(arguments) - 1 {
			result +=  ","
		}
	}
	result += ")"
	
	return result
}

func (l *language) Function(q *qlova.Script, name string, arguments []qlova.Type, returns []qlova.Type) {
	q.Body.WriteString("func ")
	q.Body.WriteString(name)
	q.Body.WriteString("(")
	for i := range arguments {
		q.Body.WriteString(arguments[i].String())
		q.Body.WriteString(" ")
		q.Body.WriteString(l.GoTypeOf(arguments[i]))
		
		if i < len(arguments) - 1 {
			q.Body.WriteString(",")
		}
	}
	q.Body.WriteString(") (")
	for i := range returns {
		q.Body.WriteString(l.GoTypeOf(arguments[i]))
		
		if i < len(returns) - 1 {
			q.Body.WriteString(",")
		}
	}
	q.Body.WriteString(") {\n")
	
	q.Neck, q.Head = q.Head, q.Neck
}

func (l *language) EndFunction(q *qlova.Script) {
	q.Body.WriteString("}\n")
}

func (l *language) Return(q *qlova.Script, t qlova.Type) {
	q.Body.WriteString("return ")
	q.Body.WriteString(t.String())
	q.Body.WriteString("\n")
}

func (l *language) ScopeFunctionType(name string) string {
	return name
}

func (l *language) NewFunctionType(q *qlova.Script, name string, args []qlova.Type, ret []qlova.Type, value ...qlova.FunctionType) {
	if args == nil && ret == nil {
		New("func()", q, name, value[0].String())
		return
	}
	panic("unimplemented")
}

func (l *language) SetFunctionType(q *qlova.Script, name string, value qlova.FunctionType) {
	Set(q, name, value.String())
}

func (l *language) CallFunctionType(q *qlova.Script, name string, args []qlova.Type) string {
	if args == nil {
		return name+"()"
	}
	panic("unimplemented")
}
