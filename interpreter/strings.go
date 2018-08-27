package Interpreter

import "github.com/qlova/script/interpreter/internal"
import "github.com/qlova/script/language"

//Returns a Symbol that the Go style literal represents ('').
func (l *implementation) LiteralSymbol(literal string) language.Symbol {
	panic("Error in "+Name+".LiteralSymbol("+literal+"): Unimplemented")
	return nil
}

type String struct {	
	internal.Variable
	
	IsLiteral bool
	Literal string
}

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
	return String{Literal: literal, IsLiteral: true}
}


//Returns a new String that concatenates 'a' and 'b'.
func (l *implementation) JoinString(a, b language.String) language.String {
	panic("Error in "+Name+".JoinString(String, String): Unimplemented")
	return nil
}
