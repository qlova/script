package Null

import "github.com/qlova/script/language"

//Returns an Error with 'code' and 'message'.
func (l *implementation) LiteralError(code language.Number, message language.String) language.Error {
	panic("Error in "+Name+".LiteralError(Number, String): Unimplemented")
	return nil
}

//Returns a Statement that throws an error to the thread, this should not halt the program.
func (l *implementation) Throw(err language.Error) language.Statement {
	panic("Error in "+Name+".Throw(String, String): Unimplemented")
	return ""
}

//Returns a Error that is sitting on the thread.
func (l *implementation) Catch() language.Error {
	panic("Error in "+Name+".Catch(): Unimplemented")
	return nil
}
