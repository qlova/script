# Qlovascript
Programmably create programs with Qlovascript!

Example:
```
	import qlova "github.com/qlova/script"
	
	---

	var HelloWorld = qlova.NewProgram(func (q *qlova.Script) {

		q.Main(func(q *qlova.Script) {
			q.Print(q.String("Hello World"))
		})

	})
	
	HelloWorld.Run()
```
Qlovascript has a more powerful feature, it can output the Go source code of a program.
```
	import "github.com/qlova/script/language/go"
	
	---

	Source, _ := HelloWorld.Source(Go.Language())
	fmt.Println(Source)
	/*
		package main
		
		import "fmt"
		
		func main() {
			fmt.Println("Hello World")
		}
		
	*/
```

Qlovascript is not ready to be used yet. A full Interpreter and Go implementation is in the works.  
Please check back later!
