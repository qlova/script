package runtime

import "github.com/qlova/script"

//Plus implements implements script.Language.Plus
func (runtime *Runtime) Plus(a, b script.Int) script.Int {
	var x, y = *a.T().Runtime, *b.T().Runtime

	return script.Int{Type: Value(a.Ctx, func() interface{} {
		return x().(int) + y().(int)
	})}
}
