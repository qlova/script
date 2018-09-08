package Interpreter

import "github.com/qlova/script/language"
import "github.com/qlova/script/interpreter/internal"

type Boolean struct {
	language.BooleanType
	
	Address internal.BooleanAddress
	
	internal.Variable
}

func (Boolean) SameAs(i interface{}) bool {_, ok := i.(Boolean); return ok } 

//Returns a Boolean that the Go style literal represents (true false).
func (l *implementation) LiteralBoolean(literal bool) language.Boolean {
	var b Boolean
	b.Literal = &literal
	return b
}

//Returns a Boolean that is the logical and of 'a' and 'b'.
func (l *implementation) And(a, b language.Boolean) language.Boolean {
	A, B := a.(Boolean), b.(Boolean)
	if A.Literal != nil && B.Literal != nil {
		var b Boolean
		var result = *A.Literal && *B.Literal
		b.Literal = &result
		return b
	}
	
	Block := l.loadBlock()
	var Address = Block.CreateBoolean()
	var boolean = Boolean{}
	boolean.Address = Address
	boolean.BlockPointer = Block 
	
	if A.Literal == nil && B.Literal == nil {
		Block.AddInstruction(func() {
			Block.SetBoolean(Address, A.BlockPointer.GetBoolean(A.Address) && B.BlockPointer.GetBoolean(B.Address))
		})
		return boolean
	}
	
	if A.Literal == nil && B.Literal != nil {
		var literal = *B.Literal
		Block.AddInstruction(func() {
			Block.SetBoolean(Address, A.BlockPointer.GetBoolean(A.Address) && literal)
		})
		return boolean
	}
	
	if A.Literal != nil && B.Literal == nil {
		var literal = *A.Literal
		Block.AddInstruction(func() {
			Block.SetBoolean(Address, literal && B.BlockPointer.GetBoolean(B.Address))
		})
		return boolean
	}
	
	panic("Error in "+Name+".Add(Number, Number): Unimplemented")
	return nil
}

//Returns a Boolean that is the logical or of 'a' and 'b'.
func (l *implementation) Or(a, b language.Boolean) language.Boolean {
	A, B := a.(Boolean), b.(Boolean)
	if A.Literal != nil && B.Literal != nil {
		var b Boolean
		var result = *A.Literal || *B.Literal
		b.Literal = &result
		return b
	}
	
	Block := l.loadBlock()
	var Address = Block.CreateBoolean()
	var boolean = Boolean{}
	boolean.Address = Address
	boolean.BlockPointer = Block 
	
	if A.Literal == nil && B.Literal == nil {
		Block.AddInstruction(func() {
			Block.SetBoolean(Address, A.BlockPointer.GetBoolean(A.Address) || B.BlockPointer.GetBoolean(B.Address))
		})
		return boolean
	}
	
	if A.Literal == nil && B.Literal != nil {
		var literal = *B.Literal
		Block.AddInstruction(func() {
			Block.SetBoolean(Address, A.BlockPointer.GetBoolean(A.Address) || literal)
		})
		return boolean
	}
	
	if A.Literal != nil && B.Literal == nil {
		var literal = *A.Literal
		Block.AddInstruction(func() {
			Block.SetBoolean(Address, literal || B.BlockPointer.GetBoolean(B.Address))
		})
		return boolean
	}
	
	panic("Error in "+Name+".Add(Number, Number): Unimplemented")
	return nil
}

//Returns a Boolean that is the logical not of 'a'.
func (l *implementation) Not(a language.Boolean) language.Boolean {
	A := a.(Boolean)
	if A.Literal != nil {
		var b Boolean
		var result = ! *A.Literal
		b.Literal = &result
		return b
	}
	
	Block := l.loadBlock()
	var Address = Block.CreateBoolean()
	var boolean = Boolean{}
	boolean.Address = Address
	boolean.BlockPointer = Block 
	
	if A.Literal == nil {
		Block.AddInstruction(func() {
			Block.SetBoolean(Address, !A.BlockPointer.GetBoolean(A.Address))
		})
		return boolean
	}
	
	panic("Error in "+Name+".Add(Number, Number): Unimplemented")
	return nil
}

func (l *implementation) If(condition language.Boolean) language.Statement {
	//println("IF")
	
	Block := l.loadBlock()
	l.pushBlock(Block)
	
	//This is tricky.
	NextBlock := new(internal.Block)
	l.BlockPointer = NextBlock
	l.pushBlock(NextBlock)
	
	A := condition.(Boolean)
	if A.Literal != nil {
		panic("Error in "+Name+".If(Literal): Unimplemented")
	} else {
		
		l.pushAddresses(A.Address)
		return ""
	}
	
	panic("Error in "+Name+".If(switch): Unimplemented")
	return ""
}

//Returns a Statement that begins an elseif.
func (l *implementation) ElseIf(condition language.Boolean) language.Statement {
	
	//println("ELSEIF")
	
	NextBlock := new(internal.Block)
	l.BlockPointer = NextBlock
	l.pushBlock(NextBlock)
	
	A := condition.(Boolean)
	if A.Literal != nil {
		panic("Error in "+Name+".ElseIf(switch): Unimplemented")
	} else {
		
		l.addAddress(A.Address)
		return ""
	}
	
	panic("Error in "+Name+".ElseIf(switch): Unimplemented")
	return ""
}
		
//Returns a Statement that begins an else.
func (l *implementation) Else() language.Statement {
	//println("ELSE")
	
	NextBlock := new(internal.Block)
	l.BlockPointer = NextBlock
	l.pushBlock(NextBlock)
	
	//Add a nil to the addresses array at the top of the stack.
	//This signifies that this is an else.
	
	l.addAddress(nil)

	return ""
}
		
//Returns a Statement that ends an if/else.
func (l *implementation) EndIf() language.Statement {
	
	//println("ENDIF")
	
	//These are the addresses for each boolean that represents a condition.
	addresses := l.popAddresses()
	
	//These are the blocks of each part of the if statement.
	ComponentBlocks := make([]*internal.Block, len(addresses))
	for i := len(addresses)-1; i >= 0; i-- {
		ComponentBlocks[i] = l.popBlock()
	}
	
	OrginalBlock := l.popBlock()
	
	//println(len(addresses))
	
	Instructions := make([]internal.Instruction, len(addresses))
	
	//Construct functions. unwind the ifelse chain.
	for i := len(addresses)-1; i >= 0; i-- {
		var j = i
		
		//Else or If is the final statement?
		if j == len(addresses)-1 {
			
			//It is an if / elseif.
			if (addresses[j] != nil) {
				
				var Address = addresses[j].(internal.BooleanAddress)
				var Target = ComponentBlocks[j]
				Instructions[j] = func() {
					if OrginalBlock.GetBoolean(Address) {
						l.JumpToBlock(Target)
					}
				}
			
			//It is an else.
			} else {
				var Target = ComponentBlocks[j]
				Instructions[j] = func() {
					l.JumpToBlock(Target)
				}
			}

		} else {
			
		
			var Address = addresses[j].(internal.BooleanAddress)
			var ElseInstruction = Instructions[j+1]
			var Target = ComponentBlocks[j]
			Instructions[j] = func() {
				if OrginalBlock.GetBoolean(Address) {
					l.JumpToBlock(Target)
				} else {
					ElseInstruction()
				}
			}
			
		}
		
	}
	
	//Add the entire ifelse block as one instruction.
	OrginalBlock.AddInstruction(Instructions[0])
	
	l.BlockPointer = OrginalBlock
	
	//println(len(addresses))
	
	//panic("Error in "+Name+".EndIf(): Unimplemented")
	return ""
}


func (l *implementation) Equals(a, b language.Type) language.Boolean {
	
	if !a.SameAs(b) {
		panic("Error in "+Name+".Equals("+a.Name()+", "+b.Name()+"): Not the same type")
		return nil
	}
	
	switch a.(type) {
		
		case language.Number:
			A, B := a.(Number), b.(Number)
			if A.Literal != nil && B.Literal != nil {
				var b Boolean
				var result = A.Literal.Cmp(B.Literal) == 0
				b.Literal = &result
				return b
			}
			
			Block := l.loadBlock()
			var Address = Block.CreateBoolean()
			var boolean = Boolean{}
			boolean.Address = Address
			boolean.BlockPointer = Block 
			
			if A.Literal == nil && B.Literal == nil {
				Block.AddInstruction(func() {
					Block.SetBoolean(Address, A.BlockPointer.GetNumber(A.Address).Cmp(B.BlockPointer.GetNumber(B.Address)) == 0)
				})
				return boolean
			}
			
			if A.Literal == nil && B.Literal != nil {
				var literal = *B.Literal
				Block.AddInstruction(func() {
					Block.SetBoolean(Address, A.BlockPointer.GetNumber(A.Address).Cmp(&literal) == 0)
				})
				return boolean
			}
			
			if A.Literal != nil && B.Literal == nil {
				var literal = *A.Literal
				Block.AddInstruction(func() {
					Block.SetBoolean(Address, literal.Cmp(B.BlockPointer.GetNumber(B.Address)) == 0)
				})
				return boolean
			}
	}
	
	panic("Error in "+Name+".Equals("+a.Name()+", "+b.Name()+"): Unimplemented")
	return nil
}
func (l *implementation) Smaller(a, b language.Type) language.Boolean {
	
	if !a.SameAs(b) {
		panic("Error in "+Name+".Equals("+a.Name()+", "+b.Name()+"): Not the same type")
		return nil
	}
	
	switch a.(type) {
		
		case language.Number:
			A, B := a.(Number), b.(Number)
			if A.Literal != nil && B.Literal != nil {
				var b Boolean
				var result = A.Literal.Cmp(B.Literal) == 0
				b.Literal = &result
				return b
			}
			
			Block := l.loadBlock()
			var Address = Block.CreateBoolean()
			var boolean = Boolean{}
			boolean.Address = Address
			boolean.BlockPointer = Block 
			
			if A.Literal == nil && B.Literal == nil {
				Block.AddInstruction(func() {
					Block.SetBoolean(Address, A.BlockPointer.GetNumber(A.Address).Cmp(B.BlockPointer.GetNumber(B.Address)) == -1)
				})
				return boolean
			}
			
			if A.Literal == nil && B.Literal != nil {
				var literal = *B.Literal
				Block.AddInstruction(func() {
					Block.SetBoolean(Address, A.BlockPointer.GetNumber(A.Address).Cmp(&literal) == -1)
				})
				return boolean
			}
			
			if A.Literal != nil && B.Literal == nil {
				var literal = *A.Literal
				Block.AddInstruction(func() {
					Block.SetBoolean(Address, literal.Cmp(B.BlockPointer.GetNumber(B.Address)) == -1)
				})
				return boolean
			}
	}
	
	panic("Error in "+Name+".EndIf("+a.Name()+", "+b.Name()+"): Unimplemented")
	return nil
}
func (l *implementation) Greater(a, b language.Type) language.Boolean {
	
	if !a.SameAs(b) {
		panic("Error in "+Name+".Equals("+a.Name()+", "+b.Name()+"): Not the same type")
		return nil
	}
	
	switch a.(type) {
		
		case language.Number:
			A, B := a.(Number), b.(Number)
			if A.Literal != nil && B.Literal != nil {
				var b Boolean
				var result = A.Literal.Cmp(B.Literal) == 0
				b.Literal = &result
				return b
			}
			
			Block := l.loadBlock()
			var Address = Block.CreateBoolean()
			var boolean = Boolean{}
			boolean.Address = Address
			boolean.BlockPointer = Block 
			
			if A.Literal == nil && B.Literal == nil {
				Block.AddInstruction(func() {
					Block.SetBoolean(Address, A.BlockPointer.GetNumber(A.Address).Cmp(B.BlockPointer.GetNumber(B.Address)) == 1)
				})
				return boolean
			}
			
			if A.Literal == nil && B.Literal != nil {
				var literal = *B.Literal
				Block.AddInstruction(func() {
					Block.SetBoolean(Address, A.BlockPointer.GetNumber(A.Address).Cmp(&literal) == 1)
				})
				return boolean
			}
			
			if A.Literal != nil && B.Literal == nil {
				var literal = *A.Literal
				Block.AddInstruction(func() {
					Block.SetBoolean(Address, literal.Cmp(B.BlockPointer.GetNumber(B.Address)) == 1)
				})
				return boolean
			}
	}
	
	panic("Error in "+Name+".EndIf("+a.Name()+", "+b.Name()+"): Unimplemented")
	return nil
}
