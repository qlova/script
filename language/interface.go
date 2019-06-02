package language

type Interface interface {
	Real(r float64) Real
	Rational() Rational
	Natural(n uint) Natural
	Octonion() Octonion
	Integer(i int) Integer
	Duplex() Duplex
	Complex() Complex
	Quaternion() Quaternion
	Sedenion() Sedenion
	Color() Color
	Sound() Sound
	Stream() Stream
	Image() Image
	Video() Video
	Time() Time
	Symbol(r rune) Symbol
	String(s string) String
	Bit(b bool) Bit
	Byte(b byte) Byte
	Name() string
	Init()
	Build(path string) func()
	//Hooks
	Head() Statement
	Neck() Statement
	Body() Statement
	Tail() Statement
	Last() Statement
	//Context
	Buffer() Buffer
	Flush(buffer Buffer)
	//Variables
	Register(register string, value Type) (Statement, Type)
	Set(variable, value Type) Statement
	//Structures
	Index(structure, index Type) Type
	Modify(structure, index, value Type) Statement
	//Branching
	If(condition Bit) Statement
	ElseIf(condition Bit) Statement
	Else() Statement
	EndIf() Statement
	//Logic
	And(a, b Bit) Bit
	Or(a, b Bit) Bit
	Not(b Bit) Bit
	Equals(a, b Type) Bit
	Smaller(a, b Type) Bit
	Greater(a, b Type) Bit
	//Loops
	Loop() Statement
	EndLoop() Statement
	Break() Statement
	While(condition Bit) Statement
	EndWhile() Statement
	ForRange(i string, a, b Number) (Statement, Type)
	EndForRange() Statement
	ForEach(i, v string, list Type) (Statement, Type, Type)
	EndForEach() Statement
	For(i string, condition Bit, action Statement) (Statement, Type)
	EndFor() Statement
	//Entrypoint
	Main() Statement
	EndMain() Statement
	Exit() Statement
	//Function
	Function(name string, registers []string, arguments []Type, returns Type) (Statement, Function)
	EndFunction() Statement
	Call(f Function, arguments []Type) Type
	Run(f Function, arguments []Type) Statement
	Return(value Type) Statement
	//Threading
	Thread(name string, distance int, arguments []Type) Stream
	//IO
	Print(values ...Type) Statement
	Write(stream Stream, values ...Type) Statement
	//Streams
	Open(protocol string, path String) Type
	Load(protocol string, path String) Type
	Read(stream Stream, mode Type) Type
	Stop(stream Stream) Statement
	Seek(stream Stream, amount Integer) Statement
	Info(stream Stream, query String) String
	Move(stream Stream, location String) Statement
	//Errors
	Error() Error
	Trace(line int, file string) Statement
	Throw(code Number, message String) Statement
	Catch() Bit
	//Operations
	Add(a, b Number) Number
	Sub(a, b Number) Number
	Mul(a, b Number) Number
	Div(a, b Number) Number
	Pow(a, b Number) Number
	Mod(a, b Number) Number
	//Collections
	ArrayOf(t Type, length int) Array
	Join(a, b Type) Type
	Length(t Type) Integer
	Push(value Type, list Type) Statement
	Pop(list Type) Type
	TableOf(t Type) Table
	ListOf(t Type) List
	List(t ...Type) List
	Array(t ...Type) Array
	Table(index String, value Type) Table
	//Pointers
	PointerOf(t Type) Pointer
	Dereference(p Pointer) Type
	//Reflection
	DynamicOf(t Type) Dynamic
	StaticOf(d Dynamic) Type
	MetatypeOf(t Type) Metatype
	//Casts
	Cast(a, b Type) Type
	//Custom
	Type(name string, registers []string, elements []Type) Statement
	Method(t string, name string, registers []string, arguments []Type, returns Type) Statement
	This() Type
	New(name string) Type
	Invoke(t Type, method string, arguments []Type) Type
	Execute(t Type, method string, arguments []Type) Statement
	EndMethod() Statement
	//Util
	Copy(t Type) Type
}
