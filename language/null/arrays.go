package Null

import "github.com/qlova/script/language"

//Returns a new String that concatenates 'a' and 'b'.
func (l *implementation) Join(a, b language.Type) language.Type {
	panic("Error in "+Name+".Join("+a.Name()+", "+b.Name()+"): Unimplemented")
	return nil
}


//Returns an Array of type T with 'length' length.
func (l *implementation) Literal(array interface{}) language.Type {
	panic("Error in "+Name+".Literal(Type, Number): Unimplemented")
	return nil
}

//Returns a Number representing the length of 'array'
func (l *implementation) Array(T language.Type, length language.Number) language.Array {
	panic("Error in "+Name+".Array("+T.Name()+", Number): Unimplemented")
	return nil
}


//Returns a Number representing the length of 'array'
func (l *implementation) Fill(T language.Type, elements []language.Type) language.Type {
	
	panic("Error in "+Name+".Fill("+T.Name()+", []language.Type): Unimplemented")
	return nil
}


//Returns a List of type T.
func (l *implementation) MakeList(T language.Type) language.List {
	panic("Error in "+Name+".MakeList(Type): Unimplemented")
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
