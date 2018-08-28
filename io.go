package script

import "github.com/qlova/script/language"

//Prints 'values' to Stdout with a trailing newline.
func (q *Script) Print(value ...language.Type) {	
	q.indent()
	q.write(q.lang.Print(value...))
}
