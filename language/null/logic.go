package Null

import . "github.com/qlova/script/language"


//Returns a Number that the Go style literal represents (true false).
func (l *language) LiteralSwitch(literal string) Switch {
	panic("Error in "+Name+".LiteralSwitch("+literal+"): Unimplemented")
	return nil
}

//Returns a Switch that is the logical and of 'a' and 'b'.
func (l *language) And(a, b Switch) Switch {
	panic("Error in "+Name+".And(Switch, Switch): Unimplemented")
	return nil
}

//Returns a Switch that is the logical or of 'a' and 'b'.
func (l *language) Or(a, b Switch) Switch {
	panic("Error in "+Name+".Or(Switch, Switch): Unimplemented")
	return nil
}

//Returns a Switch that is the logical not of 'a'.
func (l *language) Not(a Switch) Switch {
	panic("Error in "+Name+".Not(Switch): Unimplemented")
	return nil
}

//Returns a Switch that is the logical xor of 'a' and 'b'.
func (l *language) Xor(a, b Number) Switch {
	panic("Error in "+Name+".Xor(Switch, Switch): Unimplemented")
	return nil
}


func (l *language) If(condition Switch) Statement {
	panic("Error in "+Name+".If(switch): Unimplemented")
	return ""
}

//Returns a Statement that begins an elseif.
func (l *language) ElseIf(condition Switch) Statement {
	panic("Error in "+Name+".ElseIf(switch): Unimplemented")
	return ""
}
		
//Returns a Statement that begins an else.
func (l *language) Else() Statement {
	panic("Error in "+Name+".Else(): Unimplemented")
	return ""
}
		
//Returns a Statement that ends an if/else.
func (l *language) EndIf() Statement {
	panic("Error in "+Name+".EndIf(): Unimplemented")
	return ""
}
