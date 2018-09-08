package Interpreter

import "github.com/qlova/script/language"
import "github.com/qlova/script/interpreter/internal"

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
	panic("Error in "+Name+".LengthArray("+array.Name()+"): Unimplemented")
	return nil
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
	Empty bool
}

//Returns a Number representing the length of 'array'
func (l *implementation) Array(T language.Type, length language.Number) language.Array {
	if (length.(Number).Literal == nil) {
		panic("Error in "+Name+".Array("+T.Name()+", []language.Type): length is not a literal.")
	}
	
	var result = Array{}
	result.Empty = true
	result.Size = int(length.(Number).Literal.Int64())
	result.Subtype = T
	return result
}


//Returns a Number representing the length of 'array'
func (l *implementation) Fill(T language.Type, elements []language.Type) language.Type {
	
	panic("Error in "+Name+".Fill("+T.Name()+", []language.Type): Unimplemented")
	return nil
}
