package Null

import . "github.com/qlova/script/language"

//Returns an Error with 'code' and 'message'.
func (l *language) LiteralError(code Number, message String) Error {
	panic("Error in "+Name+".LiteralError(Number, String): Unimplemented")
	return nil
}

//Returns a Statement that throws an error to the thread, this should not halt the program.
func (l *language) Throw(err Error) Statement {
	panic("Error in "+Name+".Throw(String, String): Unimplemented")
	return ""
}

//Returns a Error that is sitting on the thread.
func (l *language) Catch() Error {
	panic("Error in "+Name+".Catch(): Unimplemented")
	return nil
}
