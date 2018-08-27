package Null

import "github.com/qlova/script/language"

//Returns a Symbol that the Go style literal represents ('').
func (l *implementation) LiteralSymbol(literal string) language.Symbol {
	panic("Error in "+Name+".LiteralSymbol("+literal+"): Unimplemented")
	return nil
}

//Returns a String that the Go style literal represents ("").
func (l *implementation) LiteralString(literal string) language.String {
	panic("Error in "+Name+".LiteralString("+literal+"): Unimplemented")
	return nil
}

//Returns a new String that concatenates 'a' and 'b'.
func (l *implementation) JoinString(a, b language.String) language.String {
	panic("Error in "+Name+".JoinString(String, String): Unimplemented")
	return nil
}
