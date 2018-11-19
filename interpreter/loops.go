package Interpreter

import "math/big"
import "github.com/qlova/script/language"

//Returns a Statement that begins an infinite loop.
func (l *implementation) Loop() language.Statement {
	panic("Error in "+Name+".Loop(): Unimplemented")
	return ""
}

//Returns a Statement that ends an infinite loop, do not confuse this with Break().
func (l *implementation) EndLoop() language.Statement {
	panic("Error in "+Name+".EndLoop(): Unimplemented")
	return ""
}

//Returns a Statement that breaks an infinite loop.
func (l *implementation) Break() language.Statement {
	panic("Error in "+Name+".Break(): Unimplemented")
	return ""	
}

//Returns a Statement that begins a while loop dependent on 'condition' being non zero.
func (l *implementation) While(condition language.Boolean) language.Statement {
	panic("Error in "+Name+".While(switch): Unimplemented")
	return ""
}

//Returns a Statement that ends a while loop.
func (l *implementation) EndWhile() language.Statement {
	panic("Error in "+Name+".EndWhile(): Unimplemented")
	return ""
}

//Returns a Statement that begins a for loop that iterates along the range between 'a' and 'b'.
func (l *implementation) ForRange(i string, a, b language.Number) (language.Number, language.Statement) {
	var Block = l.loadBlock()
	var Address = Block.CreateNumber() //This is the address of the iterator.
	
	A, B := a.(Number), b.(Number)
	
	if A.Literal != nil {
		var literal = *A.Literal
		Block.AddInstruction(func() {
			Block.SetNumber(Address, &literal)
		})
	} else {
		panic("Error in "+Name+".ForRange(Number, Number): Unimplemented")
	}
	
	Block.PushPointer()
	
	if B.Literal != nil {
		var literal = *B.Literal
		
		Block.AddInstruction(func() {
			var number = Block.GetNumber(Address)
			if number.Cmp(&literal) == -1 {
				number.Add(number, big.NewInt(1))
			} else {
				l.InstructionPointer = l.BreakPoint
			}
		})
	} else {
		panic("Error in "+Name+".ForRange(Number, Number): Unimplemented")
	}

	var number = Number{}
	number.BlockPointer = Block
	number.Address = Address
	
	return number, ""	
}

//Returns a Statement that ends a ranged for loop.
func (l *implementation) EndForRange() language.Statement {
	var Block = l.loadBlock()
	var Pointer = Block.PopPointer()
	
	var Break = len(Block.Instructions)+1
	
	if Block.Instructions[Pointer] != nil {
		panic("Error in EndForRange(), Invalid loop")
	}
	
	Block.Instructions[Pointer] = func() {
		l.BreakPoint = Break
	}
	
	Block.AddInstruction(func() {
		l.InstructionPointer = Pointer
	})
	
	return ""		
}

//Returns a Statement that begins an iteration over List 'list', setting 'i' to the index and 'v' to the value at that index.
func (l *implementation) ForEach(i string, v string, list language.Type) (language.Number, language.Type, language.Statement) {

	var Block = l.loadBlock()
	
	var value language.Type
	
	switch li := list.(type) {
		case language.Array:
			value, _ = l.GetVariable(v, li.SubType())
		case language.List:
			value, _ = l.GetVariable(v, li.SubType())
		default:
			panic("Error in "+Name+".ForEach("+i+", "+v+", "+list.Name()+"): Unimplemented")
	}

	
	var number = Number{}
	var Address = Block.CreateNumber()
	number.BlockPointer = Block
	number.Address = Address
	Block.AddInstruction(func() {
		Block.SetNumber(Address, big.NewInt(0))
	})
	
	Block.PushPointer()
	
	switch list.(type) {
		case language.Array:
			panic("Error in "+Name+".ForRange(Number, Number): Unimplemented")
		case language.List:
			var ListAddress = list.(List).Address
			Block.AddInstruction(func() {
				var number = Block.GetNumber(Address)
				var list = Block.GetValue(ListAddress)
				if number.Cmp(big.NewInt(int64(list.Len()))) == -1 {
					
					//Set v
					
					
					number.Add(number, big.NewInt(1))
				} else {
					l.InstructionPointer = l.BreakPoint
				}
				
				
				
			})
	}
	return number, value, ""	
}

//Returns a Statement that ends an iteration over List
func (l *implementation) EndForEach() language.Statement {
	var Block = l.loadBlock()
	var Pointer = Block.PopPointer()
	
	var Break = len(Block.Instructions)+1
	
	if Block.Instructions[Pointer] != nil {
		panic("Error in EndForEach(), Invalid loop")
	}
	
	Block.Instructions[Pointer] = func() {
		l.BreakPoint = Break
	}
	
	Block.AddInstruction(func() {
		l.InstructionPointer = Pointer
	})
	return ""
}
