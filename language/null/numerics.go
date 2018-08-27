package Null

import "github.com/qlova/script/language"


//Returns a Number that the Go style literal represents (01 1 0x1).
func (l *implementation) LiteralNumber(literal string) language.Number {
	panic("Error in "+Name+".LiteralSwitch("+literal+"): Unimplemented")
	return nil
}

//Returns a Number that is the sum of 'a' and 'b'.
func (l *implementation) Add(a, b language.Number) language.Number {
	panic("Error in "+Name+".Add(Number, Number): Unimplemented")
	return nil
}

//Returns a Number that is the difference of 'a' and 'b'.
func (l *implementation) Sub(a, b language.Number) language.Number {
	panic("Error in "+Name+".Sub(Number, Number): Unimplemented")
	return nil
}

//Returns a Number that is the product of 'a' and 'b'.
func (l *implementation) Mul(a, b language.Number) language.Number {
	panic("Error in "+Name+".Mul(Number, Number): Unimplemented")
	return nil
}

//Returns a Number that is the quotient of 'a' and 'b'.
func (l *implementation) Div(a, b language.Number) language.Number {
	panic("Error in "+Name+".Div(Number, Number): Unimplemented")
	return nil
}

//Returns a Number that is 'a' taken to the power of 'b'.
func (l *implementation) Pow(a, b language.Number) language.Number {
	panic("Error in "+Name+".Pow(Number, Number): Unimplemented")
	return nil
}

//Returns a Number that is modulos of 'a' and 'b'.
func (l *implementation) Mod(a, b language.Number) language.Number {
	panic("Error in "+Name+".Mod(Number, Number): Unimplemented")
	return nil
}
