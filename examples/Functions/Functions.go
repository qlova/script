package main

import (
	"fmt"
	"os"

	"github.com/qlova/script"
	"github.com/qlova/script/language"
	"github.com/qlova/script/runtime"
)

func main() {
	var Functions = func(q script.Ctx) {
		var Add = func(a, b script.Int) (_ script.Int) {
			q.Return(a.Plus(b)) //Use q.Return
			return              //Don't actually return anything with Go.
		}
		q.DefineFunc(&Add, "Add", "a", "b") //Now you can use the function.

		q.Main(func() {
			q.Print(Add(q.Int(2), q.Int(3)))
		})
	}

	os.Stdout.Write(language.Go(Functions))

	fmt.Println("\n[Output]")

	runtime.Execute(Functions)
}
