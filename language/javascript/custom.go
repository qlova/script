package Javascript

import "github.com/qlova/script/language"


//Returns a statement that defines a Custom type 'name' with 'tokens' corresponding to 'elements'.
func (l *implementation) Type(name string, tokens []string, elements []language.Type)  language.Statement {
	panic("Error in "+Name+".Type("+name+", []string, []Type): Unimplemented")
	return ""
}

//Returns a Statement that begins the method 'name' on Custom 'T' with 'arguments' and 'returns'.
func (l *implementation) Method(T language.Custom, name string, arguments []language.Type, returns language.Type) language.Statement {
	panic("Error in "+Name+".Method(Custom, "+name+", []Type): Unimplemented")
	return ""
}

//Returns a string representing the variable pointing to the variable of type 'T' that the current Method is acting on.
func (l *implementation) This(T language.Custom) string {
	return "this"
}

//Returns the Type resuting from calling the method 'name' on Custom 'T' with 'arguments'.
func (l *implementation) CallMethod(name string, arguments []language.Type, T language.Custom) language.Type {
	panic("Error in "+Name+".CallMethod("+name+", []Type, Custom): Unimplemented")
	return nil
}

//Returns a Statement that runs the method 'name' on Custom 'T' with 'arguments'.
func (l *implementation) RunMethod(name string, arguments []language.Type, T language.Custom) language.Statement {
	panic("Error in "+Name+".RunMethod("+name+", []Type, Custom): Unimplemented")
	return ""
}

//Returns a Statement that closes the last method.
func (l *implementation) EndMethod() language.Statement {
	panic("Error in "+Name+".EndMethod(): Unimplemented")
	return ""
}
