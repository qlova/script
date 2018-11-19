package script

import "github.com/qlova/script/language"

type EmbeddedScript struct {
	q Script
}
func (e EmbeddedScript) script() Script { return e.q }

type Type interface {
	language.Type
	script() Script
}

func (q Script) Wrap(t language.Type) Type {
	return q.wrap(t)
}

func (q Script) wrap(t language.Type) Type {
	switch value := t.(type) {
		case language.String:
			return String{String: value, EmbeddedScript:EmbeddedScript{q:q}}
			
		case language.Number:
			return Number{Number: value, EmbeddedScript:EmbeddedScript{q:q}}
			
		case language.Symbol:
			return Symbol{Symbol: value, EmbeddedScript:EmbeddedScript{q:q}}
		
		case language.Function:
			return Function{Function: value, EmbeddedScript:EmbeddedScript{q:q}}
		
		case language.List:
			return List{List: value, EmbeddedScript:EmbeddedScript{q:q}}
		
		case language.Array:
			return Array{Array: value, EmbeddedScript:EmbeddedScript{q:q}}
		
		case language.Boolean:
			return Boolean{Boolean: value, EmbeddedScript:EmbeddedScript{q:q}}
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
				
			case Error:
				return value.Error
				
			case Boolean:
				if value.Literal == nil {
					return value.Boolean
				} else {
					return value.q.lang.Literal(*value.Literal)
				}
				
			case Number:
				if value.Literal == nil {
					return value.Number
				} else {
					return value.q.lang.Literal(value.Literal)
				}
			
			case Array:
				if value.Literal == nil {
					return value.Array
				} else {
					//Special case for arrays.
					var elements = make([]language.Type, len(value.Literal))

					for i := 0; i < len(value.Literal); i++ {
						elements[i] = convert(value.Literal[i])
					}

					return value.q.lang.Fill(value.q.lang.Array(elements[0], len(value.Literal)), elements)
				}
				
			case List:
				if value.Literal == nil {
					return value.List
				} else {
					panic("Unimplemented")
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
