package Go

import "github.com/qlova/script/language"

type Error struct {
	language.ErrorType
}

//Returns a Statement that embeds an error within another error.
func (l *implementation) Embed(a, b language.Error) language.Statement {
	l.AddHelper(`var Errors []*Error
type Error struct {
	Code int
	Message string
	Embedded *Error
}
func Throw(err *Error) *Error {
	Errors = append(Errors, err)
}
func Catch() *Error {
	var err = Errors[len(Errors)-1]
	Errors = Errors[:len(Errors)-1]
	return err
}
`)
	
	return language.Statement(l.GetExpression(a)+".Embedded = "+l.GetExpression(b))
}

//Returns a Statement that throws an error to the thread, this should not halt the program.
func (l *implementation) Throw(err language.Error) language.Statement {
	return language.Statement("Throw("+l.GetExpression(err)+")")
}

//Returns a Error that is sitting on the thread.
func (l *implementation) Catch() language.Error {
	var result Error
	result.Expression = "Catch()"
	return result
}
