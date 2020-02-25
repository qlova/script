package runtime

import (
	"reflect"

	"github.com/qlova/script"
)

//Index implements script.Language.Index
func (runtime *Runtime) Index(collection script.Value, index script.Int) script.Result {
	var a = *collection.T().Runtime
	var i = *index.T().Runtime
	var result = func() interface{} {
		return reflect.ValueOf(a()).Index(i().(int))
	}
	return &result
}

//Mutate implements script.Language.Mutate
func (runtime *Runtime) Mutate(collection script.Value, index script.Int, value script.Value) {
	var a = *collection.T().Runtime
	var i = *index.T().Runtime
	var v = *value.T().Runtime
	runtime.WriteStatement(func() {
		reflect.ValueOf(a()).Index(i().(int)).Set(reflect.ValueOf(v()))
	})
}

//Lookup implements script.Language.Lookup
func (runtime *Runtime) Lookup(collection script.Value, index script.String) script.Result {
	var a = *collection.T().Runtime
	var i = *index.T().Runtime
	var result = func() interface{} {
		return reflect.ValueOf(a()).MapIndex(reflect.ValueOf(i().(string)))
	}
	return &result
}

//Update implements script.Language.Update
func (runtime *Runtime) Insert(collection script.Value, index script.String, value script.Value) {
	var a = *collection.T().Runtime
	var i = *index.T().Runtime
	var v = *value.T().Runtime
	runtime.WriteStatement(func() {
		reflect.ValueOf(a()).SetMapIndex(reflect.ValueOf(i().(string)), reflect.ValueOf(v()))
	})
}
