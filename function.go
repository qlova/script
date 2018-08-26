package script

//import "errors"

type FunctionType struct {
	data string
	args []Type
	rets []Type
}

func (s FunctionType) String() string {
	return s.data
}

func (s FunctionType) Arguments() []Type {
	return s.args
}

/*func (q Script) LiteralFunctionType(s FunctionType) FunctionType {
	return FunctionType(s)
}*/

func (q *Script) ScopeFunctionType(name string, args []Type, rets []Type) FunctionType {
	return FunctionType{data:q.Language.ScopeFunctionType(name), args: args, rets: rets}
}

func (q *Script) NewFunctionType(name string, args []Type, rets []Type, value FunctionType) {
	q.IndentBody()
	q.Language.NewFunctionType(q, name, args, rets, value)
}

func (q *Script) RunFunctionType(name string, args []Type) {
	q.IndentBody()
	q.Body.Write([]byte(q.Language.CallFunctionType(q, name, args)))
	q.Body.WriteByte('\n')
}

func (q *Script) CallFunctionType(name string, args []Type) string {
	return q.Language.CallFunctionType(q, name, args)
}

func (q *Script) SetFunctionType(name string, value FunctionType) {
	q.IndentBody()
	q.Language.SetFunctionType(q, name, value)
}

func (n FunctionType) Name() string {
	return "function"
}

func (n FunctionType) Equals(t interface{}) bool {
	if _, ok := t.(FunctionType); ok {
		return true
	}
	return false
}
