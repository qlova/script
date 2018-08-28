package Null

import "math/big"
import "github.com/qlova/script/language"

//Returns a Number based on the passed big Integer.
func (l *implementation) LiteralNumber(literal *big.Int) language.Number {
	panic("Error in "+Name+".LiteralNumber(Number): Unimplemented")
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
