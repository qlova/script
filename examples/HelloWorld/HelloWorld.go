package main

import qlova "github.com/qlova/script"
import "github.com/qlova/script/languages/go"

import "fmt"

func main() {
	var software = qlova.NewScript()
	
	software.SetLanguage(Go.Language())
	
	software.Main(func(q *qlova.Script) {
		q.Print(q.StringLiteral("Hello World\n"))
	})
	
	fmt.Println(software)
}
