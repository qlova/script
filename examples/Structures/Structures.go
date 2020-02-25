package main

import (
	"fmt"
	"os"

	"github.com/qlova/script"
	"github.com/qlova/script/language"
	"github.com/qlova/script/runtime"
)

func main() {
	var Structures = func(q script.Ctx) {

		type Vector2 struct {
			//Define Type.
			script.Type `name:"Vector2"`
			Set         func(Vector2)

			//Member Variables
			X, Y script.Int

			//Method signatures.
			Add func(Vector2) Vector2 `name:"Add" args:"b"`
		}

		//Method Definitions.
		var v Vector2
		v.Add = func(b Vector2) (_ Vector2) {
			var c Vector2
			q.Var("c").New(&c)

			c.X.Set((v.X).Plus(b.X))
			c.Y.Set((v.Y).Plus(b.Y))

			q.Return(c)
			return
		}
		q.DefineType(&v, "v")

		q.Main(func() {
			var a = Vector2{
				X: q.Int(2), Y: q.Int(1),
			}
			q.Var("a").New(&a)
			var b = Vector2{
				X: q.Int(3), Y: q.Int(2),
			}
			q.Var("b").New(&b)

			q.Var("c")
			var c = a.Add(b)

			q.Print(c.X, c.Y)
		})
	}

	os.Stdout.Write(language.Go(Structures))

	fmt.Println("\n[Output]")

	runtime.Execute(Structures)
}
