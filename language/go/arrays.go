package Go

import "math/big"
import "github.com/qlova/script/language"

//Returns a new String that concatenates 'a' and 'b'.
func (l *implementation) Join(a, b language.Type) language.Type {
	
	if _, ok := a.(String); ok {
		if _, ok := b.(String); ok {
			
			return String(l.GetExpression(a)+" + "+l.GetExpression(b))
			
		}
	}
	
	panic("Error in "+Name+".Join("+a.Name()+", "+b.Name()+"): Unimplemented")
	return nil
}


type Array struct {
	language.ArrayType
}

//Returns a Number representing the length of 'list'
func (l *implementation) Length(list language.Type) language.Number {
	var result = Number{}
	if array, ok := list.(Array); ok {
		
		result.Literal = big.NewInt(int64(array.Length()))
		
	} else {
		result.Expression = "len("+l.GetExpression(list)+")"
	}
	return result
}

//Returns a statement that adds 't' to the 'list'.
func (l *implementation) Grow(list language.List, t language.Type) language.Statement {
	var expression = l.GetExpression(list)
	return language.Statement(expression+" = append("+expression+", "+l.GetExpression(t)+")")
}

//Returns a statement that shrinks the 'list' by removing the element at the end..
func (l *implementation) Shrink(list language.List) language.Statement {
	var expression = l.GetExpression(list)
	return language.Statement(expression+" = "+expression+"[:len( "+expression+")-1]")
}

//Returns a Number representing the length of 'array'
func (l *implementation) Array(T language.Type, length language.Number) language.Array {
	if (length.(Number).Literal == nil) {
		panic("Error in "+Name+".Array("+T.Name()+", []language.Type): length is not a literal.")
	}
	
	var result Array
	
	result.Expression = "["+length.(Number).Literal.String()+"]"+l.GoTypeOf(T)+"{}"
	result.Size = int(length.(Number).Literal.Int64())
	result.Subtype = T
	
	return result
}


//Returns a Number representing the length of 'array'
func (l *implementation) Fill(T language.Type, elements []language.Type) language.Type {
	
	panic("Error in "+Name+".Fill("+T.Name()+", []language.Type): Unimplemented")
	return nil
}
