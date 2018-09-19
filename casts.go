package script

func (q *Script) ToString(T Type) String {
	
	switch value := T.(type) {
		case String:
			return value
		
		case Number:
			if value.Literal != nil {
				var cast = value.Literal.String()
				return String{EmbeddedScript: EmbeddedScript{q:q}, Literal: &cast}
			}
	}
	
	return q.wrap(q.lang.ToString(convert(T))).(String)
}

func (q *Script) ToNumber(T Type) Number {
	return q.wrap(q.lang.ToNumber(convert(T))).(Number)
}

func (q *Script) ToSymbol(T Type) Symbol {
	return q.wrap(q.lang.ToSymbol(convert(T))).(Symbol)
}
