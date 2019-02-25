package interpreter

import "strconv"
import "github.com/qlova/script/language"
import "github.com/qlova/script/interpreter/dynamic"

func (implementation Implementation) Register(name string, value language.Type) (language.Statement, language.Type) {
	var register int
	
	switch v := value.(type) {
		case Function:
			var other, _ = strconv.Atoi(string(v.Expression))
			register = implementation.ReserveRegister()
			implementation.AddInstruction(func(thread *dynamic.Thread) {
				thread.Set(register, other)
			})

		default:
			var other = implementation.RegisterOf(value)
			register = implementation.ReserveRegister()
			implementation.AddInstruction(func(thread *dynamic.Thread) {
				thread.Set(register, thread.Get(other))
			})
	}
	
	
	
	return language.Statement(""), value.Register(strconv.Itoa(register))
}

func (implementation Implementation) Set(variable, value language.Type) language.Statement {

	var register = implementation.RegisterOf(value)
	var other = implementation.RegisterOf(value)
	implementation.AddInstruction(func(thread *dynamic.Thread) {
		thread.Set(register, thread.Get(other))
	})
	
	return language.Statement("")
}
