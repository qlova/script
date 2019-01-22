package interpreter

import "strconv"
import "github.com/qlova/script/language"
import "github.com/qlova/script/interpreter/dynamic"

func (implementation Implementation) Register(name string, value language.Type) (language.Statement, language.Type) {
	
	var register int
	
	if literal := implementation.Literal(value); literal != nil {
		register = implementation.ReserveRegister()
		implementation.AddInstruction(func(thread *dynamic.Thread) {
			thread.Set(register, literal)
		})
	} else {
		var other = implementation.RegisterOf(value)
		register = implementation.ReserveRegister()
		implementation.AddInstruction(func(thread *dynamic.Thread) {
			thread.Set(register, thread.Get(other))
		})
	}
	
	return language.Statement(""), value.Register(strconv.Itoa(register))
}

func (implementation Implementation) Set(register string, value language.Type) language.Statement {
	panic(implementation.Name()+".Set() Unimplemented")
	return language.Statement("")
}

func (implementation Implementation) Get(register string, value language.Type) language.Type {
	panic(implementation.Name()+".Get() Unimplemented")
	return nil
}
