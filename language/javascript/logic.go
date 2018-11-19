package Javascript

import "github.com/qlova/script/language"

type Boolean string

func (Boolean) Name() string { return "boolean" }
func (Boolean) SameAs(i interface{}) bool { _, ok := i.(Boolean); return ok }
func (Boolean) Boolean() {}

//Returns a Boolean that is the logical and of 'a' and 'b'.
func (l *implementation) And(a, b language.Boolean) language.Boolean {
	return Boolean("("+l.GetExpression(a)+" && "+l.GetExpression(b)+")")
}

//Returns a Boolean that is the logical or of 'a' and 'b'.
func (l *implementation) Or(a, b language.Boolean) language.Boolean {
	return Boolean("("+l.GetExpression(a)+" || "+l.GetExpression(b)+")")
}

//Returns a Boolean that is the logical not of 'a'.
func (l *implementation) Not(a language.Boolean) language.Boolean {
	return Boolean("!"+l.GetExpression(a))
}


func (l *implementation) If(condition language.Boolean) language.Statement {
	return language.Statement("if ("+l.GetExpression(condition)+") {\n")
	
	panic("Error in "+Name+".If(switch): Unimplemented")
	return ""
}

//Returns a Statement that begins an elseif.
func (l *implementation) ElseIf(condition language.Boolean) language.Statement {
	return  language.Statement("} else if "+l.GetExpression(condition)+" {\n")
	
	panic("Error in "+Name+".ElseIf(switch): Unimplemented")
	return ""
}
		
//Returns a Statement that begins an else.
func (l *implementation) Else() language.Statement {
	return  language.Statement("} else {\n")
	
	panic("Error in "+Name+".Else(): Unimplemented")
	return ""
}
		
//Returns a Statement that ends an if/else.
func (l *implementation) EndIf() language.Statement {
	return  language.Statement("}\n")
	
	panic("Error in "+Name+".EndIf(): Unimplemented")
	return ""
}


func (l *implementation) Equals(a, b language.Type) language.Boolean {
	return Boolean("("+l.GetExpression(a)+" == "+l.GetExpression(b)+")")
	
	panic("Error in "+Name+".Equals("+a.Name()+", "+b.Name()+"): Unimplemented")
	return nil
}
func (l *implementation) Smaller(a, b language.Type) language.Boolean {
	return Boolean("("+l.GetExpression(a)+".Cmp("+l.GetExpression(b)+") == -1)")
	
	panic("Error in "+Name+".EndIf("+a.Name()+", "+b.Name()+"): Unimplemented")
	return nil
}
func (l *implementation) Greater(a, b language.Type) language.Boolean {
	return Boolean("("+l.GetExpression(a)+".Cmp("+l.GetExpression(b)+") == 1)")
	
	panic("Error in "+Name+".EndIf("+a.Name()+", "+b.Name()+"): Unimplemented")
	return nil
}
