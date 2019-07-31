package interpreter

import "fmt"

import "github.com/qlova/script/language"
import "github.com/qlova/script/interpreter/dynamic"

func (implementation Implementation) Print(values ...language.Type) language.Statement {
	for i, value := range values {

		var register = implementation.RegisterOf(value)
		implementation.AddInstruction(func(thread *dynamic.Thread) {
			fmt.Print(thread.Get(register))
		})

		if i < len(values)-1 {
			implementation.AddInstruction(func(thread *dynamic.Thread) {
				fmt.Print(" ")
			})
		}
	}

	implementation.AddInstruction(func(thread *dynamic.Thread) {
		fmt.Println()
	})

	return language.Statement("")
}

func (implementation Implementation) Write(stream language.Stream, values ...language.Type) language.Statement {

	if stream == nil {
		for _, value := range values {

			var register = implementation.RegisterOf(value)
			implementation.AddInstruction(func(thread *dynamic.Thread) {
				fmt.Print(thread.Get(register))
			})

		}

		return language.Statement("")
	}

	panic(implementation.Name() + ".Write() Unimplemented")
	return language.Statement("")
}
