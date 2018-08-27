package Go

import "github.com/qlova/script/language"

const Name = "Go"

type implementation struct {
	imports []string
}

func (l *implementation) Import(pkg string) {
	for i := range l.imports {
		if l.imports[i] == pkg {
			return
		}
	}
	l.imports = append(l.imports, pkg)
}

func Language() *implementation {
	return new(implementation)
}


//TODO remove.
func init() {
	language.Default = language.Interface(Language())
}


func (l *implementation) Init() {}
func (l *implementation) Head() language.Statement {
	var result = "package main\n"

	for _, pkg := range l.imports {
		result += `import "`+pkg+`"`
	}

	return language.Statement(result)
}
func (l *implementation) Neck() language.Statement { return "" }
func (l *implementation) Body() language.Statement { return "" }
func (l *implementation) Tail() language.Statement { return "" }
func (l *implementation) Last() language.Statement { return "" }

//Returns a Statement that begins the main entry point to the program.
func (l *implementation) Main() language.Statement {
	return "func main() {\n"
}

//Returns a Statement that exits the program.
func (l *implementation) Exit() language.Statement {
	l.Import("os")
	l.imports = append(l.imports, "os")
	
	return "os.Exit(0)"
}

//Returns a Statement that ends the main entry point to the program.
func (l *implementation) EndMain() language.Statement {
	return "}"
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
