package Null

import . "github.com/qlova/script/language"


//Returns a Number that the Go style literal represents (01 1 0x1).
func (l *language) LiteralNumber(literal string) Number {
	panic("Error in "+Name+".LiteralSwitch("+literal+"): Unimplemented")
	return nil
}

//Returns a Number that is the sum of 'a' and 'b'.
func (l *language) Add(a, b Number) Number {
	panic("Error in "+Name+".Add(Number, Number): Unimplemented")
	return nil
}

//Returns a Number that is the difference of 'a' and 'b'.
func (l *language) Sub(a, b Number) Number {
	panic("Error in "+Name+".Sub(Number, Number): Unimplemented")
	return nil
}

//Returns a Number that is the product of 'a' and 'b'.
func (l *language) Mul(a, b Number) Number {
	panic("Error in "+Name+".Mul(Number, Number): Unimplemented")
	return nil
}

//Returns a Number that is the quotient of 'a' and 'b'.
func (l *language) Div(a, b Number) Number {
	panic("Error in "+Name+".Div(Number, Number): Unimplemented")
	return nil
}

//Returns a Number that is 'a' taken to the power of 'b'.
func (l *language) Pow(a, b Number) Number {
	panic("Error in "+Name+".Pow(Number, Number): Unimplemented")
	return nil
}

//Returns a Number that is modulos of 'a' and 'b'.
func (l *language) Mod(a, b Number) Number {
	panic("Error in "+Name+".Mod(Number, Number): Unimplemented")
	return nil
}
