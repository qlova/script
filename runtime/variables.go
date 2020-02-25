package runtime

import (
	"reflect"

	"github.com/qlova/script"
)

//DefineVariable implements script.Language.DefineVariable
func (runtime *Runtime) DefineVariable(name string, value script.Value) script.Result {
	var v = *value.T().Runtime
	runtime.WriteStatement(func() {
		runtime.Current.define(name, v())
	})
	f := func() interface{} {
		return runtime.Current.get(name)
	}
	return &f
}

//Set implements script.Language.Set
func (runtime *Runtime) Set(a script.Value, b script.Value) {
	var A = *a.T().Runtime
	var B = *b.T().Runtime
	runtime.WriteStatement(func() {
		reflect.ValueOf(A).Set(reflect.ValueOf(B()))
	})
}
