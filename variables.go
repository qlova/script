package script

import "github.com/qlova/script/go"

func (q Script) Define(name Go.String, value Type) Type {
	q.indent()
	t, statement := q.lang.Define(name, convert(value))
	q.write(statement)
	return q.wrap(t)
}

func (q Script) Set(variable Type, value Type) {
	q.indent()
	q.write(q.lang.Set(variable, convert(value)))
}
