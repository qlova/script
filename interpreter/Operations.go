package interpreter

import "strconv"
import "github.com/qlova/script/language"
import "github.com/qlova/script/interpreter/dynamic"

func (implementation Implementation) Add(a, b language.Number) language.Number {
	var register int
	
	A, B := a.(Integer).Literal, b.(Integer).Literal
	
	if A != nil && B != nil {
		return Integer{
			Literal: A.(int) + B.(int),
		}
	}
	if A == nil && B == nil {
		register = implementation.ReserveRegister()
		RA, RB := implementation.RegisterOf(a), implementation.RegisterOf(b)
		implementation.AddInstruction(func(thread *dynamic.Thread) {
			thread.Set(register, thread.Get(RA).(int) + thread.Get(RB).(int))
		})
	}
	if A == nil && B != nil {
		register = implementation.ReserveRegister()
		RA, RB := implementation.RegisterOf(a), B.(int)
		implementation.AddInstruction(func(thread *dynamic.Thread) {
			thread.Set(register, thread.Get(RA).(int) + RB)
		})
	}
	if A != nil && B == nil {
		register = implementation.ReserveRegister()
		RA, RB :=  A.(int), implementation.RegisterOf(b)
		implementation.AddInstruction(func(thread *dynamic.Thread) {
			thread.Set(register, RA + thread.Get(RB).(int))
		})
	}

	return a.Register(strconv.Itoa(register)).(language.Number)
}

func (implementation Implementation) Sub(a, b language.Number) language.Number {
	var register int
	
	A, B := a.(Integer).Literal, b.(Integer).Literal
	
	if A != nil && B != nil {
		return Integer{
			Literal: A.(int) - B.(int),
		}
	}
	if A == nil && B == nil {
		register = implementation.ReserveRegister()
		RA, RB := implementation.RegisterOf(a), implementation.RegisterOf(b)
		implementation.AddInstruction(func(thread *dynamic.Thread) {
			thread.Set(register, thread.Get(RA).(int) - thread.Get(RB).(int))
		})
	}
	if A == nil && B != nil {
		register = implementation.ReserveRegister()
		RA, RB := implementation.RegisterOf(a), B.(int)
		implementation.AddInstruction(func(thread *dynamic.Thread) {
			thread.Set(register, thread.Get(RA).(int) - RB)
		})
	}
	if A != nil && B == nil {
		register = implementation.ReserveRegister()
		RA, RB :=  A.(int), implementation.RegisterOf(b)
		implementation.AddInstruction(func(thread *dynamic.Thread) {
			thread.Set(register, RA - thread.Get(RB).(int))
		})
	}

	return a.Register(strconv.Itoa(register)).(language.Number)
}

func (implementation Implementation) Mul(a, b language.Number) language.Number {
	
	var register int
	
	A, B := a.(Integer).Literal, b.(Integer).Literal
	
	if A != nil && B != nil {
		return Integer{
			Literal: A.(int) * B.(int),
		}
	}
	if A == nil && B == nil {
		register = implementation.ReserveRegister()
		RA, RB := implementation.RegisterOf(a), implementation.RegisterOf(b)
		implementation.AddInstruction(func(thread *dynamic.Thread) {
			thread.Set(register, thread.Get(RA).(int) * thread.Get(RB).(int))
		})
	}
	if A == nil && B != nil {
		register = implementation.ReserveRegister()
		RA, RB := implementation.RegisterOf(a), B.(int)
		implementation.AddInstruction(func(thread *dynamic.Thread) {
			thread.Set(register, thread.Get(RA).(int) * RB)
		})
	}
	if A != nil && B == nil {
		register = implementation.ReserveRegister()
		RA, RB :=  A.(int), implementation.RegisterOf(b)
		implementation.AddInstruction(func(thread *dynamic.Thread) {
			thread.Set(register, RA * thread.Get(RB).(int))
		})
	}

	return a.Register(strconv.Itoa(register)).(language.Number)
}

func (implementation Implementation) Div(a, b language.Number) language.Number {
	var register int
	
	A, B := a.(Integer).Literal, b.(Integer).Literal
	
	if A != nil && B != nil {
		return Integer{
			Literal: A.(int) / B.(int),
		}
	}
	if A == nil && B == nil {
		register = implementation.ReserveRegister()
		RA, RB := implementation.RegisterOf(a), implementation.RegisterOf(b)
		implementation.AddInstruction(func(thread *dynamic.Thread) {
			thread.Set(register, thread.Get(RA).(int) / thread.Get(RB).(int))
		})
	}
	if A == nil && B != nil {
		register = implementation.ReserveRegister()
		RA, RB := implementation.RegisterOf(a), B.(int)
		implementation.AddInstruction(func(thread *dynamic.Thread) {
			thread.Set(register, thread.Get(RA).(int) / RB)
		})
	}
	if A != nil && B == nil {
		register = implementation.ReserveRegister()
		RA, RB :=  A.(int), implementation.RegisterOf(b)
		implementation.AddInstruction(func(thread *dynamic.Thread) {
			thread.Set(register, RA / thread.Get(RB).(int))
		})
	}

	return a.Register(strconv.Itoa(register)).(language.Number)
}

func pow(base, exp int) int {
	var result = 1;
	for {
		if (exp & 1 > 0) {
			result *= base
	}
		exp >>= 1
		if (exp == 0) {
			break
	}
		base *= base
	}

	return result
}

func (implementation Implementation) Pow(a, b language.Number) language.Number {
	var register int
	
	A, B := a.(Integer).Literal, b.(Integer).Literal
	
	if A != nil && B != nil {
		return Integer{
			Literal: pow(A.(int),  B.(int)),
		}
	}
	if A == nil && B == nil {
		register = implementation.ReserveRegister()
		RA, RB := implementation.RegisterOf(a), implementation.RegisterOf(b)
		implementation.AddInstruction(func(thread *dynamic.Thread) {
			thread.Set(register, pow(thread.Get(RA).(int), thread.Get(RB).(int)))
		})
	}
	if A == nil && B != nil {
		register = implementation.ReserveRegister()
		RA, RB := implementation.RegisterOf(a), B.(int)
		implementation.AddInstruction(func(thread *dynamic.Thread) {
			thread.Set(register, pow(thread.Get(RA).(int), RB))
		})
	}
	if A != nil && B == nil {
		register = implementation.ReserveRegister()
		RA, RB :=  A.(int), implementation.RegisterOf(b)
		implementation.AddInstruction(func(thread *dynamic.Thread) {
			thread.Set(register, pow(RA,  thread.Get(RB).(int)))
		})
	}

	return a.Register(strconv.Itoa(register)).(language.Number)
}

func (implementation Implementation) Mod(a, b language.Number) language.Number {
	var register int
	
	A, B := a.(Integer).Literal, b.(Integer).Literal
	
	if A != nil && B != nil {
		return Integer{
			Literal: A.(int) % B.(int),
		}
	}
	if A == nil && B == nil {
		register = implementation.ReserveRegister()
		RA, RB := implementation.RegisterOf(a), implementation.RegisterOf(b)
		implementation.AddInstruction(func(thread *dynamic.Thread) {
			thread.Set(register, thread.Get(RA).(int) % thread.Get(RB).(int))
		})
	}
	if A == nil && B != nil {
		register = implementation.ReserveRegister()
		RA, RB := implementation.RegisterOf(a), B.(int)
		implementation.AddInstruction(func(thread *dynamic.Thread) {
			thread.Set(register, thread.Get(RA).(int) % RB)
		})
	}
	if A != nil && B == nil {
		register = implementation.ReserveRegister()
		RA, RB :=  A.(int), implementation.RegisterOf(b)
		implementation.AddInstruction(func(thread *dynamic.Thread) {
			thread.Set(register, RA % thread.Get(RB).(int))
		})
	}

	return a.Register(strconv.Itoa(register)).(language.Number)
}

