package Null

import . "github.com/qlova/script/language"


//Returns a statement that defines a Custom type 'name' with 'tokens' corresponding to 'elements'.
func (l *language) Type(name string, tokens []string, elements []Type)  Statement {
	panic("Error in "+Name+".Type("+name+", []string, []Type): Unimplemented")
	return ""
}

//Returns a Statement that begins the method 'name' on Custom 'T' with 'arguments' and 'returns'.
func (l *language) Method(T Custom, name string, arguments []Type, returns Type) Statement {
	panic("Error in "+Name+".Method(Custom, "+name+", []Type): Unimplemented")
	return ""
}

//Returns a string representing the variable pointing to the variable of type 'T' that the current Method is acting on.
func (l *language) This(T Custom) string {
	panic("Error in "+Name+".This(Custom): Unimplemented")
	return ""
}

//Returns the Type resuting from calling the method 'name' on Custom 'T' with 'arguments'.
func (l *language) CallMethod(name string, arguments []Type, T Custom) Type {
	panic("Error in "+Name+".CallMethod("+name+", []Type, Custom): Unimplemented")
	return nil
}

//Returns a Statement that runs the method 'name' on Custom 'T' with 'arguments'.
func (l *language) RunMethod(name string, arguments []Type, T Custom) Statement {
	panic("Error in "+Name+".RunMethod("+name+", []Type, Custom): Unimplemented")
	return ""
}

//Returns a Statement that closes the last method.
func (l *language) EndMethod() Statement {
	panic("Error in "+Name+".EndMethod(): Unimplemented")
	return ""
}

//Returns a Custom intialised with 'tokens' corresponding to 'elements' AKA a tuple when tokens are empty.
func (l *language) LiteralCustom(tokens []string, elements []Type) Custom {
	panic("Error in "+Name+".LiteralCustom([]string, []Type): Unimplemented")
	return nil
}
