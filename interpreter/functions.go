package Interpreter

import "github.com/qlova/script/interpreter/internal"
import "github.com/qlova/script/language"

//Returns a Stream connected to 'function' called with 'arguments' starting on another thread, netowork, coroutine or process.
func (l *implementation) Thread(functiom language.Function, arguments []language.Type) language.Stream {
	panic("Error in "+Name+".Thread(Function, []Type): Unimplemented")
	return nil
}

type Function struct {
	language.FunctionType
	internal.Variable
	Address internal.FunctionAddress
	
	IsLiteral bool
}

func (l *implementation) CreateVariable(name string, T language.Type) language.Type {
	var block = l.loadBlock()
	
	switch T.(type) {
	
		case language.Function:
			var f Function
			f.BlockPointer = block
			f.Address = block.CreateFunction()
			return f
			

	}
	panic("Error in "+Name+".CreateVariable("+name+"): Not implemented!")
}

//Returns a Statement that begins the function 'name' with 'arguments' and 'returns'.
func (l *implementation) Function(name string, names []string, arguments []language.Type, returns language.Type) (language.Function, language.Statement) {
	
	Block := l.loadBlock()
	l.pushBlock(Block)
	
	//This is tricky.
	NextBlock := new(internal.Block)
	l.BlockPointer = NextBlock
	
	var f Function
	f.IsLiteral = true
	f.BlockPointer = NextBlock
	
	if len(arguments) == 0 {
		f.Args = nil 
	} else {
		f.Args = make([]language.Type, len(arguments))
		for i := range arguments {
			f.Args[i] = l.CreateVariable(names[i], arguments[i])
		}
	}
	f.Rets = returns
	
	return f, ""
}

//Returns a Statement that closes the last function.
func (l *implementation) EndFunction() language.Statement {
	OrginalBlock := l.popBlock()
	l.BlockPointer = OrginalBlock
	return ""
}
		
//Returns the resulting Type from calling 'function' with 'arguments'
func (l *implementation) Call(function language.Function, arguments []language.Type) language.Type {
	panic("Error in "+Name+".Call(Function, []Type): Unimplemented")
	return nil
}

func (l *implementation) SetVariable(A language.Type, B language.Type) {
	var block = l.loadBlock()
	
	switch A.(type) {
	
		case language.Function:
			
			var Address = A.(Function).Address
			var Block = A.(Function).BlockPointer
			var BlockPointer = B.(Function).BlockPointer
			var ValueAddress = B.(Function).Address
			
			block.AddInstruction(func() {
				Block.SetFunction(Address, BlockPointer.GetFunction(ValueAddress))
			})
		
		default:
			panic("Error in "+Name+".SetVariable("+A.Name()+", "+B.Name()+"): Not implemented!")

	}
	
}
		
//Returns a Statement that calls 'function' with 'arguments'.
func (l *implementation) Run(function language.Function, arguments []language.Type) language.Statement {
	block := l.loadBlock()
	
	if function.(Function).IsLiteral {
		
		if arguments != nil {
			var Args = function.(Function).Arguments()
			for i := range arguments {
				l.SetVariable(Args[i], arguments[i])
			}
		}
		
		var BlockPointer = function.(Function).BlockPointer
		block.AddInstruction(func() {
			l.JumpToBlock(BlockPointer)
		})
		
		
	} else {
		
		if arguments != nil {
			var Args = function.(Function).Arguments()
			for i := range arguments {
				l.SetVariable(Args[i], arguments[i])
			}
		}
		
		var BlockPointer = function.(Function).BlockPointer
		var Address = function.(Function).Address
		block.AddInstruction(func() {
			l.JumpToBlock(BlockPointer.GetFunction(Address))
		})
	}
	return ""
}

//Returns a Statement that returns T from the current function.
func (l *implementation) Return(T language.Type) language.Statement {
	panic("Error in "+Name+".Function(Type): Unimplemented")
	return ""
}
