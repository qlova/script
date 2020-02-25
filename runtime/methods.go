package runtime

import (
	"reflect"

	"github.com/qlova/script"
)

//RunMethod implements script.Language.RunMethod.
func (runtime *Runtime) RunMethod(structure script.Value, name string, args []script.Value) {
	var NumberOfArguments = len(args)
	var Args = make(map[int]func() interface{}, NumberOfArguments)

	for i, arg := range args {
		Args[i] = (*arg.T().Runtime)
	}

	Args[-1] = (*structure.T().Runtime)

	name = reflect.TypeOf(structure).Elem().Name() + "." + name

	runtime.WriteStatement(func() {
		var block, ok = runtime.Functions[name]
		if !ok {
			panic("invalid block: " + name)
		}
		var RuntimeArgs = make(map[int]interface{}, NumberOfArguments)
		for i := range Args {
			RuntimeArgs[i] = Args[i]()
		}
		block.Args = RuntimeArgs
		block.Jump()
		runtime.returned = runtime.returning
		runtime.returning = nil
	})
}

//CallMethod implements script.Language.CallMethod
func (runtime *Runtime) CallMethod(structure script.Value, name string, args []script.Value) script.Result {
	var NumberOfArguments = len(args)
	var Args = make(map[int]func() interface{}, NumberOfArguments)

	for i, arg := range args {
		Args[i] = (*arg.T().Runtime)
	}

	Args[-1] = (*structure.T().Runtime)

	name = reflect.TypeOf(structure).Elem().Name() + "." + name

	f := func() interface{} {

		var block, ok = runtime.Functions[name]
		if !ok {
			panic("invalid function: " + name)
		}
		var RuntimeArgs = make(map[int]interface{}, NumberOfArguments)
		for i := range Args {
			RuntimeArgs[i] = Args[i]()
		}
		block.Args = RuntimeArgs
		block.Jump()
		runtime.returned = runtime.returning
		runtime.returning = nil

		return runtime.returned
	}
	return &f
}
