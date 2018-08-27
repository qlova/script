package Go

import "github.com/qlova/script/language"


//Returns Error cast to String.
func (l *implementation) ErrorToString(language.Error) language.String {
	panic("Error in "+Name+".ErrorToString(Error): Unimplemented")
	return nil
}
		
//Returns Error cast to Number.
func (l *implementation) ErrorToNumber(language.Error) language.Number {
	panic("Error in "+Name+".ErrorToNumber(Error): Unimplemented")
	return nil
}
		
//Returns Number cast to String.
func (l *implementation) NumberToString(language.Number) language.String {
	panic("Error in "+Name+".NumberToString(Number): Unimplemented")
	return nil
}

//Returns String cast to Number.
func (l *implementation) StringToNumber(language.String) language.Number {
	panic("Error in "+Name+".StringToNumber(String): Unimplemented")
	return nil
}

//Returns Symbol cast to String.
func (l *implementation) SymbolToString(language.Symbol) language.String {
	panic("Error in "+Name+".SymbolToString(Symbol): Unimplemented")
	return nil
}

//Returns Symbol cast to Number.
func (l *implementation) SymbolToNumber(language.Symbol) language.Number {
	panic("Error in "+Name+".SymbolToNumber(Symbol): Unimplemented")
	return nil
}

//Returns Symbol cast to Number.
func (l *implementation) NumberToSymbol(language.Number) language.Symbol {
	panic("Error in "+Name+".NumberToSymbol(Number): Unimplemented")
	return nil
}
