package Null

import "github.com/qlova/script/language"


//Returns Type cast to String.
func (l *implementation) ToString(T language.Type) language.String {
	panic("Error in "+Name+".ToString("+T.Name()+"): Unimplemented")
	return nil
}
		
//Returns Type cast to Number.
func (l *implementation) ToNumber(T language.Type) language.Number {
	panic("Error in "+Name+".ToNumber("+T.Name()+"): Unimplemented")
	return nil
}
		
//Returns Type cast to Number.
func (l *implementation) ToSymbol(T language.Type) language.Symbol {
	panic("Error in "+Name+".ToSymbol("+T.Name()+"): Unimplemented")
	return nil
}
