package Go

import "github.com/qlova/script/language"


//Returns an Array of type T with 'length' length.
func (l *implementation) MakeArray(T language.Type, length language.Number) language.Array {
	panic("Error in "+Name+".MakeArray(Type, Number): Unimplemented")
	return nil
}

//Returns a Number representing the length of 'array'
func (l *implementation) LengthArray(array language.Array) language.Number {
	panic("Error in "+Name+".LengthArray(Array): Unimplemented")
	return nil
}

//Returns a List of type T.
func (l *implementation) MakeList(T language.Type) language.List {
	panic("Error in "+Name+".MakeList(Type): Unimplemented")
	return nil
}

//Returns a List of type T, intialised with optional 'elements'.
func (l *implementation) LiteralList(T language.Type, elements ...language.Type) language.List {
	panic("Error in "+Name+".LiteralList(Type, ...Type): Unimplemented")
	return nil
}

//Returns a Number representing the length of 'list'
func (l *implementation) LengthList(list language.List) language.Number {
	panic("Error in "+Name+".LengthList(List): Unimplemented")
	return nil
}

//Returns a statement that adds 't' to the 'list'.
func (l *implementation) Grow(list language.List, t language.Type) language.Statement {
	panic("Error in "+Name+".Grow(List, Type): Unimplemented")
	return ""
}

//Returns a statement that shrinks the 'list' by removing the element at the end..
func (l *implementation) Shrink(list language.List) language.Statement {
	panic("Error in "+Name+".Shrink(List): Unimplemented")
	return ""
}
