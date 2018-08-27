package Null

import . "github.com/qlova/script/language"

const Name = "Null"

type language struct {}

func Language() *language {
	return new(language)
}

var _ = Interface(Language())

func (l *language) Init() {}
func (l *language) Head() Statement { return "" }
func (l *language) Neck() Statement { return "" }
func (l *language) Body() Statement { return "" }
func (l *language) Tail() Statement { return "" }
func (l *language) Last() Statement { return "" }

//Returns a Statement that begins the main entry point to the program.
func (l *language) Main() Statement {
	panic("Error in "+Name+".Main(): Unimplemented")
	return ""
}

//Returns a Statement that exits the program.
func (l *language) Exit() Statement {
	panic("Error in "+Name+".Exit(): Unimplemented")
	return ""
}

//Returns a Statement that ends the main entry point to the program.
func (l *language) EndMain() Statement {
	panic("Error in "+Name+".EndMain(): Unimplemented")
	return ""
}

//Returns a statement that defines 'name' to be of type 'T' with optional 'value'.
func (l *language) Define(name string, T Type, value ...Type) Statement {
	var PanicName = "Error in "+Name+".Define("+name+", "+T.Name()
	if len(value) == 0 {
		PanicName += ", "+value[0].Name()+")"
	} else {
		PanicName += ")"
	}
	
	switch T.(type) {
		case Switch, Number, Symbol, String, 
			Custom, Stream, List, Array, 
			Table, Error, Float, Pointer, 
			Dynamic, Function, Metatype, FunctionType:
		
		panic(PanicName+": Unimplented")
			
		default:
			panic(PanicName+": Invalid Type")
	}
	
	return ""
}

//Returns a Statement that sets the type 'T' variable 'name' to be set to 'value'.
func (l *language) Set(name string, T Type, value ...Type) Statement {
	var PanicName = "Error in "+Name+".Set("+name+", "+T.Name()
	if len(value) == 0 {
		PanicName += ", "+value[0].Name()+")"
	} else {
		PanicName += ")"
	}
	
	switch T.(type) {
		case Switch, Number, Symbol, String, 
			Custom, Stream, List, Array, 
			Table, Error, Float, Pointer, 
			Dynamic, Function, Metatype, FunctionType:
		
		panic(PanicName+": Unimplented")
			
		default:
			panic(PanicName+": Invalid Type")
	}
	
	return ""
}

//Returns the Type at 'index' of 'T'.
func (l *language) Index(T Type, index Type) Type {
	var PanicName = "Error in "+Name+".Index("+T.Name()+", "+index.Name()+")"
	
	switch T.(type) {
		case Switch, Number, Symbol, String, 
			Custom, Stream, List, Array, 
			Table, Error, Float, Pointer, 
			Dynamic, Function, Metatype, FunctionType:
		
		panic(PanicName+": Unimplented")
			
		default:
			panic(PanicName+": Invalid Type")
	}
	
	return nil
}

//Returns a statement that modifies type T at 'index' to be 'value'.
func (l *language) Modify(T Type, index Type, value Type) Statement {
	var PanicName = "Error in "+Name+".Modify("+T.Name()+", "+index.Name()
	PanicName += ", "+value.Name()+")"
	
	switch T.(type) {
		case Switch, Number, Symbol, String, 
			Custom, Stream, List, Array, 
			Table, Error, Float, Pointer, 
			Dynamic, Function, Metatype, FunctionType:
		
		panic(PanicName+": Unimplented")
			
		default:
			panic(PanicName+": Invalid Type")
	}
	
	return ""
}
