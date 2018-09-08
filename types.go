package script

import "github.com/qlova/script/language"

type EmbeddedScript struct {
	q *Script
}
func (e EmbeddedScript) script() *Script { return e.q }

type Type interface {
	language.Type
	script() *Script
}

func (q *Script) wrap(t language.Type) Type {
	switch value := t.(type) {
		case language.String:
			return String{String: value, EmbeddedScript:EmbeddedScript{q:q}}
			
		case language.Number:
			return Number{Number: value, EmbeddedScript:EmbeddedScript{q:q}}
			
		case language.Symbol:
			return Symbol{Symbol: value, EmbeddedScript:EmbeddedScript{q:q}}
		
		case language.Function:
			return Function{Function: value, EmbeddedScript:EmbeddedScript{q:q}}
	}
	
	panic("Cannot wrap type "+t.Name())
}

func convert(l Type) language.Type {
	
	if t, ok := l.(Type); ok {
		switch value := t.(type) {
			case String:
				if value.Literal == nil {
					return value.String
				} else {
					return value.q.lang.Literal(*value.Literal)
				}
			case Symbol:
				if value.Literal == nil {
					return value.Symbol
				} else {
					return value.q.lang.Literal(*value.Literal)
				}
			case Number:
				if value.Literal == nil {
					return value.Number
				} else {
					return value.q.lang.Literal(value.Literal)
				}
			case Function:
				if value.Literal == nil {
					return value.Function
				} else {
					panic("Unimplemented")
					return value.q.lang.Literal(value.Literal)
				}
		}
	}

	panic("Cannot convert type "+l.Name())
}
