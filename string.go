package script

import "errors"

type String string

func (s String) String() string {
	return string(s)
}

func (q Script) AddStrings(a, b String) String {
	return String(q.Language.AddStrings(a.String(), b.String()))
}

func (q Script) LiteralString(s string) String {
	return String(s)
}

func (q *Script) NewString(name string, value String) {
	q.IndentBody()
	q.Language.NewString(q, name, value)
}

func (q *Script) SetString(name string, value String) {
	q.IndentBody()
	q.Language.SetString(q, name, value)
}

func (n String) Name() string {
	return "string"
}

func (n String) Equals(t interface{}) bool {
	if _, ok := t.(String); ok {
		return true
	}
	return false
}

func (q *Script) ToString(T Type) (String, error) {
	switch T.(type) {
		
		case Number:
			return q.LiteralString(q.Language.NumberToString(string(T.(Number)))), nil
			
		case Symbol:
			return q.LiteralString(q.Language.SymbolToString(string(T.(Symbol)))), nil
		
		case String:
			return T.(String), nil
		
		default:
			return "", errors.New("Cannot convert "+T.Name()+" to string!")
	}
}
