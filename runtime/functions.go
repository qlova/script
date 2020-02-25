package runtime

import (
	"fmt"

	"github.com/qlova/script"
)

//Return implements script.Language.Return
func (runtime *Runtime) Return(v script.Value) {
	var V = *v.T().Runtime
	runtime.WriteStatement(func() {
		runtime.returning = V()
	})
}

//DefineFunction implements script.Language.DefineFunction
func (runtime *Runtime) DefineFunction(function script.Function) {
	runtime.Functions[function.Name] = runtime.compile(function.Block, true)
}

//Argument implements script.Language.Argument
func (runtime *Runtime) Argument(name string, nth int) script.Result {
	f := func() interface{} {
		fmt.Println(nth, runtime.Current.Args[nth])
		return runtime.Current.Args[nth]
	}
	return &f
}

//CallFunction implements script.Language.CallFunction
func (runtime *Runtime) CallFunction(name string, args []script.Value) script.Result {
	var NumberOfArguments = len(args)
	var RuntimeArgs = make(map[int]interface{}, NumberOfArguments)

	for i, arg := range args {
		RuntimeArgs[i] = (*arg.T().Runtime)()
	}

	f := func() interface{} {

		var block, ok = runtime.Functions[name]
		if !ok {
			panic("invalid function: " + name)
		}
		block.Args = RuntimeArgs
		block.Jump()
		runtime.returned = runtime.returning
		runtime.returning = nil

		return runtime.returned
	}
	return &f
}

//RunFunction implements script.Language.RunFunction.
func (runtime *Runtime) RunFunction(name string, args []script.Value) {

	var NumberOfArguments = len(args)
	var RuntimeArgs = make(map[int]interface{}, NumberOfArguments)

	for i, arg := range args {
		RuntimeArgs[i] = (*arg.T().Runtime)()
	}

	runtime.WriteStatement(func() {
		var block, ok = runtime.Functions[name]
		if !ok {
			panic("invalid block: " + name)
		}
		block.Args = RuntimeArgs
		block.Jump()
		runtime.returned = runtime.returning
		runtime.returning = nil
	})
}
