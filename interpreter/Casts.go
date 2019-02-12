package interpreter

import "strconv"
import "fmt"
import "github.com/qlova/script/language"
import "github.com/qlova/script/interpreter/dynamic"

func (implementation Implementation) Cast(a, b language.Type) language.Type {
	
	var register = implementation.RegisterOf(a)
	switch a.(type) {
		case Integer:
			switch b.(type) {
				case String:
					var result = implementation.ReserveRegister()
					implementation.AddInstruction(func(thread *dynamic.Thread) {
						thread.Set(result, fmt.Sprint(thread.Get(register)))
					})
					return b.Register(strconv.Itoa(result))
					
					
				case Symbol:
					var result = implementation.ReserveRegister()
					implementation.AddInstruction(func(thread *dynamic.Thread) {
						thread.Set(result, rune(thread.Get(register).(int)))
					})
					return b.Register(strconv.Itoa(result))
			}
			
		case Symbol:
			switch b.(type) {
				case String:
					var result = implementation.ReserveRegister()
					implementation.AddInstruction(func(thread *dynamic.Thread) {
						thread.Set(result, string(thread.Get(register).(rune)))
					})
					return b.Register(strconv.Itoa(result))
					
				case Integer:
					var result = implementation.ReserveRegister()
					implementation.AddInstruction(func(thread *dynamic.Thread) {
						thread.Set(result, int(thread.Get(register).(rune)))
					})
					return b.Register(strconv.Itoa(result))
			}
	}
	
	panic(implementation.Name()+".Cast("+a.Name()+", "+b.Name()+") Unimplemented")
	return nil
}

