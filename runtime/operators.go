package runtime

import "github.com/qlova/script"

//Plus implements implements script.Language.Plus
func (runtime *Runtime) Plus(a, b script.Int) script.Int {
	var x, y = *a.T().Runtime, *b.T().Runtime

	return script.Int{Type: Value(a.Ctx, func() interface{} {
		return x().(int) + y().(int)
	})}
}

//Join implements implements script.Language.Join
func (runtime *Runtime) Join(a, b script.String) script.String {
	var x, y = *a.T().Runtime, *b.T().Runtime

	return script.String{Type: Value(a.Ctx, func() interface{} {
		return x().(string) + y().(string)
	})}
}

//Same implements implements script.Language.Same
func (runtime *Runtime) Same(a, b script.Int) script.Bool {
	var x, y = *a.T().Runtime, *b.T().Runtime

	return script.Bool{Type: Value(a.Ctx, func() interface{} {
		return x().(string) == y().(string)
	})}
}

//Not implements implements script.Language.Not
func (runtime *Runtime) Not(a script.Bool) script.Bool {
	var x = *a.T().Runtime

	return script.Bool{Type: Value(a.Ctx, func() interface{} {
		return !x().(bool)
	})}
}
