package Interpreter

import "os"
import "os/exec"

import "github.com/qlova/script/interpreter/internal"
import "github.com/qlova/script/language"

//Returns a Stream connected to 'function' called with 'arguments' starting on another thread, netowork, coroutine or process.
func (l *implementation) Thread(functiom language.Function, arguments []language.Type) language.Stream {
	panic("Error in "+Name+".Thread(Function, []Type): Unimplemented")
	return nil
}

//Returns a Function that refers to external 'name'.
func (l *implementation) External(name language.String) (language.Function) {
	
	var f Function
	var p = name.(String)
	f.External = &p
	f.Rets = String{}

	return f
}

type Function struct {
	language.FunctionType
	internal.Variable
	Address internal.FunctionAddress
	
	IsLiteral bool
	
	External *String
}

func (l *implementation) CreateVariable(name string, T language.Type) language.Type {
	var block = l.loadBlock()
	
	switch T.(type) {
	
		case language.Function:
			var f Function
			f.BlockPointer = block
			f.Address = block.CreateFunction()
			return f
			
		case language.Number:
			var n Number
			n.BlockPointer = block
			n.Address = block.CreateNumber()
			return n
	}
	panic("Error in "+Name+".CreateVariable("+T.Name()+"): Not implemented!")
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

//Returns a Statement that begins the function 'name' with 'arguments' and 'returns'.
func (l *implementation) UpdateFunction(f language.Function, name string, names []string, arguments []language.Type, returns language.Type) (language.Function, language.Statement) {
	
	fn := f.(Function)
	fn.Rets = returns
	
	return fn, ""
}

//Returns a Statement that closes the last function.
func (l *implementation) EndFunction() language.Statement {
	OrginalBlock := l.popBlock()
	l.BlockPointer = OrginalBlock
	return ""
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
			
		case language.Number:
			
			var Address = A.(Number).Address
			var Block = A.(Number).BlockPointer
			var BlockPointer = B.(Number).BlockPointer
			var ValueAddress = B.(Number).Address
			
			block.AddInstruction(func() {
				Block.SetNumber(Address, BlockPointer.GetNumber(ValueAddress))
			})
		
		default:
			panic("Error in "+Name+".SetVariable("+A.Name()+", "+B.Name()+"): Not implemented!")

	}
}

func (l *implementation) ReturnVariable(T language.Type) {
	var block = l.loadBlock()
	var ReturnBlock = block.ReturnsBlock
	
	switch T.(type) {
	
		case language.Number:
			var BlockPointer = T.(Number).BlockPointer
			var ValueAddress = T.(Number).Address

			block.AddInstruction(func() {
				ReturnBlock.SetNumber(0, BlockPointer.GetNumber(ValueAddress))
			})
		
		default:
			panic("Error in "+Name+".ReturnVariable("+T.Name()+"): Not implemented!")

	}
}

func (l *implementation) GetReturnValue(function language.Function) language.Type {
	var block = l.loadBlock()
	
	switch function.Returns().(type) {
	
		case language.Number:
			var n Number
			if function.(Function).IsLiteral {
				n.BlockPointer = function.(Function).BlockPointer.ReturnsBlock
				n.Address = 0
				
			} else {
				n.Address = block.CreateNumber()
				n.BlockPointer = block
				
				var BlockPointer = function.(Function).BlockPointer
				var Address = function.(Function).Address
				block.AddInstruction(func() {
					block.SetNumber(n.Address, BlockPointer.GetFunction(Address).ReturnsBlock.GetNumber(0))
				})
			}
			block.AddInstruction(func(){})
			return n
		
		default:
			panic("Error in "+Name+".GetReturnValue("+function.Returns().Name()+"): Not implemented!")

	}
	
}

//Returns the resulting Type from calling 'function' with 'arguments'
func (l *implementation) Call(function language.Function, arguments []language.Type) language.Type {
	block := l.loadBlock()
	
	//External functions eg commands.
	if external := function.(Function).External; external != nil {
		
		//Create the return string of the external function.
		var ReturnString String
		ReturnString.Address = block.CreateString()
		ReturnString.BlockPointer = block
		
		var ReturnAddress = ReturnString.Address
		
		var BlockPointer = external.BlockPointer
		var Address = external.Address

		block.AddInstruction(func() {
			//Ouch? can we be smart and precache this?
			var args []string
			for _, argument := range arguments {
				var s = argument.(String)
				if s.IsLiteral {
					args = append(args, s.Literal)
				} else {
					args = append(args, s.BlockPointer.GetString(s.Address))
				}
			}
			
			var result, _ = exec.Command(BlockPointer.GetString(Address), args...).CombinedOutput()
			block.SetString(ReturnAddress, string(result))
		})
		
		return ReturnString
	}
	
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
		
		return l.GetReturnValue(function)
		
	}
		
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
	
	return l.GetReturnValue(function)
}
		
//Returns a Statement that calls 'function' with 'arguments'.
func (l *implementation) Run(function language.Function, arguments []language.Type) language.Statement {
	block := l.loadBlock()
	
	//External functions eg commands.
	if external := function.(Function).External; external != nil {
		
		var BlockPointer = external.BlockPointer
		var Address = external.Address

		block.AddInstruction(func() {
			//Ouch? can we be smart and precache this?
			var args []string
			for _, argument := range arguments {
				var s = argument.(String)
				if s.IsLiteral {
					args = append(args, s.Literal)
				} else {
					args = append(args, s.BlockPointer.GetString(s.Address))
				}
			}
			
			var cmd = exec.Command(BlockPointer.GetString(Address), args...)
			cmd.Stdout = os.Stdout
			cmd.Run()
		})
		
		return ""
	}
	
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
	block := l.loadBlock()
	if block.ReturnsBlock == nil {
		block.ReturnsBlock = new(internal.Block)
		l.BlockPointer = block.ReturnsBlock
		l.CreateVariable("return", T)
		l.BlockPointer = block
	}
	l.ReturnVariable(T)
	
	block.AddInstruction(func() {
		l.BlockPointer = l.popBlock()
		l.popPointer()
	})
	return ""
}
