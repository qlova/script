package script

import "fmt"

//PrintL prints literal values to Stdout.
func (q Ctx) PrintL(args ...interface{}) {
	var strings []AnyValue
	for _, arg := range args {
		strings = append(strings, q.String(fmt.Sprint(arg)))
	}
	q.Print(strings...)
}

//Print prints values to Stdout.
func (q Ctx) Print(args ...AnyValue) {
	var values Values
	for _, arg := range args {
		values = append(values, arg.ValueFromCtx(q))
	}
	q.Language.Print(values)
}
