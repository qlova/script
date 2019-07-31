package interpreter

import "strconv"
import "github.com/qlova/script/language"
import "github.com/qlova/script/interpreter/dynamic"

func (implementation Implementation) And(a, b language.Bit) language.Bit {
	var register = implementation.ReserveRegister()

	RA, RB := implementation.RegisterOf(a), implementation.RegisterOf(b)

	implementation.AddInstruction(func(thread *dynamic.Thread) {
		thread.Set(register, thread.Get(RA).(bool) && thread.Get(RB).(bool))
	})

	return Bit{
		Expression: language.Statement(strconv.Itoa(register)),
	}
}

func (implementation Implementation) Or(a, b language.Bit) language.Bit {
	var register = implementation.ReserveRegister()

	RA, RB := implementation.RegisterOf(a), implementation.RegisterOf(b)

	implementation.AddInstruction(func(thread *dynamic.Thread) {
		thread.Set(register, thread.Get(RA).(bool) || thread.Get(RB).(bool))
	})

	return Bit{
		Expression: language.Statement(strconv.Itoa(register)),
	}
}

func (implementation Implementation) Not(b language.Bit) language.Bit {
	panic(implementation.Name() + ".Not() Unimplemented")
	return nil
}

func (implementation Implementation) Equals(a, b language.Type) language.Bit {

	var register = implementation.ReserveRegister()

	RA, RB := implementation.RegisterOf(a), implementation.RegisterOf(b)

	implementation.AddInstruction(func(thread *dynamic.Thread) {
		thread.Set(register, thread.Get(RA) == thread.Get(RB))
	})

	return Bit{
		Expression: language.Statement(strconv.Itoa(register)),
	}
}

func (implementation Implementation) Smaller(a, b language.Type) language.Bit {
	panic(implementation.Name() + ".Smaller() Unimplemented")
	return nil
}

func (implementation Implementation) Greater(a, b language.Type) language.Bit {
	panic(implementation.Name() + ".Greater() Unimplemented")
	return nil
}
