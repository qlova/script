package language

type Interface interface {
	Complex() Complex
	Sedenion() Sedenion
	Real(r float64) Real
	Rational() Rational
	Natural(n uint) Natural
	Integer(i int) Integer
	Duplex() Duplex
	Quaternion() Quaternion
	Octonion() Octonion
	Sound() Sound
	Stream() Stream
	String(s string) String
	Bit(b bool) Bit
	Color() Color
	Video() Video
	Time() Time
	Symbol(r rune) Symbol
	Byte(b byte) Byte
	Image() Image
	Name() string
	Init()
	Build(path string) func()
	//Hooks
	Head() Statement
	Neck() Statement
	Body() Statement
	Tail() Statement
	Last() Statement
	//Variables
	Register(register string, value Type) (Statement, Type)
	Set(register string, value Type) Statement
	Get(register string, value Type) Type
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
	ForRange(i string, a, b Number) Statement
	EndForRange() Statement
	ForEach(i, v string, list Type) Statement
	EndForEach() Statement
	For(i string, condition Bit, action Statement) Statement
	EndFor() Statement
	//Entrypoint
	Main() Statement
	EndMain() Statement
	Exit() Statement
	//Function
	Function(name string, registers []string, arguments []Type, returns Type) Statement
	EndFunction() Statement
	Call(name string, arguments []Type) Type
	Run(name string, arguments []Type) Statement
	Return(value Type) Statement
	//Threading
	Thread(name string, distance int, arguments []Type) Stream
	//IO
	Print(values ...Type) Statement
	Write(stream Stream, values ...Type) Statement
	//Streams
	Open(protocol string, path String) Type
	Load(protocol string, path String) Type
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
}
