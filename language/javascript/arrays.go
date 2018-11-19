package Javascript

import "fmt"
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

type List struct {
	language.ListType
}

//Returns a Number representing the length of 'list'
func (l *implementation) Length(list language.Type) language.Number {
	l.Import(NumberImport)
	l.AddHelper(NumberTypeDefinition)
	
	var result = Number{}
	result.Expression = "Number{Small:int64(len("+l.GetExpression(list)+"))}"
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
func (l *implementation) Array(T language.Type, length int) language.Array {
	var result Array
	
	result.Expression = "["+fmt.Sprint(length)+"]"+l.GoTypeOf(T)+"{}"
	result.Size = length
	result.Subtype = T
	
	return result
}


//Returns a Number representing the length of 'array'
func (l *implementation) Fill(T language.Type, elements []language.Type) language.Type {
	
	//Fill an array.
	if array, ok := T.(Array); ok {
		if array.Full {
			panic("Error in "+Name+".Fill("+T.Name()+", []language.Type): Cannot fill something that is already full!")
		}
		
		var result = "["+fmt.Sprint(array.Length())+"]"+l.GoTypeOf(array.SubType())+"{"
		
		for i := range elements {
			result += l.GetExpression(elements[i])
			if i < len(elements)-1 {
				result += ","
			}
		}
		result += "}"
		
		var a Array
		
		a.Expression = result
		a.Size = array.Length()
		a.Subtype = array.SubType()
		a.Full = true
		
		return a
	}
	
	panic("Error in "+Name+".Fill("+T.Name()+", []language.Type): Unimplemented")
	return nil
}
