package main

import "fmt"
import "github.com/qlova/script"
import "github.com/qlova/script/language/go"

const X, Y = 100, 7

func main() {
	
	var software = script.NewProgram(func(q script.Script) {
		q.Main(func(q script.Script) {
			
			q.Print(q.String("My Favourite Number:"), q.Number(22))
			q.Print()
			
			x := q.Define("x", q.Number(X)).(script.Number)
			y := q.Define("y", q.Number(Y)).(script.Number)

			q.Print(q.String("Operations on"), x, q.String("&"), y)
			q.Print(q.String("======================"))
			
			q.Print(q.String("Double:"), q.Mul(x, q.Number(2)), q.Mul(q.Number(2), y))
			q.Print(q.String("Add:"), q.Add(x, y))
			q.Print(q.String("Sub:"), q.Sub(x, y))
			q.Print(q.String("Mul:"), q.Mul(x, y))
			q.Print(q.String("Div:"), q.Div(x, y))
			q.Print(q.String("Mod:"), q.Mod(x, y))
			q.Print(q.String("Pow:"), q.Pow(x, y))
		})
	})
	
	Source, err := software.Source(Go.Language())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(Source)
	
	fmt.Println("\nOutput:")
	
	err = software.Run()
	if err != nil {
		fmt.Println(err)
	}
}
 
