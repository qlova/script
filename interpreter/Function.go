package interpreter

import "github.com/qlova/script/interpreter/dynamic"
import "github.com/qlova/script/language"

func (implementation Implementation) Function(name string, registers []string, arguments []language.Type, returns language.Type) (language.Statement, language.Function) {
	
	if name != "" && len(arguments) == 0 {
		
		var block = implementation.program.CreateBlock()
		//println("creating function ", name, " as ", block)

		implementation.Activate(block)
		
		return "", Function{
			Expression: language.Statement(block.String()),
		}
	}
	
	panic(implementation.Name()+".Function() Unimplemented")
	return language.Statement(""), nil
}

func (implementation Implementation) EndFunction() language.Statement {
	implementation.Deactivate()

	return language.Statement("")
}

func (implementation Implementation) Call(f language.Function, arguments []language.Type) language.Type {
	panic(implementation.Name()+".Call() Unimplemented")
	return nil
}

func (implementation Implementation) Run(f language.Function, arguments []language.Type) language.Statement {
	if len(arguments) == 0 {
		var block = implementation.BlockOf(f)

		implementation.AddInstruction(func(thread *dynamic.Thread) {
			thread.JumpTo(block)
		})
		
		return ""
	}
	
	panic(implementation.Name()+".Run() Unimplemented")
	return language.Statement("")
}

func (implementation Implementation) Return(value language.Type) language.Statement {
	panic(implementation.Name()+".Return() Unimplemented")
	return language.Statement("")
}

