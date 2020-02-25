package runtime

import (
	"fmt"
	"reflect"

	"github.com/qlova/script"
)

func (runtime *Runtime) DefineStruct(def script.Struct) {
	for _, method := range def.Methods {
		runtime.Functions[def.Name+"."+method.Name] = runtime.compile(method.Block, true)
	}
}

//Field implements script.Language.Field
func (runtime *Runtime) Field(structure script.Value, name string) script.Result {
	var s = *structure.T().Runtime
	var result = func() interface{} {
		fmt.Println(s())
		return (reflect.ValueOf(s()).Interface().(map[string]interface{}))[name]
	}
	return &result
}
