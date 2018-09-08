package Go

import "github.com/qlova/script/language"


//Returns a Statement that prints a String to os.Stdout with a newline.
func (l *implementation) Print(values ...language.Type) language.Statement {
	l.Import("fmt")
	
	var PanicName = "Error in "+Name+".Print("
	for i := range values {
		PanicName += values[i].Name()
		if i < len(values)-1 {
			PanicName += ","
		}
	}
	PanicName += ")"
	
	var result = "fmt.Println("
	
	for i := range values {
		
		if i > 0 {
			result += " "
		}
		
		if values[i] == nil {
			result += `"nil"`
		} else {
		
			switch values[i].(type) {
				case language.Number, language.String, language.Boolean:
					result += l.GetExpression(values[i])
				
				case language.Symbol, 
					language.Custom, language.Stream, language.List, language.Array, 
					language.Table, language.Error, language.Float, language.Pointer, 
					language.Dynamic, language.Function, language.Metatype, language.FunctionType:
				
				panic(PanicName+": Unimplented")
					
				default:
					panic(PanicName+": Invalid Type")
			}

		}

		if i < len(values)-1 {
			result += ","
		}
	}
	
	result += ")\n"
	
	return language.Statement(result)
}

//Returns a Statement that writes a String to Stream (or Stdout) without a newline.
func (l *implementation) WriteString(language.Stream, language.String) language.Statement {
	panic("Error in "+Name+".WriteString(Stream, String): Unimplemented")
	return ""
}

//Returns a Statement that writes the contents of Array to a Stream (or Stdout) without a newline.
func (l *implementation) WriteArray(language.Stream, language.Array) language.Statement {
	panic("Error in "+Name+".WriteArray(Stream, String): Unimplemented")
	return ""
}

//Returns a statement that sends Type 't' over Stream 'c'.
func (l *implementation) Send(c language.Stream, t language.Type) language.Statement {
	panic("Error in "+Name+".Send(Stream, Type): Unimplemented")
	return ""
}

//Returns Type 't' from Stream 'c'.
func (l *implementation) Read(c language.Stream, t language.Type) language.Type {
	panic("Error in "+Name+".Read(Stream, Type): Unimplemented")
	return nil
}

//Reads Symbols from Stream (or Stdin) until Symbol is reached, returns a String of all Symbols up until Symbol.
func (l *implementation) ReadSymbol(language.Stream, language.Symbol) language.String {
	panic("Error in "+Name+".ReadSymbol(Stream, Symbol): Unimplemented")
	return nil
}

//Reads 'amount' bytes from Stream (or Stdin), returns Array of all Bytes up until 'amount'. 
func (l *implementation) ReadNumber(s language.Stream, amount language.Number) language.Array {
	panic("Error in "+Name+".ReadNumber(Stream, Number): Unimplemented")
	return nil
}

//Returns a Statement that Reads bytes from Stream (or Stdin) and fills Array. 
func (l *implementation) ReadArray(s language.Stream, fill language.Array) language.Statement {
	panic("Error in "+Name+".ReadArray(Stream, Array): Unimplemented")
	return ""
}
