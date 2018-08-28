package script

import "github.com/qlova/script/language"

//Converts a Go string to a language.String.
func (q Script) String(s string) language.String {
	return q.lang.LiteralString(s)
}
