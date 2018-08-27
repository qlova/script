package Null

import . "github.com/qlova/script/language"


//Returns Error cast to String.
func (l *language) ErrorToString(Error) String {
	panic("Error in "+Name+".ErrorToString(Error): Unimplemented")
	return nil
}
		
//Returns Error cast to Number.
func (l *language) ErrorToNumber(Error) Number {
	panic("Error in "+Name+".ErrorToNumber(Error): Unimplemented")
	return nil
}
		
//Returns Number cast to String.
func (l *language) NumberToString(Number) String {
	panic("Error in "+Name+".NumberToString(Number): Unimplemented")
	return nil
}

//Returns String cast to Number.
func (l *language) StringToNumber(String) Number {
	panic("Error in "+Name+".StringToNumber(String): Unimplemented")
	return nil
}

//Returns Symbol cast to String.
func (l *language) SymbolToString(Symbol) String {
	panic("Error in "+Name+".SymbolToString(Symbol): Unimplemented")
	return nil
}

//Returns Symbol cast to Number.
func (l *language) SymbolToNumber(Symbol) Number {
	panic("Error in "+Name+".SymbolToNumber(Symbol): Unimplemented")
	return nil
}

//Returns Symbol cast to Number.
func (l *language) NumberToSymbol(Number) Symbol {
	panic("Error in "+Name+".NumberToSymbol(Number): Unimplemented")
	return nil
}
