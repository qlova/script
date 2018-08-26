package script

type Language interface {
	
	//Hooks that run at the beginning of the compile and at the end.
	Init(*Script)
	Last(*Script)
	
	//The main entry point for the program.
	Main(*Script)
	EndMain(*Script) //Closing the main entry point.
	
	//Creating a function or subroutine
	Function(q *Script, name string, arguments []Type, returns []Type)
	EndFunction(*Script) //Closing the function.
	Call(q *Script, name string, arguments []Type) string //Calling a function.
	
	//Print a string.
	Print(*Script, String)
	
	//Strings
	NewString(*Script, string, ...String)
	SetString(*Script, string, String)
	AddStrings(string, string) string
	
	//Numbers
	LiteralNumber(string) string
	NewNumber(*Script, string, ...Number)
	SetNumber(*Script, string, Number)
	
	//Arithmetic (for numbers)
	Add(string, string) string
	Subtract(string, string) string
	Multiply(string, string) string
	Divide(string, string) string
	Power(string, string) string
	Modulo(string, string) string
	
	//Symbols
	LiteralSymbol(string) string
	NewSymbol(*Script, string, ...Symbol)
	SetSymbol(*Script, string, Symbol)
	
	//Arrays
	LiteralArray(Type, ...Type) string
	NewArray(Type, *Script, string, ...Array)
	SetArray(Type, *Script, string, Array)
	IndexArray(string, string) string
	LengthArray(string) string 
	
	NumberToArray(t Type, size string) string //Create an array of the requested size.
	
	//Lists
	LiteralList(Type, ...Type) string
	NewList(Type, *Script, string, ...List)
	SetList(Type, *Script, string, List)
	IndexList(string, string) string
	LengthList(string) string 
	
	NumberToList(t Type, size string) string //Create an array of the requested size.
	
	//Functions
	Return(*Script, Type)
	ScopeFunctionType(string) string
	NewFunctionType(*Script, string, []Type, []Type, ...FunctionType)
	SetFunctionType(*Script, string, FunctionType)
	CallFunctionType(q *Script, name string, arguments []Type) string //Calling a function variable.
	
	//Loops
	ForEachList(i string, v string, a string) string
	EndForEachList() string
	
	//Conversions
	NumberToString(string) string
	StringToNumber(string) string
	
	SymbolToString(string) string
	
	SymbolToNumber(string) string
	NumberToSymbol(string) string
	
	//Input / Output
	ReadSymbol(string) string
}
