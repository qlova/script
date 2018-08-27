package Null

import . "github.com/qlova/script/language"


//Returns a Statement that begins an infinite loop.
func (l *language) Loop() Statement {
	panic("Error in "+Name+".Loop(): Unimplemented")
	return ""
}

//Returns a Statement that ends an infinite loop, do not confuse this with Break().
func (l *language) EndLoop() Statement {
	panic("Error in "+Name+".EndLoop(): Unimplemented")
	return ""
}

//Returns a Statement that breaks an infinite loop.
func (l *language) Break() Statement {
	panic("Error in "+Name+".Break(): Unimplemented")
	return ""	
}

//Returns a Statement that begins a while loop dependent on 'condition' being non zero.
func (l *language) While(condition Switch) Statement {
	panic("Error in "+Name+".While(switch): Unimplemented")
	return ""
}

//Returns a Statement that ends a while loop.
func (l *language) EndWhile() Statement {
	panic("Error in "+Name+".EndWhile(): Unimplemented")
	return ""
}

//Returns a Statement that begins a for loop that iterates along the range between 'a' and 'b'.
func (l *language) ForRange(a, b Number) Statement {
	panic("Error in "+Name+".ForRange(Number, Number): Unimplemented")
	return ""	
}

//Returns a Statement that ends a ranged for loop.
func (l *language) EndForRange() Statement {
	panic("Error in "+Name+".EndForRange(): Unimplemented")
	return ""		
}

//Returns a Statement that begins an iteration over List 'list', setting 'i' to the index and 'v' to the value at that index.
func (l *language) ForEach(i string, v string, list List) Statement {
	panic("Error in "+Name+".ForEach("+i+", "+v+", List): Unimplemented")
	return ""		
}

//Returns a Statement that ends an iteration over List
func (l *language) EndForEach() Statement {
	panic("Error in "+Name+".EndForEach(): Unimplemented")
	return ""		
}
