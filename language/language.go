package language

import "math/big"

var Default Interface

type Statement string

type Type interface{
	Name() string
	SameAs(interface{}) bool
}

type Number interface {
	Type
	Number()
}

type Float interface {
	Type
	Float()
}

type Switch interface {
	Type
	Switch()
}

type Symbol interface {
	Type
	Symbol() string
}

type String interface {
	Type
	String()
}

type Error interface {
	Type
	Error()
}

type List interface {
	Type
	List()
	
	SubType() Type
}

type Array interface {
	Type
	Array()
	
	SubType() Type
	Length()  int
}

type Function interface {
	Type 
	Function()
	
	Arguments() []Type
	Returns() Type
}

//Type of type.
type Metatype interface {
	Type
	Metatype()
}

type Dynamic interface {
	Type
	Dynamic()
}

type FunctionType interface {
	Type
	FunctionType()
	
	Arguments() []Type
	Returns() Type
}

type Custom interface {
	Type
	Custom()
	
	Tokens() []string
	Elements() []Type
}

type Table interface {
	Type
	Table() string
	
	SubType() Type
}

type Pointer interface {
	Type
	Pointer()
	
	SubType() Type
}

type Stream interface {
	Type
	Stream()
}

type LanguageWithFloats interface {
	Interface
	// Floats

	//Returns a Float that the Go style literal represents (0.2 etc).
	LiteralFloat(literal string) Float
	
	//Returns a Float that is the sum of 'a' and 'b'.
	AddFloat(a, b Float) Float
	
	//Returns a Float that is the difference of 'a' and 'b'.
	SubFloat(a, b Float) Float
	
	//Returns a Float that is the product of 'a' and 'b'.
	MulFloat(a, b Float) Float
	
	//Returns a Float that is the quotient of 'a' and 'b'.
	DivFloat(a, b Float) Float
	
	//Returns a Float that is 'a' taken to the power of 'b'.
	PowFloat(a, b Float) Float
	
	//TODO maybe add trig functions.
}

// Defines a Language, all methods can panic on error. Make sure to deal with this accordingly.
type Interface interface {

	// Hooks
	
		//Hook that is run before compilation.
		Init()
		
		//These hooks are executed at the end of compilation:
		// Statements are added in the respective order.
		// Body is where all other output is placed.
		Head() Statement
		Neck() Statement
		Body() Statement
		Tail() Statement
		Last() Statement
		
	//Variables

		//Returns the type and statement that defines 'name' to be of type 'value' initialised to 'value'.
 		Define(name string, value Type) (Type, Statement)
 		
 		//Returns a Statement that sets the type 'T' variable 'name' to be set to 'value'.
 		Set(name string, T Type, value Type) Statement
 		
 		//Returns the Type at 'index' of 'T'.
 		Index(T Type, index Type) Type
 		
 		//Returns a statement that modifies type T at 'index' to be 'value'.
 		Modify(T Type, index Type, value Type) Statement

	// Logic
		
		//Returns a Statement that begins an if.
		If(condition Switch) Statement
		
		//Returns a Statement that begins an elseif.
		ElseIf(condition Switch) Statement
		
		//Returns a Statement that begins an else.
		Else() Statement
		
		//Returns a Statement that ends an if/else.
		EndIf() Statement
		
	// Loops
		
		//Returns a Statement that begins an infinite loop.
		Loop() Statement
		
		//Returns a Statement that ends an infinite loop, do not confuse this with Break().
		EndLoop() Statement
		
		//Returns a Statement that breaks an infinite loop.
		Break() Statement
		
		//Returns a Statement that begins a while loop dependent on 'condition' being non zero.
		While(condition Switch) Statement
		
		//Returns a Statement that ends a while loop.
		EndWhile() Statement
		
		//Returns an iterator and Statement that begins a for loop that iterates along the range between 'a' and 'b'.
		ForRange(i string, a, b Number) (Number, Statement)
		
		//Returns a Statement that ends a ranged for loop.
		EndForRange() Statement
		
		//Returns a Statement that begins an iteration over List 'list', setting 'i' to the index and 'v' to the value at that index.
		ForEach(i string, v string, list List) Statement
		
		//Returns a Statement that ends an iteration over List
		EndForEach() Statement
	
	// Entrypoint
	
		//Returns a Statement that begins the main entry point to the program.
		Main() Statement
		
		//Returns a Statement that exits the program.
		Exit() Statement
		
		//Returns a Statement that ends the main entry point to the program.
		EndMain() Statement
	
	//Functions
		
		//Returns a Statement that begins the function 'name' with 'arguments' and 'returns'.
		Function(name string, arguments []Type, returns Type) Statement
		
		//Returns a Statement that closes the last function.
		EndFunction() Statement
		
		//Returns the resulting Type from calling 'function' with 'arguments'
		Call(function Function, arguments []Type) Type
		
		//Returns a Statement that calls 'function' with 'arguments'.
		Run(functiom Function, arguments []Type) Statement
		
		//Returns a Statement that returns T from the current function.
		Return(T Type) Statement
		
	//Threading
	
		//Returns a Stream connected to 'function' called with 'arguments' starting on another thread, netowork, coroutine or process.
		Thread(functiom Function, arguments []Type) Stream
		
	// Stdin, Stderr and Stdout
	
		//Returns a Statement that prints Types to os.Stdout with a newline.
		//TODO specify how types should be printed.
		Print(...Type) Statement
		
		//Returns a Statement that writes a String to Stream (or Stdout) without a newline.
		WriteString(Stream, String) Statement
		
		//Returns a Statement that writes the contents of Array to a Stream (or Stdout) without a newline.
		WriteArray(Stream, Array) Statement

		//Returns a statement that sends Type 't' over Stream 'c'.
		Send(c Stream, t Type) Statement
		
		//Returns Type 't' from Stream 'c'.
		Read(c Stream, t Type) Type
	
		//Reads Symbols from Stream (or Stdin) until Symbol is reached, returns a String of all Symbols up until Symbol.
		ReadSymbol(Stream, Symbol) String
		
		//Reads 'amount' bytes from Stream (or Stdin), returns Array of all Bytes up until 'amount'. 
		ReadNumber(s Stream, amount Number) Array

		//Returns a Statement that Reads bytes from Stream (or Stdin) and fills Array. 
		ReadArray(s Stream, fill Array) Statement
		
	//Streams
	
		//Returns a Stream at 'path' associated with the given 'protocol'.
		Open(protocol String, path String) Stream
		
		//Returns a String at 'path' associated with the given 'protocol'.
		Load(protocol String, path String) String
		
		//Returns a statement that stops Stream 'c'.
		Stop(c Stream) Statement
		
		//Returns a statement that seeks Stream 'c' by 'amount'.
		Seek(c Stream, amount Number) Statement
		
		//Returns a String that is the result of a 'query' on Stream 'c'.
		Info(c Stream, query String) String
		
		//Returns a Statement that moves Stream 'c' to 'location'.
		Move(c Stream, location String) Statement

	// Errors
		//Returns an Error with 'code' and 'message'.
		LiteralError(code Number, message String) Error
		
		//Returns a Statement that throws an error to the thread, this should not halt the program.
		Throw(err Error) Statement
		
		//Returns a Error that is sitting on the thread.
		Catch() Error
		
	// Switches aka Booleans
	
		//Returns a Number that the Go style literal represents (true false).
		LiteralSwitch(literal string) Switch
		
		//Returns a Switch that is the logical and of 'a' and 'b'.
		And(a, b Switch) Switch
		
		//Returns a Switch that is the logical or of 'a' and 'b'.
		Or(a, b Switch) Switch
		
		//Returns a Switch that is the logical not of 'a'.
		Not(a Switch) Switch
		
		//Returns a Switch that is the logical xor of 'a' and 'b'.
		Xor(a, b Number) Switch
	
	// Numbers
	
		//Returns a Number that the Go style literal represents (01 1 0x1).
		LiteralNumber(literal *big.Int) Number
		
		//Returns a Number that is the sum of 'a' and 'b'.
		Add(a, b Number) Number
		
		//Returns a Number that is the difference of 'a' and 'b'.
		Sub(a, b Number) Number
		
		//Returns a Number that is the product of 'a' and 'b'.
		Mul(a, b Number) Number
		
		//Returns a Number that is the quotient of 'a' and 'b'.
		Div(a, b Number) Number
		
		//Returns a Number that is 'a' taken to the power of 'b'.
		Pow(a, b Number) Number
		
		//Returns a Number that is modulos of 'a' and 'b'.
		Mod(a, b Number) Number
		
	//Symbols
	
		//Returns a Symbol that the Go style literal represents ('').
		LiteralSymbol(literal string) Symbol
	
	//Strings

		//Returns a String that the Go style literal represents ("").
		LiteralString(literal string) String
		
		//Returns a new String that concatenates 'a' and 'b'.
		JoinString(a, b String) String

	//Arrays, fixed-size collection of elements.

		//Returns an Array of type T with 'length' length.
		MakeArray(T Type, length Number) Array
		
		//Returns a Number representing the length of 'array'
		LengthArray(array Array) Number
		
	//Lists, unordered collection of elements.

		//Returns a List of type T.
		MakeList(T Type) List
		
		//Returns a List of type T, intialised with optional 'elements'.
		LiteralList(T Type, elements ...Type) List
		
		//Returns a Number representing the length of 'list'
		LengthList(list List) Number
		
		//Returns a statement that adds 't' to the 'list'.
		Grow(list List, t Type) Statement
		
		//Returns a statement that shrinks the 'list' by removing the element at the end..
		Shrink(list List) Statement
		
	//Function Types

		//Returns a FunctionType based on 'function'.
		LiteralFunctionType(function Function) FunctionType
		
		//Returns the resulting Type from calling 'function' with 'arguments'.
		CallFunctionType(function FunctionType, arguments []Type) Type
		
		//Returns a statement that calls the FunctionType 'function' with 'arguments'.
		RunFunctionType(function FunctionType, arguments []Type) Statement
	
	// Tables.
		
		//Returns a Table of type T.
		MakeTable(T Type) Table
		
		//Returns a Table intialised with 'indices' corresponding to 'elements'.
		LiteralTable(indices []String, elements []Type) Table
		
	//Custom types.

		//Returns a statement that defines a Custom type 'name' with 'tokens' corresponding to 'elements'.
		Type(name string, tokens []string, elements []Type) Statement
		
		//Returns a Statement that begins the method 'name' on Custom 'T' with 'arguments' and 'returns'.
		Method(T Custom, name string, arguments []Type, returns Type) Statement
		
		//Returns a string representing the variable pointing to the variable of type 'T' that the current Method is acting on.
		This(T Custom) string
		
		//Returns the Type resuting from calling the method 'name' on Custom 'T' with 'arguments'.
		CallMethod(name string, arguments []Type, T Custom) Type
		
		//Returns a Statement that runs the method 'name' on Custom 'T' with 'arguments'.
		RunMethod(name string, arguments []Type, T Custom) Statement
		
		//Returns a Statement that closes the last method.
		EndMethod() Statement
		
		//Returns a Custom intialised with 'tokens' corresponding to 'elements' AKA a tuple when tokens are empty.
		LiteralCustom(tokens []string, elements []Type) Custom
		
	// Pointers.
		
		//Returns a Pointer type based of value 'T'.
		PointerTo(value Type) Pointer
		
		//Returns the refernce of the Pointer 'pointer'.
		Dereference(pointer Pointer) Type
		
	//Dynamic types.
		
		//Returns a Dynamic type based of value 'T'.
		ToDynamic(value Type) Dynamic
		
		//Returns a Type cast from value 'T'.
		DynamicTo(value Type) Type
		
		//Returns Dynamic's type as a Metatype.
		DynamicMetatype(value Dynamic) Metatype
	
	//Casting
		
		//Returns Error cast to String.
		ErrorToString(Error) String
		
		//Returns Error cast to Number.
		ErrorToNumber(Error) Number
		
		//Returns Number cast to String.
		NumberToString(Number) String
		
		//Returns String cast to Number.
		StringToNumber(String) Number
		
		//Returns Symbol cast to String.
		SymbolToString(Symbol) String
		
		//Returns Symbol cast to Number.
		SymbolToNumber(Symbol) Number
		
		//Returns Symbol cast to Number.
		NumberToSymbol(Number) Symbol
}
