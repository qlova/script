package Go

import "strconv"
import "github.com/qlova/script/language"

//Returns a Symbol that the Go style literal represents ('').
func (l *implementation) LiteralSymbol(literal string) language.Symbol {
	panic("Error in "+Name+".LiteralSymbol("+literal+"): Unimplemented")
	return nil
}

type String string

func (String) Name() string {
	return "string"
}

func (String) SameAs(i interface{}) bool {
	_, ok := i.(String)
	return ok
}

func (String) String() {}

//Returns a String that the Go style literal represents ("").
func (l *implementation) LiteralString(literal string) language.String {
	return String(strconv.Quote(literal))
}

//Returns a new String that concatenates 'a' and 'b'.
func (l *implementation) JoinString(a, b language.String) language.String {
	panic("Error in "+Name+".JoinString(String, String): Unimplemented")
	return nil
}
