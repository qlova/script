package runtime

import (
	"fmt"

	"github.com/qlova/script"
)

//Runtime implements the script.Language interface.
//Providing a way to execute scripts at runtime.
type Runtime struct {
	//The current block we are focused on.
	Current *Block

	//Entrypoint is main block of the script.
	Entrypoint *Block

	//Functions are all of the functions in the script.
	Functions map[string]*Block

	//Has the current loop been broken?
	broken bool

	//The returning value.
	returning, returned interface{}
}

func (runtime *Runtime) Raw(value script.Value) string {
	return fmt.Sprint((*value.T().Runtime)())
}

//Main function.
func (runtime *Runtime) Main(f func()) {
	runtime.Entrypoint = runtime.compile(f, false)
}

func (runtime *Runtime) compile(function func(), parent bool) *Block {
	var block = new(Block)
	block.runtime = runtime
	if parent {
		block.Parent = runtime.Current
	}

	var backup = runtime.Current

	runtime.Current = block
	function()
	runtime.Current = backup

	return block
}

//Not implements implements script.Language.Not
func (runtime *Runtime) Write([]byte) (int, error) {
	return 0, nil
}

//Execute compiles and executes a script.
func Execute(s script.Script) {
	Compile(s)()
}

//Compile returns a function created from the given script.
func Compile(s script.Script) func() {
	var q = script.NewCtx()
	var r = new(Runtime)
	r.Functions = make(map[string]*Block)
	q.Language = r
	s(q)

	return r.Entrypoint.Jump
}
