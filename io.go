package script

import "github.com/qlova/script/language"

//Prints 'values' to Stdout with a trailing newline.
func (q *Script) Print(value ...Type) {
	
	var values = make([]language.Type, len(value))
	for i := range value {
		values[i] = convert(value[i])
	}
	
	q.indent()
	q.write(q.lang.Print(values...))
}
