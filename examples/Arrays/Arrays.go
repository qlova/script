package main

import qlova "github.com/qlova/script"

import "fmt"

func main() {
	var Arrays = qlova.Program(func(q qlova.Script) {

		q.Main(func() {

			a := q.Int().Array(1).Var("a")

			q.Print(a.Index(q.Int(0)))

			a.Modify(q.Int(0), q.Int(1))

			q.Print(a.Index(q.Int(0)))
		})

	})

	//Print out the source code of the program in Go.
	code := Arrays.Go()
	if code.Error {
		fmt.Println(code.ErrorMessage)
		return
	}
	fmt.Println(code)

	//Run the program and get the output.
	fmt.Println("\nOutput:")

	err := Arrays.Run()
	if err != nil {
		fmt.Println(err)
	}
}
