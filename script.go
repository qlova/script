package script

import "bytes"

type Script struct {
	Depth int
	
	Language Language
	
	Head bytes.Buffer
	Neck bytes.Buffer
	Body bytes.Buffer
	Tail bytes.Buffer
	
	FakeScript bool
}

func NewScript() *Script {
	s := new(Script)
	return s
}

func (q *Script) Fake() *Script {
	s := new(Script)
	s.Language = q.Language
	s.Depth = q.Depth
	s.FakeScript = true
	return s
}

func (q *Script) String() string {
	return string(q.Head.Bytes())+string(q.Neck.Bytes())+string(q.Body.Bytes())+string(q.Tail.Bytes())
}

func (q *Script) SetLanguage(f Language) {
	q.Language = f
}

func (q *Script) Init() {
	if q != nil {
		q.Language.Init(q)
	}
}

func (q *Script) Last() {
	if q != nil {
		q.Language.Last(q)
	}
}

func (q *Script) Main(f func(*Script)) {
	q.Language.Main(q)
	q.Depth++
	f(q)
	q.Depth--
	q.Language.EndMain(q)
}

func (q *Script) ForEachList(i string, variable string, list List, f func(*Script)) {
	q.IndentBody()
	q.Body.WriteString(q.Language.ForEachList(i, variable, list.String()))
	q.Depth++
	f(q)
	q.Depth--
	q.IndentBody()
	q.Body.WriteString(q.Language.EndForEachList())
}

func (q *Script) Return(t Type) {
	q.IndentBody()
	q.Language.Return(q, t)
}

func (q *Script) Function(name string, f func(*Script), arguments []Type, returns []Type) {
	var s = NewScript()
	s.SetLanguage(q.Language)
	
	q.Language.Function(s, name, arguments, returns)
	s.Depth++
	f(s)
	s.Depth--
	q.Language.EndFunction(s)
	
	q.Head.Write(s.Head.Bytes())
	q.Neck.Write(s.Neck.Bytes())
	q.Neck.Write(s.Body.Bytes())
}

func (q *Script) Call(name string, arguments []Type, returns Type) Type {
	return q.Raw(returns, q.Language.Call(q, name, arguments))
}

func (q *Script) Run(name string, arguments []Type) {
	q.IndentBody()
	q.Body.WriteString(q.Language.Call(q, name, arguments))
	q.Body.WriteByte('\n')
}


func (q *Script) ScopeFunction(name string, arguments []Type) {
	q.IndentBody()
	q.Language.Call(q, name, arguments)
}

func (q *Script) IndentBody() {
	for i := 0; i < q.Depth; i++ {
		q.Body.WriteByte('\t')
	}
}

func (q *Script) Print(value String) {
	q.IndentBody()
	q.Language.Print(q, value)
}

func (q *Script) Raw(T Type, name string) Type {
	switch T.(type) {
		case String:
			return String(name)
			
		case Number:
			return Number(name)
			
		case Symbol:
			return Symbol(name)
			
		case Array:
			return Array{subtype: T.(Array).subtype, data:name}
			
		case List:
			return List{subtype: T.(List).subtype, data:name}
			
		case FunctionType:
			return FunctionType{rets: T.(FunctionType).rets, args: T.(FunctionType).args, data:name}
		
		default:
			panic("ERROR invalid call to qlova.Script.Literal("+T.Name()+", "+name+")")
	}
	return nil
}

func (q *Script) Set(T interface{}, name string) {
	switch T.(type) {
		case String:
			q.SetString(name, T.(String))
			
		case Number:
			q.SetNumber(name, T.(Number))
		
		default:
			panic("ERROR invalid call to qlova.Script.Set(T, name)")
	}
}

func (q *Script) New(T Type, name string) {
	switch T.(type) {
		case String:
			q.NewString(name, T.(String))
			
		case Number:
			q.NewNumber(name, T.(Number))
	
		case Symbol:
			q.NewSymbol(name, T.(Symbol))
		
		case Array:
			q.NewArray(T.(Array).subtype, name, T.(Array))

		case List:
			q.NewList(T.(List).subtype, name, T.(List))
			
		case FunctionType:
			q.NewFunctionType(name, T.(FunctionType).args, T.(FunctionType).rets, T.(FunctionType))
			
		default:
			panic("ERROR invalid call to qlova.Script.New("+T.Name()+", "+name+")")
	}
}
