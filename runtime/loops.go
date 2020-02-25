package runtime

import "github.com/qlova/script"

//While implements script.Language.While
func (runtime *Runtime) While(condition script.Bool, f func()) {
	var block = runtime.compile(f, true)

	runtime.WriteStatement(func() {
		for !runtime.broken && (*condition.T().Runtime)().(bool) {
			block.Jump()
		}
		runtime.broken = false
	})
}

//Loop implements script.Language.Loop
func (runtime *Runtime) Loop(f func()) {
	var block = runtime.compile(f, true)

	runtime.WriteStatement(func() {
		for !runtime.broken {
			block.Jump()
		}
		runtime.broken = false
	})
}

//Break implements script.Language.Break
func (runtime *Runtime) Break() {
	runtime.WriteStatement(func() {
		runtime.broken = true
	})
}
