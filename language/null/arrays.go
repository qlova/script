package Null

import . "github.com/qlova/script/language"


//Returns an Array of type T with 'length' length.
func (l *language) MakeArray(T Type, length Number) Array {
	panic("Error in "+Name+".MakeArray(Type, Number): Unimplemented")
	return nil
}

//Returns a Number representing the length of 'array'
func (l *language) LengthArray(array Array) Number {
	panic("Error in "+Name+".LengthArray(Array): Unimplemented")
	return nil
}

//Returns a List of type T.
func (l *language) MakeList(T Type) List {
	panic("Error in "+Name+".MakeList(Type): Unimplemented")
	return nil
}

//Returns a List of type T, intialised with optional 'elements'.
func (l *language) LiteralList(T Type, elements ...Type) List {
	panic("Error in "+Name+".LiteralList(Type, ...Type): Unimplemented")
	return nil
}

//Returns a Number representing the length of 'list'
func (l *language) LengthList(list List) Number {
	panic("Error in "+Name+".LengthList(List): Unimplemented")
	return nil
}

//Returns a statement that adds 't' to the 'list'.
func (l *language) Grow(list List, t Type) Statement {
	panic("Error in "+Name+".Grow(List, Type): Unimplemented")
	return ""
}

//Returns a statement that shrinks the 'list' by removing the element at the end..
func (l *language) Shrink(list List) Statement {
	panic("Error in "+Name+".Shrink(List): Unimplemented")
	return ""
}
