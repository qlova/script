package interpreter

import "strconv"
import "github.com/qlova/script/language"
import "github.com/qlova/script/interpreter/dynamic"

func (implementation Implementation) ArrayOf(t language.Type, length int) language.Array {
	switch t.(type) {
		case Integer:
			var register = implementation.ReserveRegister()
			implementation.AddInstruction(func(thread *dynamic.Thread) {
				thread.Set(register, make([]int, length))
			})
			return Array{
				Length: length,
				Subtype: t,
				Expression: language.Statement(strconv.Itoa(register)),
			}
			
		//TODO create reflection version.
	}
	panic(implementation.Name()+".ArrayOf("+t.Name()+", int) Unimplemented")
	return nil
}

func (implementation Implementation) Join(a, b language.Type) language.Type {
	
	switch a.(type) {
		case String:
			switch b.(type) {
				case String:
					var a = implementation.RegisterOf(a)
					var b = implementation.RegisterOf(b)
					var register = implementation.ReserveRegister()
					implementation.AddInstruction(func(thread *dynamic.Thread) {
						thread.Set(register, thread.Get(a).(string)+thread.Get(b).(string))
					})
					return String{
						Expression: language.Statement(strconv.Itoa(register)),
					}
			}
	}
	
	panic(implementation.Name()+".Join() Unimplemented")
	return nil
}

func (implementation Implementation) Length(t language.Type) language.Integer {
	panic(implementation.Name()+".Length() Unimplemented")
	return nil
}

func (implementation Implementation) Push(value language.Type, list language.Type) language.Statement {
	panic(implementation.Name()+".Push() Unimplemented")
	return language.Statement("")
}

func (implementation Implementation) Pop(list language.Type) language.Type {
	panic(implementation.Name()+".Pop() Unimplemented")
	return nil
}

func (implementation Implementation) TableOf(t language.Type) language.Table {
	panic(implementation.Name()+".TableOf() Unimplemented")
	return nil
}

func (implementation Implementation) ListOf(t language.Type) language.List {
	panic(implementation.Name()+".ListOf() Unimplemented")
	return nil
}

