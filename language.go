package script

import "io"

type Result *func() interface{}

//Language is a script language implementation.
type Language interface {
	io.Writer
	Raw(Value) string

	Main(func())
	Print(Values)
	If(If, []If, func())

	//Types
	DefineStruct(Struct)

	//Structures
	Field(Value, string) Result
	//MutateField(Value, string, Value)

	//Methods
	RunMethod(Value, string, []Value)
	CallMethod(Value, string, []Value) Result

	//Variables
	DefineVariable(string, Value) Result
	Set(Value, Value)

	//Lists
	Index(Value, Int) Result
	Mutate(Value, Int, Value)

	//Tables
	Lookup(Value, String) Result
	Insert(Value, String, Value)

	//Loops.
	For(set Int, condition ForLoopCondition, action ForLoopAction, f func(Int))
	While(Bool, func())
	Loop(func())
	Break()

	//Functions
	DefineFunction(Function)
	RunFunction(name string, args []Value)
	CallFunction(name string, args []Value) Result
	Argument(name string, index int) Result
	Return(Value)

	Plus(Int, Int) Int
	Same(Int, Int) Bool
	Not(Bool) Bool

	Join(String, String) String
}
