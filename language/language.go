package language

type Statement interface {
	Statement() string
}

type Type interface{
	Name() string
	SameAs(Type)
}

type Number interface {
	Type
	Number() string
}

type Float interface {
	Type
	Float() string
}

type Switch interface {
	Type
	Switch() string
}

type Symbol interface {
	Type
	Symbol() string
}

type String interface {
	Type
	String() string
}

type Error interface {
	Type
	Error() string
}

type List interface {
	Type
	List() string
	SubType() Type
}

type Array interface {
	Type
	SubType() Type
	Array() string
	Length()  int
}

type Function interface {
	Type 
	Function() string
	Arguments() []Type
	Returns() Type
}

//Type of type.
type Metatype interface {
	Type
	Metatype() string
}

type Dynamic interface {
	Type
	Dynamic() string
}

type FunctionType interface {
	Type
	FunctionType() string
	Arguments() []Type
	Returns() Type
}

type Custom interface {
	Type
	Custom() string
	
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
	Pointer() string
	
	SubType() Type
}

type Stream interface {
	Type
	
	Stream() string
}

type LanguageWithFloats interface {
	Language
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
type Language interface {

	// Hooks
	
		//Hook that is run before compilation.
		Init() Statement
		
		//These hooks are executed at the end of compilation:
		// Statements are added in the respective order.
		// Body is where all other output is placed.
		Head() Statement
		Neck() Statement
		Body() Statement
		Tail() Statement
		Last() Statement
		
	//Variables

		//Returns a statement that defines 'name' to be of type 'T' with optional 'value'.
 		Define(name string, T Type, value ...Type) Statement
 		
 		//Returns a Statement that sets the type 'T' variable 'name' to be set to 'value'.
 		Set(name string, T Type, value ...Type) Statement
 		
 		//Returns the Type at 'index' of 'T'.
 		Index(T Type, index Type) Type
 		
 		//Returns a statement that modifies type T at 'index' to be 'value'.
 		Modify(T Type, index Type, value Type) Statement

	// Logic
		
		//Returns a Statement that begins an if.
		If(condition Switch)
		
		//Returns a Statement that begins an elseif.
		ElseIf(condition Switch)
		
		//Returns a Statement that begins an else.
		Else()
		
		//Returns a Statement that ends an if/else.
		EndIf()
		
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
		
		//Returns a Statement that begins a for loop that iterates along the range between 'a' and 'b'.
		ForRange(a, b Number) Statement
		
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
		Return(T Type)
		
	//Threading
	
		//Returns a Stream connected to 'function' called with 'arguments' starting on another thread, netowork, coroutine or process.
		Thread(functiom Function, arguments []Type) Stream
		
	// Stdin, Stderr and Stdout
	
		//Returns a Statement that prints a Strings to os.Stdout with a newline.
		Print(...String) Statement
		
		//Returns a Statement that prints a String to os.Stdout without a newline.
		Write(...String) Statement
	
		//Reads Symbols from Stdin until Symbol is reached, returns a String of all Symbols up until Symbol.
		ReadSymbol(Symbol) String
		
		//Reads 'amount' bytes from Stdin, returns Array of all Bytes up until 'amount'. 
		ReadNumber(amount Number) Array

		//Returns a Statement that Reads bytes from Stdin and fills Array. 
		ReadArray(fill Array) Statement
		
	//Streams
	
		//Returns a Stream at 'path' associated with the given 'protocol'.
		OpenStream(protocol String, path String) Stream
		
		//Returns a String at 'path' associated with the given 'protocol'.
		LoadStream(protocol String, path String) String
		
		//Returns a statement that sends Type 't' over Stream 'c'.
		SendStream(c Stream, t Type) Statement
		
		//Returns Type 't' from Stream 'c'.
		ReadStream(c Stream, t Type) Type
		
		//Returns a statement that stops Stream 'c'.
		StopStream(c Stream) Statement
		
		//Returns a statement that seeks Stream 'c' by 'amount'.
		SeekStream(c Stream, amount Number) Statement
		
		//Returns a String that is the result of a 'query' on Stream 'c'.
		InfoStream(c Stream, query String) String
		
		//Returns a Statement that moves Stream 'c' to 'location'.
		MoveStream(c Stream, location String) Statement

	// Errors
		//Returns an Error with 'code' and 'message'.
		LiteralError(code Number, message String) Error
		
		//Returns a Statement that throws an error to the thread, this should not halt the program.
		ThrowError(err Error) Statement
		
		//Returns a Error that is sitting on the thread.
		CatchError() Error
		
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
		LiteralNumber(literal string) Number
		
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
		
		//Returns a statement that adds grows the 'list' by 'amount'.
		Grow(list List, amount Number)
		
		//Returns a statement that shrinks the 'list' by 'amount'.
		Shrink(list List, amount Number)
		
	//Function Types

		//Returns a FunctionType based on 'function'.
		LiteralFunctionType(function Function) FunctionType
		
		//Returns the resulting Type from calling 'function' with 'arguments'.
		CallFunctionType(function FunctionType, arguments []Type) Type
		
		//Returns a statement that calls the FunctionType 'function' with 'arguments'.
		RunFunctionType(function FunctionType, arguments []Type)
	
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
