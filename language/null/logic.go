package Null

import "fmt"
import "github.com/qlova/script/language"

//Returns a Boolean that is the logical and of 'a' and 'b'.
func (l *implementation) And(a, b language.Boolean) language.Boolean {
	panic("Error in "+Name+".And(Boolean, Boolean): Unimplemented")
	return nil
}

//Returns a Boolean that is the logical or of 'a' and 'b'.
func (l *implementation) Or(a, b language.Boolean) language.Boolean {
	panic("Error in "+Name+".Or(Boolean, Boolean): Unimplemented")
	return nil
}

//Returns a Boolean that is the logical not of 'a'.
func (l *implementation) Not(a language.Boolean) language.Boolean {
	panic("Error in "+Name+".Not(Boolean): Unimplemented")
	return nil
}


func (l *implementation) If(condition language.Boolean) language.Statement {
	panic("Error in "+Name+".If(switch): Unimplemented")
	return ""
}

//Returns a Statement that begins an elseif.
func (l *implementation) ElseIf(condition language.Boolean) language.Statement {
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

func (l *implementation) Equals(a, b language.Type) language.Boolean {
	panic("Error in "+Name+".Equals("+a.Name()+", "+b.Name()+"): Unimplemented")
	return nil
}
func (l *implementation) Smaller(a, b language.Type) language.Boolean {
	panic("Error in "+Name+".EndIf("+a.Name()+", "+b.Name()+"): Unimplemented")
	return nil
}
func (l *implementation) Greater(a, b language.Type) language.Boolean {
	panic("Error in "+Name+".EndIf("+a.Name()+", "+b.Name()+"): Unimplemented")
	return nil
}
