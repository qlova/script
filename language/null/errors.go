package Null

import "github.com/qlova/script/language"


//Returns a Statement that embeds an error within another error.
func (l *implementation) Embed(a, b language.Error) language.Statement {
	panic("Error in "+Name+".Embed(Error, Error): Unimplemented")
	return ""
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
