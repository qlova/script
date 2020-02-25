package main

import (
	"fmt"
	"os"

	"github.com/qlova/script"
	"github.com/qlova/script/language"
	"github.com/qlova/script/runtime"
)

func main() {
	var HelloWorld = func(q script.Ctx) {
		q.Main(func() {
			q.PrintL("Hello World")
		})
	}

	os.Stdout.Write(language.Go(HelloWorld))

	fmt.Println("\n[Output]")

	runtime.Execute(HelloWorld)
}
