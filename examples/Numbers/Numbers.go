package main

import "fmt"
import . "github.com/qlova/script"

const X, Y = 100, 7

func main() {
	
	var software = Program(func(q Script) {
		q.Main(func() {
			
			q.Print(String("My Favourite Number:"), Int(22))
			q.Print()

			x, y := q.Int(Int(X)), q.Int(Int(X))

			q.Print(String("Operations on"), x, String("&"), y)
			q.Print(String("======================"))
			
			q.Print(String("Double:"), x.Mul(Int(2)), Int(2).Mul(y))
			q.Print(String("Add:"), x.Add(y))
			q.Print(String("Sub:"), x.Sub(y))
			q.Print(String("Mul:"), x.Mul(y))
			q.Print(String("Div:"), x.Div(y))
			q.Print(String("Mod:"), x.Mod(y))
			q.Print(String("Pow:"), x.Pow(y))
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
 
