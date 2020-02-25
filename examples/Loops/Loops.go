package main

import (
	"fmt"
	"os"

	"github.com/qlova/script"
	"github.com/qlova/script/language"
	"github.com/qlova/script/runtime"
)

func main() {
	var Loops = func(q script.Ctx) {
		q.Main(func() {
			//for i := 0; i < 5; i++
			q.ForL(0, q.LessThanL(5), q.Plus1(), func(i script.Int) {
				q.Print(i)
			})

			q.WhileL(false, func() {
				q.PrintL("This loop never runs")
			})

			q.Loop(func() {
				q.PrintL("This would be an infinite loop if there was no break")
				q.Break()
			})
		})
	}

	os.Stdout.Write(language.Go(Loops))

	fmt.Println("\n[Output]")

	runtime.Execute(Loops)
}
