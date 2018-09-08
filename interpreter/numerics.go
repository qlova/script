package Interpreter

import "math/big"
import "github.com/qlova/script/language"
import "github.com/qlova/script/interpreter/internal"

type Number struct {
	internal.Variable
	
	Address internal.NumberAddress
	Literal *big.Int
}

func (Number) Name() string { return "number" }
func (Number) SameAs(i interface{}) bool { _, ok := i.(Number); return ok }
func (Number) Number() {}


//Returns a Number that the Go style literal represents (01 1 0x1).
func (l *implementation) LiteralNumber(literal *big.Int) language.Number {
	return Number{Literal: literal}
}

//Returns a Number that is the sum of 'a' and 'b'.
func (l *implementation) Add(a, b language.Number) language.Number {
	
	A, B := a.(Number), b.(Number)
	if A.Literal != nil && B.Literal != nil {
		var z big.Int
		z.Add(A.Literal, B.Literal)
		return Number{Literal: &z}
	}
	
	Block := l.loadBlock()
	var Address = Block.CreateNumber()
	var number = Number{}
	number.Address = Address
	number.BlockPointer = Block 
	
	if A.Literal == nil && B.Literal == nil {
		
		Block.AddInstruction(func() {
			var z big.Int
			z.Add(A.BlockPointer.GetNumber(A.Address), B.BlockPointer.GetNumber(B.Address))
			Block.SetNumber(Address, &z)
		})
		
		return number
	}
	
	if A.Literal == nil && B.Literal != nil {
		

		Block.AddInstruction(func() {
			var z big.Int
			z.Add(A.BlockPointer.GetNumber(A.Address), B.Literal)
			Block.SetNumber(Address, &z)
		})
		
		return number
	}
	
	if A.Literal != nil && B.Literal == nil {
		
		Block.AddInstruction(func() {
			var z big.Int
			z.Add(A.Literal, B.BlockPointer.GetNumber(B.Address))
			Block.SetNumber(Address, &z)
		})
		
		return number
	}
	
	panic("Error in "+Name+".Add(Number, Number): Unimplemented")
	return nil
}

//Returns a Number that is the difference of 'a' and 'b'.
func (l *implementation) Sub(a, b language.Number) language.Number {
	
	A, B := a.(Number), b.(Number)
	if A.Literal != nil && B.Literal != nil {
		var z big.Int
		z.Sub(A.Literal, B.Literal)
		return Number{Literal: &z}
	}
	
	Block := l.loadBlock()
	var Address = Block.CreateNumber()
	var number = Number{}
	number.Address = Address
	number.BlockPointer = Block 
	
	if A.Literal == nil && B.Literal == nil {
		
		Block.AddInstruction(func() {
			var z big.Int
			z.Sub(A.BlockPointer.GetNumber(A.Address), B.BlockPointer.GetNumber(B.Address))
			Block.SetNumber(Address, &z)
		})
		
		return number
	}
	
	if A.Literal == nil && B.Literal != nil {
		

		Block.AddInstruction(func() {
			var z big.Int
			z.Sub(A.BlockPointer.GetNumber(A.Address), B.Literal)
			Block.SetNumber(Address, &z)
		})
		
		return number
	}
	
	if A.Literal != nil && B.Literal == nil {
		
		Block.AddInstruction(func() {
			var z big.Int
			z.Sub(A.Literal, B.BlockPointer.GetNumber(B.Address))
			Block.SetNumber(Address, &z)
		})
		
		return number
	}
	
	panic("Error in "+Name+".Add(Number, Number): Unimplemented")
	return nil
}

//Returns a Number that is the product of 'a' and 'b'.
func (l *implementation) Mul(a, b language.Number) language.Number {
	
	A, B := a.(Number), b.(Number)
	if A.Literal != nil && B.Literal != nil {
		var z big.Int
		z.Mul(A.Literal, B.Literal)
		return Number{Literal: &z}
	}
	
	Block := l.loadBlock()
	var Address = Block.CreateNumber()
	var number = Number{}
	number.Address = Address
	number.BlockPointer = Block 
	
	if A.Literal == nil && B.Literal == nil {
		
		Block.AddInstruction(func() {
			var z big.Int
			z.Mul(A.BlockPointer.GetNumber(A.Address), B.BlockPointer.GetNumber(B.Address))
			Block.SetNumber(Address, &z)
		})
		
		return number
	}
	
	if A.Literal == nil && B.Literal != nil {
		

		Block.AddInstruction(func() {
			var z big.Int
			z.Mul(A.BlockPointer.GetNumber(A.Address), B.Literal)
			Block.SetNumber(Address, &z)
		})
		
		return number
	}
	
	if A.Literal != nil && B.Literal == nil {
		
		Block.AddInstruction(func() {
			var z big.Int
			z.Mul(A.Literal, B.BlockPointer.GetNumber(B.Address))
			Block.SetNumber(Address, &z)
		})
		
		return number
	}
	
	panic("Error in "+Name+".Mul(Number, Number): Unimplemented")
	return nil
}

//Returns a Number that is the quotient of 'a' and 'b'.
func (l *implementation) Div(a, b language.Number) language.Number {
	
	A, B := a.(Number), b.(Number)
	if A.Literal != nil && B.Literal != nil {
		var z big.Int
		z.Div(A.Literal, B.Literal)
		return Number{Literal: &z}
	}
	
	Block := l.loadBlock()
	var Address = Block.CreateNumber()
	var number = Number{}
	number.Address = Address
	number.BlockPointer = Block 
	
	if A.Literal == nil && B.Literal == nil {
		
		Block.AddInstruction(func() {
			var z big.Int
			z.Div(A.BlockPointer.GetNumber(A.Address), B.BlockPointer.GetNumber(B.Address))
			Block.SetNumber(Address, &z)
		})
		
		return number
	}
	
	if A.Literal == nil && B.Literal != nil {
		

		Block.AddInstruction(func() {
			var z big.Int
			z.Div(A.BlockPointer.GetNumber(A.Address), B.Literal)
			Block.SetNumber(Address, &z)
		})
		
		return number
	}
	
	if A.Literal != nil && B.Literal == nil {
		
		Block.AddInstruction(func() {
			var z big.Int
			z.Div(A.Literal, B.BlockPointer.GetNumber(B.Address))
			Block.SetNumber(Address, &z)
		})
		
		return number
	}
	
	panic("Error in "+Name+".Add(Number, Number): Unimplemented")
	return nil
}

//Returns a Number that is 'a' taken to the power of 'b'.
func (l *implementation) Pow(a, b language.Number) language.Number {
	
	A, B := a.(Number), b.(Number)
	if A.Literal != nil && B.Literal != nil {
		var z big.Int
		z.Exp(A.Literal, B.Literal, nil)
		return Number{Literal: &z}
	}
	
	Block := l.loadBlock()
	var Address = Block.CreateNumber()
	var number = Number{}
	number.Address = Address
	number.BlockPointer = Block 
	
	if A.Literal == nil && B.Literal == nil {
		
		Block.AddInstruction(func() {
			var z big.Int
			z.Exp(A.BlockPointer.GetNumber(A.Address), B.BlockPointer.GetNumber(B.Address), nil)
			Block.SetNumber(Address, &z)
		})
		
		return number
	}
	
	if A.Literal == nil && B.Literal != nil {
		

		Block.AddInstruction(func() {
			var z big.Int
			z.Exp(A.BlockPointer.GetNumber(A.Address), B.Literal, nil)
			Block.SetNumber(Address, &z)
		})
		
		return number
	}
	
	if A.Literal != nil && B.Literal == nil {
		
		Block.AddInstruction(func() {
			var z big.Int
			z.Exp(A.Literal, B.BlockPointer.GetNumber(B.Address), nil)
			Block.SetNumber(Address, &z)
		})
		
		return number
	}
	
	panic("Error in "+Name+".Add(Number, Number): Unimplemented")
	return nil
}

//Returns a Number that is modulos of 'a' and 'b'.
func (l *implementation) Mod(a, b language.Number) language.Number {
	
	A, B := a.(Number), b.(Number)
	if A.Literal != nil && B.Literal != nil {
		var z big.Int
		z.Mod(A.Literal, B.Literal)
		return Number{Literal: &z}
	}
	
	Block := l.loadBlock()
	var Address = Block.CreateNumber()
	var number = Number{}
	number.Address = Address
	number.BlockPointer = Block 
	
	if A.Literal == nil && B.Literal == nil {
		
		Block.AddInstruction(func() {
			var z big.Int
			z.Mod(A.BlockPointer.GetNumber(A.Address), B.BlockPointer.GetNumber(B.Address))
			Block.SetNumber(Address, &z)
		})
		
		return number
	}
	
	if A.Literal == nil && B.Literal != nil {
		

		Block.AddInstruction(func() {
			var z big.Int
			z.Mod(A.BlockPointer.GetNumber(A.Address), B.Literal)
			Block.SetNumber(Address, &z)
		})
		
		return number
	}
	
	if A.Literal != nil && B.Literal == nil {
		
		Block.AddInstruction(func() {
			var z big.Int
			z.Mod(A.Literal, B.BlockPointer.GetNumber(B.Address))
			Block.SetNumber(Address, &z)
		})
		
		return number
	}
	
	panic("Error in "+Name+".Add(Number, Number): Unimplemented")
	return nil
}
