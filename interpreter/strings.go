package Interpreter

import "github.com/qlova/script/interpreter/internal"

type String struct {	
	internal.Variable
	
	Address internal.StringAddress
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

type Symbol struct {	
	internal.Variable
	
	Address internal.SymbolAddress
	IsLiteral bool
	Literal rune
}

func (Symbol) Name() string {
	return "symbol"
}

func (Symbol) SameAs(i interface{}) bool {
	_, ok := i.(Symbol)
	return ok
}

func (Symbol) Symbol() {}
