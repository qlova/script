package Javascript

import "github.com/qlova/script/language"

func (l *implementation) varaidic(fname string, values ...language.Type) language.Statement {
	
	var PanicName = "Error in "+Name+".varaidic("
	for i := range values {
		PanicName += values[i].Name()
		if i < len(values)-1 {
			PanicName += ","
		}
	}
	PanicName += ")"
	
	var result = fname
	
	for i := range values {
		
		if i > 0 {
			result += " "
		}
		
		if values[i] == nil {
			result += `"null"`
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
	
	result += ");"
	
	return language.Statement(result)
}

//Returns a Statement that prints a String to os.Stdout with a newline.
func (l *implementation) Print(values ...language.Type) language.Statement {
	//l.Import("fmt")
	return l.varaidic("console.log(", values...)
}

//Returns a Statement that writes a String to Stream (or Stdout) without a newline.
func (l *implementation) Write(stream language.Stream, values ...language.Type) language.Statement {
	var PanicName = "Error in "+Name+".Write(Stream, "
	for i := range values {
		PanicName += values[i].Name()
		if i < len(values)-1 {
			PanicName += ","
		}
	}
	PanicName += ")"
	
	if stream == nil {
		l.Import("fmt")
		return l.varaidic("fmt.Print(", values...)
	}
	
	panic(PanicName+": Unimplemented")
	return ""
}


//Returns a statement that sends Type 't' over Stream 'c'.
func (l *implementation) Send(c language.Stream, t language.Type) language.Statement {
	panic("Error in "+Name+".Send(Stream, Type): Unimplemented")
	return ""
}

//Returns Type 't' from Stream 'c'.
func (l *implementation) Read(c language.Stream, t language.Type) language.Type {
	if c != nil {
		panic("Error in "+Name+".Read(Stream, "+t.Name()+"): Unimplemented")
	}
	
	switch t.(type) {
		case Symbol:
			
			l.Import("bufio")
			l.Import("os")
			l.AddHelper(`var BStdin = bufio.NewReader(os.Stdin) 
	func ReadSymbol(symbol rune) string {
	result, _ := BStdin.ReadString(byte(symbol))
	return result
}

`)

			return String("ReadSymbol("+l.GetExpression(t)+")")
			
		default:
			panic("Error in "+Name+".Read(nil, "+t.Name()+"): Unimplemented")
	}
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
