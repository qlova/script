package interpreter

import "github.com/qlova/script/language"
import "github.com/qlova/script/interpreter/dynamic"

type IfState struct {
	If int
	ElseIf []int
	Else int
}

func (implementation Implementation) PushIfState() {
	implementation.ifstates = append(implementation.ifstates, IfState{If:implementation.Instruction(), Else: -1})
}

func (implementation Implementation) IfState() *IfState {
	return &implementation.ifstates[len(implementation.ifstates)-1]
}

func (implementation Implementation) PopIfState() IfState {
	var result = implementation.ifstates[len(implementation.ifstates)-1]
	implementation.ifstates = implementation.ifstates[:len(implementation.ifstates)-1]
	return result
}

func (implementation Implementation) If(condition language.Bit) language.Statement {
	
	var register = implementation.RegisterOf(condition)
	
	//We create two intructions, one is the condition.
	//The second will jump to the next component of the if-else statement.
	
	implementation.PushIfState()
	
	implementation.AddInstruction(func(thread *dynamic.Thread) {
		if thread.Get(register).(bool) {
			thread.InstructionCounter++ //Skip next instruction.
		}
	})
	
	//Since we don't know where the next if-else component is, we will leave this instruction blank for now.
	implementation.AddInstruction(func(thread *dynamic.Thread) {})

	return language.Statement("")
}

func (implementation Implementation) ElseIf(condition language.Bit) language.Statement {
	panic(implementation.Name()+".ElseIf() Unimplemented")
	return language.Statement("")
}

func (implementation Implementation) Else() language.Statement {
	
	//This instruction will jump to the end of the if statement.
	implementation.AddInstruction(func(thread *dynamic.Thread) {})
	
	var state = implementation.IfState()
		state.Else = implementation.Instruction()

	return language.Statement("")
}

func (implementation Implementation) EndIf() language.Statement {
	
	var End = implementation.Instruction()
	
	var state = implementation.PopIfState()
	
	if len(state.ElseIf) > 0 {
		panic(implementation.Name()+".EndIf() Unimplemented")
	}
	
	if state.Else >= 0 {
		implementation.SetInstruction(state.If+1, func(thread *dynamic.Thread) {
			thread.InstructionCounter = state.Else-1
		})
		implementation.SetInstruction(state.Else-1, func(thread *dynamic.Thread) {
			thread.InstructionCounter = End-1
		})
	} else {
		implementation.SetInstruction(state.If+1, func(thread *dynamic.Thread) {
			thread.InstructionCounter = End-1
		})
	}
	
	return language.Statement("")
}

