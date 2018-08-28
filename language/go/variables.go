package Go

import "github.com/qlova/script/language"

func GetVariable(name string, T language.Type) language.Type {
	var PanicName = "Error in "+Name+".GetVariable("+T.Name()+")"
	
	switch T.(type) {
		case language.Number:
			return Number{Expression: name}
			
		case language.String:
			return String(name)
		
		case language.Switch, language.Symbol, 
			language.Custom, language.Stream, language.List, language.Array, 
			language.Table, language.Error, language.Float, language.Pointer, 
			language.Dynamic, language.Function, language.Metatype, language.FunctionType:
		
		panic(PanicName+": Unimplented")
			
		default:
			panic(PanicName+": Invalid Type")
	}
}

//Returns a statement that defines 'name' to be of type 'value', initialised to 'value'.
func (l *implementation) Define(name string, value language.Type) (language.Type, language.Statement) {

	var statement = "var "+name+" = "+l.GetExpression(value)+"\n"
	var variable = GetVariable(name, value)

	return variable, language.Statement(statement)
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
