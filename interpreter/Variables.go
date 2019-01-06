package interpreter

import "github.com/qlova/script/language"
import "github.com/qlova/script/interpreter/dynamic"

func (implementation Implementation) Register(register string, value language.Type) language.Statement {
	
	if literal := implementation.Literal(value); literal != nil {
		implementation.ReserveRegister()
		implementation.AddInstruction(func(thread *dynamic.Thread) {
			dynamic.Thread.
		})
	} else {
		implementation.AddInstruction(func(thread *dynamic.Thread) {
			
		})
	}
	
	return language.Statement("")
}

func (implementation Implementation) Set(register string, value language.Type) language.Statement {
	panic(implementation.Name()+".Set() Unimplemented")
	return language.Statement("")
}

func (implementation Implementation) Get(register string, value language.Type) language.Type {
	panic(implementation.Name()+".Get() Unimplemented")
	return nil
}
