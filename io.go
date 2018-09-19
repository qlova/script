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

//Prints 'values' to Stdout with a trailing newline.
func (q *Script) Write(value ...Type) {
	
	var values = make([]language.Type, len(value))
	for i := range value {
		values[i] = convert(value[i])
	}
	
	q.indent()
	q.write(q.lang.Write(nil, values...))
}

//Prints 'values' to Stdout with a trailing newline.
func (q *Script) Read(value Type) Type {	
	switch value.(type) {
		case Symbol:
			
			return q.wrap(q.lang.Read(nil, convert(value)))
			
		default:
			panic("Script.Read("+value.Name()+"): Unimplemented")
			return nil
	}
}
