package Interpreter

import "github.com/qlova/script/language"
import "github.com/qlova/script/interpreter/internal"

import "math/big"

//Returns a new String that concatenates 'a' and 'b'.
func (l *implementation) Join(a, b language.Type) language.Type {
	block := l.loadBlock()
	
	if _, ok := a.(String); ok {
		if _, ok := b.(String); ok {
			
			var Address = block.CreateString()
			
			var BlockPointer1 = a.(String).BlockPointer
			var A = a.(String).Address
			
			var BlockPointer2 = b.(String).BlockPointer
			var B = b.(String).Address
			
			block.AddInstruction(func() {
				block.SetString(Address, BlockPointer1.GetString(A) + BlockPointer2.GetString(B))
			})
			
			var result String
			result.BlockPointer= block
			result.Address = Address
			return result
			
		}
	}
	
	
	panic("Error in "+Name+".Join("+a.Name()+", "+b.Name()+"): Unimplemented")
	return nil
}


//Returns an Array of type T with 'length' length.
func (l *implementation) LiteralArray(array interface{}) language.Array {
	panic("Error in "+Name+".MakeArray(Type, Number): Unimplemented")
	return nil
}

//Returns a Number representing the length of 'array'
func (l *implementation) Length(array language.Type) language.Number {
	block := l.loadBlock()
	
	var n = l.NewNumber()
	
	var Address = n.Address
	var Length = int64(array.(Array).Length())
	if Length == 0 {
		panic("Error in "+Name+".Length("+array.Name()+"): Unimplemented for length == 0")
	}
	
	block.AddInstruction(func() {
		block.SetNumber(Address, big.NewInt(Length))
	})

	return n
}

//Returns a List of type T.
func (l *implementation) MakeList(T language.Type) language.List {
	panic("Error in "+Name+".MakeList(Type): Unimplemented")
	return nil
}

//Returns a List of type T, intialised with optional 'elements'.
func (l *implementation) LiteralList(T language.Type, elements ...language.Type) language.List {
	panic("Error in "+Name+".LiteralList(Type, ...Type): Unimplemented")
	return nil
}

//Returns a statement that adds 't' to the 'list'.
func (l *implementation) Grow(list language.List, t language.Type) language.Statement {
	panic("Error in "+Name+".Grow(List, Type): Unimplemented")
	return ""
}

//Returns a statement that shrinks the 'list' by removing the element at the end..
func (l *implementation) Shrink(list language.List) language.Statement {
	panic("Error in "+Name+".Shrink(List): Unimplemented")
	return ""
}

type Array struct {
	language.ArrayType
	internal.Variable
	
	Address internal.ArrayAddress
	
	Literal []language.Type
}

//Returns a Number representing the length of 'array'
func (l *implementation) Array(T language.Type, length int) language.Array {
	var result = Array{}
	result.Size = length
	result.Subtype = T
	return result
}


//Returns a filled type of type T with 'elements'
func (l *implementation) Fill(T language.Type, elements []language.Type) language.Type {
	if len(elements) != T.(Array).Length() {
		panic("Error in "+Name+".Fill("+T.Name()+", []language.Type): Invalid Length")
	}
	
	var a = T.(Array)
	a.Literal = elements
	return a
}
