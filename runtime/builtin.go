package runtime

import (
	"fmt"

	"github.com/qlova/script"
)

//Print implements script.Language.Print
func (runtime *Runtime) Print(args script.Values) {
	runtime.WriteStatement(func() {
		fmt.Println(script.Values(args).Runtime()...)
	})
}
