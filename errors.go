package script

import "github.com/qlova/script/go"
import "github.com/qlova/script/language"

type Error struct {
	language.Error
	EmbeddedScript
}

func (q Script) Trace(line int, file Go.String) {
	q.write(q.lang.Trace(line, file))
}

func (q Script) Error() Error {
	return Error{Error:q.lang.Error(), EmbeddedScript:EmbeddedScript{q:q}}
}

func (q Script) Throw(code Number, message string) {
	q.write(q.lang.Throw(code.Number, message.String))
}

func (q Script) Catch() Boolean {
	return q.wrap(q.lang.Catch()).(Boolean)
}
