package main

import (
	"fmt"
	"os"

	"github.com/qlova/script"
	"github.com/qlova/script/language"
	"github.com/qlova/script/runtime"
)

func main() {
	var Conditionals = func(q script.Ctx) {
		q.Main(func() {
			q.IfL(true, func() {
				q.PrintL("This will run")
			}).ElseIfL(false, func() {
				q.PrintL("This will not run")
			}).Else(func() {
				q.PrintL("This will not run")
			}).End() //End call is required.
		})
	}

	os.Stdout.Write(language.Go(Conditionals))

	fmt.Println("\n[Output]")

	runtime.Execute(Conditionals)
}
