package Interpreter

import "github.com/qlova/script/language"

//Returns a Statement that throws an error to the thread, this should not halt the program.
func (l *implementation) Error() language.Error {
	panic(Name+".Error() Not Implemented")
	return nil
}

//Returns a Statement that throws an error to the thread, this should not halt the program.
func (l *implementation) Trace(line int, filename string) language.Statement {
	panic(Name+".Trace() Not Implemented")
	return ""
}

//Returns a Statement that throws an error to the thread, this should not halt the program.
func (l *implementation) Throw(code language.Number, message language.String) language.Statement {
	panic(Name+".Throw() Not Implemented")
	return ""
}

//Returns a Error that is sitting on the thread.
func (l *implementation) Catch() language.Boolean {
	panic(Name+".Catch() Not Implemented")
	return nil
}
