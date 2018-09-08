package Go

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
	var result string
	
	A, B := a.(Number), b.(Number)
	
	if A.Literal != nil && B.Literal != nil {
		
		var z big.Int
		
		//Figure out the direction.
		z.Sub(A.Literal, B.Literal) 
		cmp := z.Cmp(big.NewInt(0)) 
		
		
		if cmp == -1 {
			result = "for "+i+" := "+A.Literal.String()+"; i < "+B.Literal.String()+"; i++ {\n"
		} else if cmp == 1 {
			result = "for "+i+" := "+A.Literal.String()+"; i > "+B.Literal.String()+"; i-- {\n"
		} else {
			result = "if i := 0; false {\n"
		}
		
	} else {
		
		l.AddHelper(`func dir(i int) {
	if i > 0 {
		return 1
	} 
	return -1
}`)
		
		panic("Error in "+Name+".ForRange("+i+", Number, Number): Unimplemented")
	}
	
	return Number{Expression: i}, language.Statement(result)
}

//Returns a Statement that ends a ranged for loop.
func (l *implementation) EndForRange() language.Statement {
	return "}\n"
}

//Returns a Statement that begins an iteration over List 'list', setting 'i' to the index and 'v' to the value at that index.
func (l *implementation) ForEach(i string, v string, list language.List) language.Statement {
	panic("Error in "+Name+".ForEach("+i+", "+v+", List): Unimplemented")
	return ""		
}

//Returns a Statement that ends an iteration over List
func (l *implementation) EndForEach() language.Statement {
	panic("Error in "+Name+".EndForEach(): Unimplemented")
	return ""		
}
