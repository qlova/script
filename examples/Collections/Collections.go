package main

import (
	"fmt"
	"os"

	"github.com/qlova/script"
	"github.com/qlova/script/language"
	"github.com/qlova/script/runtime"
)

func main() {
	var Collections = func(q script.Ctx) {
		q.Main(func() {
			var a = q.Var("a").IntListL(0)
			a.MutateL(0, 1)
			q.Print(a.IndexL(0))

			var t = q.Var("t").IntTable(nil)
			t.InsertL("key", 1)
			q.Print(t.LookupL("key"))
		})
	}

	os.Stdout.Write(language.Go(Collections))

	fmt.Println("\n[Output]")

	runtime.Execute(Collections)
}
