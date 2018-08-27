package Go

import "github.com/qlova/script/language"


//Returns a FunctionType based on 'function'.
func (l *implementation) LiteralFunctionType(function language.Function) language.FunctionType {
	panic("Error in "+Name+".LiteralFunctionType(Function): Unimplemented")
	return nil
}

//Returns the resulting Type from calling 'function' with 'arguments'.
func (l *implementation) CallFunctionType(function language.FunctionType, arguments []language.Type) language.Type {
	panic("Error in "+Name+".CallFunctionType(Function, []Type): Unimplemented")
	return nil
}

//Returns a statement that calls the FunctionType 'function' with 'arguments'.
func (l *implementation) RunFunctionType(function language.FunctionType, arguments []language.Type) language.Statement {
	panic("Error in "+Name+". RunFunctionType(FunctionType, []Type): Unimplemented")
	return ""
}

//Returns a Stream connected to 'function' called with 'arguments' starting on another thread, netowork, coroutine or process.
func (l *implementation) Thread(functiom language.Function, arguments []language.Type) language.Stream {
	panic("Error in "+Name+".Thread(Function, []Type): Unimplemented")
	return nil
}

//Returns a Statement that begins the function 'name' with 'arguments' and 'returns'.
func (l *implementation) Function(name string, arguments []language.Type, returns language.Type) language.Statement {
	panic("Error in "+Name+".Function("+name+", []Type, []Type): Unimplemented")
	return ""
}

//Returns a Statement that closes the last function.
func (l *implementation) EndFunction() language.Statement {
	panic("Error in "+Name+".EndFunction(): Unimplemented")
	return ""
}
		
//Returns the resulting Type from calling 'function' with 'arguments'
func (l *implementation) Call(function language.Function, arguments []language.Type) language.Type {
	panic("Error in "+Name+".Call(Function, []Type): Unimplemented")
	return nil
}
		
//Returns a Statement that calls 'function' with 'arguments'.
func (l *implementation) Run(functiom language.Function, arguments []language.Type) language.Statement {
	panic("Error in "+Name+".Run(Function, []Type): Unimplemented")
	return ""
}

//Returns a Statement that returns T from the current function.
func (l *implementation) Return(T language.Type) language.Statement {
	panic("Error in "+Name+".Function(Type): Unimplemented")
	return ""
}
