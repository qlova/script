package language

import "github.com/qlova/script"

type expression string

//Expression is a helper function.
func Expression(q script.Ctx, s string) script.Type {
	return script.NewType(q, func() interface{} {
		return expression(s)
	})
}
