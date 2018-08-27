package Null

import . "github.com/qlova/script/language"


//Returns a FunctionType based on 'function'.
func (l *language) LiteralFunctionType(function Function) FunctionType {
	panic("Error in "+Name+".LiteralFunctionType(Function): Unimplemented")
	return nil
}

//Returns the resulting Type from calling 'function' with 'arguments'.
func (l *language) CallFunctionType(function FunctionType, arguments []Type) Type {
	panic("Error in "+Name+".CallFunctionType(Function, []Type): Unimplemented")
	return nil
}

//Returns a statement that calls the FunctionType 'function' with 'arguments'.
func (l *language) RunFunctionType(function FunctionType, arguments []Type) Statement {
	panic("Error in "+Name+". RunFunctionType(FunctionType, []Type): Unimplemented")
	return ""
}

//Returns a Stream connected to 'function' called with 'arguments' starting on another thread, netowork, coroutine or process.
func (l *language) Thread(functiom Function, arguments []Type) Stream {
	panic("Error in "+Name+".Thread(Function, []Type): Unimplemented")
	return nil
}

//Returns a Statement that begins the function 'name' with 'arguments' and 'returns'.
func (l *language) Function(name string, arguments []Type, returns Type) Statement {
	panic("Error in "+Name+".Function("+name+", []Type, []Type): Unimplemented")
	return ""
}

//Returns a Statement that closes the last function.
func (l *language) EndFunction() Statement {
	panic("Error in "+Name+".EndFunction(): Unimplemented")
	return ""
}
		
//Returns the resulting Type from calling 'function' with 'arguments'
func (l *language) Call(function Function, arguments []Type) Type {
	panic("Error in "+Name+".Call(Function, []Type): Unimplemented")
	return nil
}
		
//Returns a Statement that calls 'function' with 'arguments'.
func (l *language) Run(functiom Function, arguments []Type) Statement {
	panic("Error in "+Name+".Run(Function, []Type): Unimplemented")
	return ""
}

//Returns a Statement that returns T from the current function.
func (l *language) Return(T Type) Statement {
	panic("Error in "+Name+".Function(Type): Unimplemented")
	return ""
}
