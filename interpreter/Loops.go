package interpreter

import "reflect"
import "strconv"
import "github.com/qlova/script/language"
import "github.com/qlova/script/interpreter/dynamic"

func (implementation Implementation) PushLoop() {
	implementation.loops = append(implementation.loops, implementation.Instruction())
}

func (implementation Implementation) PopLoop() int {
	var result = implementation.loops[len(implementation.loops)-1]
	implementation.loops = implementation.loops[:len(implementation.loops)-1]
	return result
}

func (implementation Implementation) Loop() language.Statement {
	panic(implementation.Name()+".Loop() Unimplemented")
	return language.Statement("")
}

func (implementation Implementation) EndLoop() language.Statement {
	panic(implementation.Name()+".EndLoop() Unimplemented")
	return language.Statement("")
}

func (implementation Implementation) Break() language.Statement {
	panic(implementation.Name()+".Break() Unimplemented")
	return language.Statement("")
}

func (implementation Implementation) While(condition language.Bit) language.Statement {
	panic(implementation.Name()+".While() Unimplemented")
	return language.Statement("")
}

func (implementation Implementation) EndWhile() language.Statement {
	panic(implementation.Name()+".EndWhile() Unimplemented")
	return language.Statement("")
}

func (implementation Implementation) ForRange(i string, a, b language.Number) (language.Statement, language.Type) {
	panic(implementation.Name()+".ForRange() Unimplemented")
	return language.Statement(""), nil
}

func (implementation Implementation) EndForRange() language.Statement {
	panic(implementation.Name()+".EndForRange() Unimplemented")
	return language.Statement("")
}

func (implementation Implementation) ForEach(i, v string, list language.Type) (language.Statement, language.Type, language.Type) {

	var register = implementation.RegisterOf(list)
	
	//Initalise index and value registers.
	var index = implementation.ReserveRegister()
	var value = implementation.ReserveRegister()
	implementation.AddInstruction(func(thread *dynamic.Thread) {
		thread.Set(index, -1)
	})
	
	//This is the beginning of the loop.
	implementation.PushLoop()

	//We create two intructions, one is the condition and sets the variables.
	//The second will jump to the end of the for loop.
	
	implementation.AddInstruction(func(thread *dynamic.Thread) {
		thread.Set(index, thread.Get(index).(int)+1) //i++
		
		if thread.Get(index).(int) < reflect.ValueOf(thread.Get(register)).Len() {
			
			thread.Set(value, reflect.ValueOf(thread.Get(register)).Index(thread.Get(index).(int)).Interface())

			thread.InstructionCounter++ //Skip next instruction.
		}
	})
	
	//Since we don't know where the end of the for loop is, we will leave this instruction blank for now.
	implementation.AddInstruction(func(thread *dynamic.Thread) {})
	
	
	return language.Statement(""), Integer{Expression:language.Statement(strconv.Itoa(index))}, list.(List).Subtype.Register(strconv.Itoa(value))
}

func (implementation Implementation) EndForEach() language.Statement {
	var StartOfLoop = implementation.PopLoop()
	
	implementation.AddInstruction(func(thread *dynamic.Thread) {
		thread.InstructionCounter = StartOfLoop-1
	})
	
	var EndOfLoop = implementation.Instruction()

	implementation.SetInstruction(StartOfLoop+1, func(thread *dynamic.Thread) {
		thread.InstructionCounter = EndOfLoop-1
	})

	return language.Statement("")
}

func (implementation Implementation) For(i string, condition language.Bit, action language.Statement) (language.Statement, language.Type) {
	panic(implementation.Name()+".For() Unimplemented")
	return language.Statement(""), nil
}

func (implementation Implementation) EndFor() language.Statement {
	panic(implementation.Name()+".EndFor() Unimplemented")
	return language.Statement("")
}

