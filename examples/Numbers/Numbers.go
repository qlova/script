package main

import "fmt"
import "github.com/qlova/script"

const X, Y = 100, 7

func main() {
	
	var software = script.Program(func(q script.Script) {
		q.Main(func() {
			
			q.Print(q.String("My Favourite Number:"), q.Int(22))
			q.Print()

			x, y := q.Int(X).Var("x"), q.Int(Y).Var("y")

			q.Print(q.String("Operations on"), x, q.String("&"), y)
			q.Print(q.String("======================"))
			
			q.Print(q.String("Double:"), x.Mul(q.Int(2)), q.Int(2).Mul(y))
			q.Print(q.String("Add:"), x.Add(y))
			q.Print(q.String("Sub:"), x.Sub(y))
			q.Print(q.String("Mul:"), x.Mul(y))
			q.Print(q.String("Div:"), x.Div(y))
			q.Print(q.String("Mod:"), x.Mod(y))
			q.Print(q.String("Pow:"), x.Pow(y))
		})
	})
	
	//Print out the source code of the program in Go.
	code := software.Go()
	if code.Error {
		fmt.Println(code.ErrorMessage)
		return
	}
	fmt.Println(code)

	//Run the program and get the output.	
	fmt.Println("\nOutput:")

	err := software.Run()
	if err != nil {
		fmt.Println(err)
	}
}
 
