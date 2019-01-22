# Qlovascript
Programmably create programs with Qlovascript!

Example:
```
	import qlova "github.com/qlova/script"
	
	---

	var HelloWorld = qlova.Program(func (q qlova.Script) {
		q.Main(func() {
			q.Print(q.String("Hello World"))
		})
	})
	
	HelloWorld.Run()
```
Qlovascript has a more powerful feature, it can output the Go source code of a program.
```
	---

	Source := HelloWorld.Go()
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
