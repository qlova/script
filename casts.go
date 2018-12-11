package script

func (q Script) ToString(T Type) string {
	
	switch value := T.(type) {
		case string:
			return value
		
		case Number:
			if value.Literal != nil {
				var cast = value.Literal.String()
				return string{EmbeddedScript: EmbeddedScript{q:q}, literal: &cast}
			}
	}
	
	return q.wrap(q.lang.ToString(convert(T))).(string)
}

func (q Script) ToNumber(T Type) Number {
	return q.wrap(q.lang.ToNumber(convert(T))).(Number)
}

func (q Script) ToSymbol(T Type) Symbol {
	return q.wrap(q.lang.ToSymbol(convert(T))).(Symbol)
}
