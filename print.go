package script

import "fmt"

//PrintL prints literal values to Stdout.
func (q Ctx) PrintL(args ...interface{}) {
	var strings Values
	for _, arg := range args {
		strings = append(strings, q.String(fmt.Sprint(arg)))
	}
	q.Print(strings...)
}

//Print prints values to Stdout.
func (q Ctx) Print(args ...Value) {
	q.Language.Print(args)
}
