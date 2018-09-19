package Go

import "github.com/qlova/script/language"

//Returns Type cast to String.
func (l *implementation) ToString(T language.Type) language.String {
	switch T.(type) {
		case String:
			return T.(String)

		case Number:
			l.AddHelper(NumberTypeDefinition)
			l.AddHelper(NumberToString)
			return String(l.GetExpression(T)+".String()")
		
		case Symbol:
			return String("string("+l.GetExpression(T)+")")
			
		case Error:
			l.Import("fmt")
			l.AddHelper(ErrorToString)
			return String(l.GetExpression(T)+".String()")
	}
	panic("Error in "+Name+".ToString("+T.Name()+"): Unimplemented")
	return nil
}
		
//Returns Type cast to Number.
func (l *implementation) ToNumber(T language.Type) language.Number {
	
	var result Number
	
	switch T.(type) {
		case Number:
			return T.(Number)
			
		case String:
			
			//TODO throw an error?
			l.Import(NumberImport)
			l.AddHelper(NumberTypeDefinition)
			l.AddHelper(`func ParseNumber(number string) Number {
	var z big.Int
	z.SetString(number, 10)
	if z.IsInt64() {
		return Number{Small: z.Int64()}
	}
	return Number{Large: &z}
}

`)
				result.Expression = ("ParseNumber("+l.GetExpression(T)+")")
				return result
	
		case Symbol:
			result.Expression = "Number{Small:int64("+l.GetExpression(T)+")}"
			return result
	}
	
	panic("Error in "+Name+".ToNumber("+T.Name()+"): Unimplemented")
	return nil
}
		
//Returns Type cast to Number.
func (l *implementation) ToSymbol(T language.Type) language.Symbol {

	switch T.(type) {
		case Symbol:
			return T.(Symbol)
	
		case Number:
			l.AddHelper(NumberTypeDefinition)
			l.AddHelper(NumberToInt64)
			return Symbol("rune("+l.GetExpression(T)+".Int64())")
	}
	
	panic("Error in "+Name+".ToSymbol("+T.Name()+"): Unimplemented")
	return nil
}


//Returns Error cast to String.
func (l *implementation) ErrorToString(err language.Error) language.String {
	return String(l.GetExpression(err)+".Message")
}
		
//Returns Error cast to Number.
func (l *implementation) ErrorToNumber(err language.Error) language.Number {
	var result Number
	result.Expression = l.GetExpression(err)+".Code"
	return result
}
