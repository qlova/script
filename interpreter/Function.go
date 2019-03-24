package interpreter

import "strconv"
import "github.com/qlova/script/interpreter/dynamic"
import "github.com/qlova/script/language"

func (implementation Implementation) Function(name string, registers []string, arguments []language.Type, returns language.Type) (language.Statement, language.Function) {
	
	if name != "" {
		
		var block = implementation.program.CreateBlock()

		implementation.Activate(block)
		
		for i, register := range registers {
			var table = implementation.program[block].Arguments[register]
				table[0] = i
				
			 implementation.program[block].Arguments[register] = table
		}

		implementation.program[block].Registers += len(registers)
		
		implementation.program[block].Name = name
		
		return "", Function{
			Expression: language.Statement(block.String()),
			Subtype: returns,
		}
	}
	
	panic(implementation.Name()+".Function() Unimplemented")
	return language.Statement(""), nil
}

func (implementation Implementation) EndFunction() language.Statement {
	implementation.Deactivate()

	return language.Statement("")
}

func (implementation Implementation) Call(f language.Function, arguments []language.Type) language.Type {
	var block = implementation.ExpressionOf(f)

	var length = len(arguments)
	var addresses = make([]int, length)
	for i := range arguments {
		addresses[i] = implementation.RegisterOf(arguments[i])
	}
	
	var returnRegister = implementation.ReserveRegister()
	
	//TODO Need to sort arguments?
	_, err := strconv.Atoi(string(block))
	
	if block[0] == '$' || err != nil {
		var register, _ = strconv.Atoi(string(block[1:]))
		implementation.AddInstruction(func(thread *dynamic.Thread) {
			var Converted = make([]interface{}, length)
			for i := range addresses {
				Converted[i] = thread.Get(addresses[i])
			}
			
			
			thread.Returns = returnRegister
			thread.JumpTo(dynamic.BlockPointer(thread.Get(register).(int)), Converted...)
		})	
	} else {
		var pointer = implementation.RegisterOf(f)
		implementation.AddInstruction(func(thread *dynamic.Thread) {
			var Converted = make([]interface{}, length)
			for i := range addresses {
				Converted[i] = thread.Get(addresses[i])
			}
			
			thread.Returns = returnRegister
			thread.JumpTo(dynamic.BlockPointer(pointer), Converted...)
		})
	}
	
	return f.(Function).Subtype.Register(strconv.Itoa(returnRegister))
}

func (implementation Implementation) Run(f language.Function, arguments []language.Type) language.Statement {
	var block = implementation.ExpressionOf(f)

	var length = len(arguments)
	var addresses = make([]int, length)
	for i := range arguments {
		addresses[i] = implementation.RegisterOf(arguments[i])
	}
	
	//TODO Need to sort arguments?
	_, err := strconv.Atoi(string(block))
	
	if block[0] == '$' || err != nil {
		var register, _ = strconv.Atoi(string(block[1:]))
		implementation.AddInstruction(func(thread *dynamic.Thread) {
			var Converted = make([]interface{}, length)
			for i := range addresses {
				Converted[i] = thread.Get(addresses[i])
			}
			
			thread.JumpTo(dynamic.BlockPointer(thread.Get(register).(int)), Converted...)
		})	
	} else {
		var pointer = implementation.RegisterOf(f)
		implementation.AddInstruction(func(thread *dynamic.Thread) {
			var Converted = make([]interface{}, length)
			for i := range addresses {
				Converted[i] = thread.Get(addresses[i])
			}
			
			thread.JumpTo(dynamic.BlockPointer(pointer), Converted...)
		})
	}
	
	return ""
}

func (implementation Implementation) Return(value language.Type) language.Statement {
	var register = implementation.RegisterOf(value)
	
	implementation.AddInstruction(func(thread *dynamic.Thread) {
		thread.Return(thread.Get(register))
	})
	
	return ""
}

