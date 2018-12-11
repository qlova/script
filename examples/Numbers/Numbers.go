package main

import "fmt"
import . "github.com/qlova/script"
import "github.com/qlova/script/language/go"

const X, Y = 100, 7

func main() {
	
	var software = script.NewProgram(func(q script.Script) {
		q.Main(func() {
			
			q.Print(q.String("My Favourite Number:"), q.Int(22))
			q.Print()

			x, y := q.Int(), q.Int()

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
 
