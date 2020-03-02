package language

import "github.com/qlova/script"

type expression string

func (e expression) String() string {
	return string(e)
}

//Expression is a helper function.
func Expression(q script.AnyCtx, s string) script.Type {
	return script.NewType(q.RootCtx(), func() interface{} {
		return expression(s)
	})
}
