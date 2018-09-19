package language

import "os/exec"

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

type Boolean interface {
	Type
	Boolean()
}

type Symbol interface {
	Type
	Symbol()
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

//Other type ideas.
/*
 * Decimal
 * Complex
 * Rational
 * Matrix
 * Duplex
 * Vector
 * 
 * Far fetched
 *========
 * Image?
 * Color?
 * Video?
 * Audio?
 */

// Defines a Language, all methods can panic on error. Make sure to deal with this accordingly.
type Interface interface {
	Name() string

	// Hooks
	
		//Hook that is run before compilation.
		Init()
		
		Build(path string) *exec.Cmd
		
		//These hooks are executed at the end of compilation:
		// Statements are added in the respective order.
		// Body is where all other output is placed.
		Head() Statement
		Neck() Statement
		Body() Statement
		Tail() Statement
		Last() Statement
		
	//Variables

		//Returns an Array of type T with 'length' length.
		Literal(interface{}) Type

		//Returns the type and statement that defines 'name' to be of type 'value' initialised to 'value'.
 		Define(name string, value Type) (Type, Statement)
 		
 		//Returns a Statement that sets the variable 'T' to be set to 'value'.
 		Set(T Type, value Type) Statement
 		
 		//Returns the Type at 'index' of 'T'.
 		Index(T Type, index Type) Type
 		
 		//Returns a statement that modifies type T at 'index' to be 'value'.
 		Modify(T Type, index Type, value Type) Statement

	// Logic
		
		//Returns a Statement that begins an if.
		If(condition Boolean) Statement
		
		//Returns a Statement that begins an elseif.
		ElseIf(condition Boolean) Statement
		
		//Returns a Statement that begins an else.
		Else() Statement
		
		//Returns a Statement that ends an if/else.
		EndIf() Statement
		
		Equals(Type, Type) Boolean
		Smaller(Type, Type) Boolean
		Greater(Type, Type) Boolean
		
	// Loops
		
		//Returns a Statement that begins an infinite loop.
		Loop() Statement
		
		//Returns a Statement that ends an infinite loop, do not confuse this with Break().
		EndLoop() Statement
		
		//Returns a Statement that breaks an infinite loop.
		Break() Statement
		
		//Returns a Statement that begins a while loop dependent on 'condition' being non zero.
		While(condition Boolean) Statement
		
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
	
		//Returns an external function, eg OS command or dynamic library etc
		External(name String) Function
		
		//Returns a Statement that begins the function 'name' with 'arguments' and 'returns'.
		Function(name string, names []string, arguments []Type, returns Type) (Function, Statement)
		
		//A hacky thing requried for the compiler so it can do concepts.
		UpdateFunction(f Function, name string, names []string, arguments []Type, returns Type) (Function, Statement)
		
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
		
		//Returns a Statement that writes Type to Stream (or Stdout if nil) without a newline.
		Write(Stream, ...Type) Statement

		//Returns a statement that sends Type 't' over Stream 'c'.
		Send(c Stream, t Type) Statement
		
		//Returns Type 't' from Stream 'c'.
		Read(c Stream, t Type) Type
		
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
	
		//Returns the current error.
		Error() Error
		
		//Set the current trace, this will be attached to all future errors.
		Trace(line int, file string) Statement
		
		//Returns a Statement that throws an error to the thread with the given parameters, this should not halt the program.
		Throw(code Number, message String) Statement
		
		//Checks if there is an error on the stack, makes it the current error and returns whether or not this was succesfull.
		Catch() Boolean
		
	// Booleanes aka Booleans
		
		//Returns a Boolean that is the logical and of 'a' and 'b'.
		And(a, b Boolean) Boolean
		
		//Returns a Boolean that is the logical or of 'a' and 'b'.
		Or(a, b Boolean) Boolean
		
		//Returns a Boolean that is the logical not of 'a'.
		Not(a Boolean) Boolean

	// Numbers
		
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

	//Strings
		
		//Returns a new Type that concatenates 'a' and 'b'.
		Join(a, b Type) Type

	//Arrays, fixed-size collection of elements.
	
		//Returns the array type of type 'T' with 'length'.
		Array(T Type, length int) Array
		
		//Fills type 'T' with 'elements'
		Fill(T Type, elements []Type) Type

	//Lists, unordered collection of elements.

		//Returns a Number representing the length of 'list'
		Length(t Type) Number
		
		//Returns a statement that adds 't' to the 'list'.
		Grow(list List, t Type) Statement
		
		//Returns a statement that shrinks the 'list' by removing the element at the end..
		Shrink(list List) Statement
	
	// Tables.
		
		//Returns a Table of type T.
		MakeTable(T Type) Table
		
		//Returns a Table intialised with 'indices' corresponding to 'elements'.
		//LiteralTable(indices []String, elements []Type) Table
		
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
		//LiteralCustom(tokens []string, elements []Type) Custom
		
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
		
		//Returns Type cast to String.
		ToString(Type) String
		
		//Returns Type cast to Number.
		ToNumber(Type) Number
		
		//Returns Type cast to Number.
		ToSymbol(Type) Symbol
}
