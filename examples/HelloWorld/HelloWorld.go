package main

import qlova "github.com/qlova/script"
import "github.com/qlova/script/language/go"

import "fmt"

func main() {
	var HelloWorld = qlova.NewProgram(func (q *qlova.Script) {

		q.Main(func(q *qlova.Script) {
			q.Print(q.String("Hello World"))
		})

	})

	Source, err := HelloWorld.Source(Go.Language())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(Source)
	
	fmt.Println("\nOutput:")
	
	err = HelloWorld.Run()
	if err != nil {
		fmt.Println(err)
	}
}
