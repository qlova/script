package Go

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

func (s String) Raw() string {
	return string(s)
}

func (String) String() {}

type Symbol string

func (Symbol) Name() string {
	return "symbol"
}

func (Symbol) SameAs(i interface{}) bool {
	_, ok := i.(Symbol)
	return ok
}

func (Symbol) Symbol() {}
