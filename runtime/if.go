package runtime

import "github.com/qlova/script"

//If implements script.Language.If
func (runtime *Runtime) If(first script.If, chain []script.If, last func()) {

	var block = runtime.compile(first.Block, true)

	var blocks = make([]*Block, len(chain))
	for i, link := range chain {
		blocks[i] = runtime.compile(link.Block, true)
	}
	var end *Block = runtime.compile(last, true)

	runtime.WriteStatement(func() {
		if (*first.Condition.T().Runtime)().(bool) {
			block.Jump()
		} else {
			var done bool
			for i, block := range blocks {
				if (*chain[i].Condition.T().Runtime)().(bool) {
					block.Jump()
					done = true
					break
				}
			}
			if !done {
				end.Jump()
			}

		}
	})
}
