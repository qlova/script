package main

import . "github.com/qlova/script"

import "fmt"

func main() {
	//Create a simple Hello World program.
	var HelloWorld = Program(func (q Script) {

		q.Main(func() {
			q.Print(String("Hello World"))
		})

		a := q.Array(Number)
		a.

		q.Number.Array().Array().Array()

	})

	//Print out the source code of the program in Go.
	code := HelloWorld.Go()
	if code.Error {
		fmt.Println(code.ErrorMessage)
		return
	}
	fmt.Println(code)

	//Run the program and get the output.	
	fmt.Println("\nOutput:")

	err := HelloWorld.Run()
	if err != nil {
		fmt.Println(err)
	}
}
