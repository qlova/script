package Null

import "github.com/qlova/script/language"

const Name = "Null"

type implementation struct {}

func Language() *implementation {
	return new(implementation)
}

var _ = language.Interface(Language())

func (l *implementation) Init() {}
func (l *implementation) Head() language.Statement { return "" }
func (l *implementation) Neck() language.Statement { return "" }
func (l *implementation) Body() language.Statement { return "" }
func (l *implementation) Tail() language.Statement { return "" }
func (l *implementation) Last() language.Statement { return "" }

//Returns a Statement that begins the main entry point to the program.
func (l *implementation) Main() language.Statement {
	panic("Error in "+Name+".Main(): Unimplemented")
	return ""
}

//Returns a Statement that exits the program.
func (l *implementation) Exit() language.Statement {
	panic("Error in "+Name+".Exit(): Unimplemented")
	return ""
}

//Returns a Statement that ends the main entry point to the program.
func (l *implementation) EndMain() language.Statement {
	panic("Error in "+Name+".EndMain(): Unimplemented")
	return ""
}

//Returns a statement that defines 'name' to be of type 'T' with optional 'value'.
func (l *implementation) Define(name string, T language.Type, value ...language.Type) language.Statement {
	var PanicName = "Error in "+Name+".Define("+name+", "+T.Name()
	if len(value) == 0 {
		PanicName += ", "+value[0].Name()+")"
	} else {
		PanicName += ")"
	}
	
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

//Returns a Statement that sets the type 'T' variable 'name' to be set to 'value'.
func (l *implementation) Set(name string, T language.Type, value ...language.Type) language.Statement {
	var PanicName = "Error in "+Name+".Set("+name+", "+T.Name()
	if len(value) == 0 {
		PanicName += ", "+value[0].Name()+")"
	} else {
		PanicName += ")"
	}
	
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
