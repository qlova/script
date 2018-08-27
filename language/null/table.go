package Null

import . "github.com/qlova/script/language"

//Returns a Table of type T.
func (l *language) MakeTable(T Type) Table {
	panic("Error in "+Name+".MakeTable(Type): Unimplemented")
	return nil
}

//Returns a Table intialised with 'indices' corresponding to 'elements'.
func (l *language) LiteralTable(indices []String, elements []Type) Table {
	panic("Error in "+Name+".LiteralTable([]String, []Type): Unimplemented")
	return nil
}
