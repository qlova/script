package Interpreter

import "github.com/qlova/script/language"


//Returns a Number that the Go style literal represents (true false).
func (l *implementation) LiteralSwitch(literal string) language.Switch {
	panic("Error in "+Name+".LiteralSwitch("+literal+"): Unimplemented")
	return nil
}

//Returns a Switch that is the logical and of 'a' and 'b'.
func (l *implementation) And(a, b language.Switch) language.Switch {
	panic("Error in "+Name+".And(Switch, Switch): Unimplemented")
	return nil
}

//Returns a Switch that is the logical or of 'a' and 'b'.
func (l *implementation) Or(a, b language.Switch) language.Switch {
	panic("Error in "+Name+".Or(Switch, Switch): Unimplemented")
	return nil
}

//Returns a Switch that is the logical not of 'a'.
func (l *implementation) Not(a language.Switch) language.Switch {
	panic("Error in "+Name+".Not(Switch): Unimplemented")
	return nil
}

//Returns a Switch that is the logical xor of 'a' and 'b'.
func (l *implementation) Xor(a, b language.Number) language.Switch {
	panic("Error in "+Name+".Xor(Switch, Switch): Unimplemented")
	return nil
}


func (l *implementation) If(condition language.Switch) language.Statement {
	panic("Error in "+Name+".If(switch): Unimplemented")
	return ""
}

//Returns a Statement that begins an elseif.
func (l *implementation) ElseIf(condition language.Switch) language.Statement {
	panic("Error in "+Name+".ElseIf(switch): Unimplemented")
	return ""
}
		
//Returns a Statement that begins an else.
func (l *implementation) Else() language.Statement {
	panic("Error in "+Name+".Else(): Unimplemented")
	return ""
}
		
//Returns a Statement that ends an if/else.
func (l *implementation) EndIf() language.Statement {
	panic("Error in "+Name+".EndIf(): Unimplemented")
	return ""
}
