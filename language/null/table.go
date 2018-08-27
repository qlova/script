package Null

import "github.com/qlova/script/language"

//Returns a Table of type T.
func (l *implementation) MakeTable(T language.Type) language.Table {
	panic("Error in "+Name+".MakeTable(Type): Unimplemented")
	return nil
}

//Returns a Table intialised with 'indices' corresponding to 'elements'.
func (l *implementation) LiteralTable(indices []language.String, elements []language.Type) language.Table {
	panic("Error in "+Name+".LiteralTable([]String, []Type): Unimplemented")
	return nil
}
