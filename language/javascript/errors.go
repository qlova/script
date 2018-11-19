package Javascript

import "fmt"
import "strconv"
import "github.com/qlova/script/language"

type Error struct {
	language.ErrorType
}

const ErrorHelper = `type Error struct {
	Code, Line int
	Message, File string
}

var err Error
var errors []Error
var line int
var filename string

`

const ErrorToString = `func (err Error) String() string {
	return err.File+":"+fmt.Sprint(err.Line)+" "+err.Message
}

`

//Returns a Statement that throws an error to the thread, this should not halt the program.
func (l *implementation) Error() language.Error {
	l.AddHelper(ErrorHelper)
	
	var e Error
	e.Expression = "err"
	return e
}

//Returns a Statement that throws an error to the thread, this should not halt the program.
func (l *implementation) Trace(line int, filename string) language.Statement {
	l.AddHelper(ErrorHelper)

	return language.Statement("line, filename = "+fmt.Sprint(line)+", "+strconv.Quote(filename)+"\n")
}

//Returns a Statement that throws an error to the thread, this should not halt the program.
func (l *implementation) Throw(code language.Number, message language.String) language.Statement {
	l.AddHelper(ErrorHelper)
	
	return language.Statement("errors = append(errors, Error{Line:line, File:filename, Code:"+l.GetExpression(code)+", Message:"+l.GetExpression(message)+"})\n")
}

//Returns a Error that is sitting on the thread.
func (l *implementation) Catch() language.Boolean {
	l.AddHelper(`func Catch() bool {
	if len(errors) > 0 {
		err = errors[len(errors)-1]
		errors = errors[:len(errors)-1]
		return true
	}
	return false
}

`)

	var result Boolean
	result = "Catch()"
	return result
}
