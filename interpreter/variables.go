package Interpreter

import "github.com/qlova/script/language"

//Returns a statement that defines 'name' to be of type 'T' with optional 'value'.
func (l *implementation) Define(name string, value language.Type) (language.Type, language.Statement) {
	var PanicName = "Error in "+Name+".Define("+name+", "+value.Name()+")"
	
	block := l.loadBlock()
	
	var Address = block.CreateNumber()
	
	switch value.(type) {
		case language.Number:
			number := value.(Number)
			
			if number.Literal != nil {
				literal := *number.Literal
				block.AddInstruction(func() {
					block.SetNumber(Address, &literal)
				})
			} else {
				var BlockPointer = number.BlockPointer
				var ValueAddress = number.Address
				block.AddInstruction(func() {
					block.SetNumber(Address, BlockPointer.GetNumber(ValueAddress))
				})
			}
			
			return number, ""
		
		case language.Switch, language.Symbol, language.String, 
			language.Custom, language.Stream, language.List, language.Array, 
			language.Table, language.Error, language.Float, language.Pointer, 
			language.Dynamic, language.Function, language.Metatype, language.FunctionType:
		
		panic(PanicName+": Unimplented")
			
		default:
			panic(PanicName+": Invalid Type")
	}
	
	return nil, ""
}

//Returns a Statement that sets the type 'T' variable 'name' to be set to 'value'.
func (l *implementation) Set(name string, T language.Type, value language.Type) language.Statement {
	var PanicName = "Error in "+Name+".Set("+name+", "+T.Name()+", "+value.Name()+")"
	
	switch T.(type) {
		case language.Switch, language.Number, language.Symbol, language.String, 
			language.Custom, language.Stream, language.List, language.Array, 
			language.Table, language.Error, language.Float, language.Pointer, 
			language.Dynamic, language.Function, language.Metatype, language.FunctionType:
		
		panic(PanicName+": Unimplented")
			
		default:
			panic(PanicName+": Invalid Type")
	}
	
	return ""
}

//Returns the Type at 'index' of 'T'.
func (l *implementation) Index(T language.Type, index language.Type) language.Type {
	var PanicName = "Error in "+Name+".Index("+T.Name()+", "+index.Name()+")"
	
	switch T.(type) {
		case language.Switch, language.Number, language.Symbol, language.String, 
			language.Custom, language.Stream, language.List, language.Array, 
			language.Table, language.Error, language.Float, language.Pointer, 
			language.Dynamic, language.Function, language.Metatype, language.FunctionType:
		
		panic(PanicName+": Unimplented")
			
		default:
			panic(PanicName+": Invalid Type")
	}
	
	return nil
}

//Returns a statement that modifies type T at 'index' to be 'value'.
func (l *implementation) Modify(T language.Type, index language.Type, value language.Type) language.Statement {
	var PanicName = "Error in "+Name+".Modify("+T.Name()+", "+index.Name()
	PanicName += ", "+value.Name()+")"
	
	switch T.(type) {
		case language.Switch, language.Number, language.Symbol, language.String, 
			language.Custom, language.Stream, language.List, language.Array, 
			language.Table, language.Error, language.Float, language.Pointer, 
			language.Dynamic, language.Function, language.Metatype, language.FunctionType:
		
		panic(PanicName+": Unimplented")
			
		default:
			panic(PanicName+": Invalid Type")
	}
	
	return ""
}
 
