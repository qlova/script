package Interpreter

import "github.com/qlova/script/language"
import "os"
import "bufio"

func (c *implementation) PrintString(text String) {
	block := c.loadBlock()
	
	if text.IsLiteral {

		str := text.Literal
		
		block.AddInstruction(func() {
			os.Stdout.Write([]byte(str))
		})

	} else {

		var BlockPointer = text.BlockPointer
		var Address = text.Address
		block.AddInstruction(func() {
			os.Stdout.Write([]byte(BlockPointer.GetString(Address)))
		})

	}
}

func (c *implementation) PrintNumber(number Number) {
	block := c.loadBlock()
	
	if number.Literal != nil {

		literal := *number.Literal
		block.AddInstruction(func() {
			print(literal.String())
		})

	} else {

		var BlockPointer = number.BlockPointer
		var Address = number.Address
		block.AddInstruction(func() {
			print(BlockPointer.GetNumber(Address).String())
		})

	}
}


func (c *implementation) PrintBoolean(boolean Boolean) {
	block := c.loadBlock()
	
	if boolean.Literal != nil {

		literal := *boolean.Literal
		block.AddInstruction(func() {
			print(literal)
		})

	} else {

		var BlockPointer = boolean.BlockPointer
		var Address = boolean.Address
		block.AddInstruction(func() {
			print(BlockPointer.GetBoolean(Address))
		})

	}
}



//Returns a Statement that prints a Strings to os.Stdout with a newline.
func (c *implementation) Print(values ...language.Type) language.Statement {
	
	c.Write(nil, values...)
	
	block := c.loadBlock()
	block.AddInstruction(func() {
		os.Stdout.Write([]byte("\n"))
	})
	
	return ""
}

//Returns a Statement that writes a String to Stream (or Stdout) without a newline.
func (c *implementation) Write(stream language.Stream, values ...language.Type) language.Statement {
	var PanicName = "Error in "+Name+".Write("
	for i := range values {
		PanicName += values[i].Name()
		if i < len(values)-1 {
			PanicName += ","
		}
	}
	PanicName += ")"
	
	if stream != nil {
		panic(PanicName+": Cannot use streams yet")
	}
	
	
	block := c.loadBlock()
	
	if len(values) == 0 {
		block.AddInstruction(func() {
			os.Stdout.Write([]byte("\n"))
		})
		return ""
	}
	
	for i := range values {
		if values[i] == nil {
			block.AddInstruction(func() {
				os.Stdout.Write([]byte("nil"))
			})
			continue
		}
		
		switch values[i].(type) {
			case language.String:
				c.PrintString(values[i].(String))
				
			case language.Number:
				c.PrintNumber(values[i].(Number))
			
			case language.Boolean:
				c.PrintBoolean(values[i].(Boolean))
				
			case language.Symbol, 
				language.Custom, language.Stream, language.List, language.Array, 
				language.Table, language.Error, language.Float, language.Pointer, 
				language.Dynamic, language.Function, language.Metatype, language.FunctionType:
			
			panic(PanicName+": Unimplented")
				
			default:
				panic(PanicName+": Invalid Type")
		}
		
		if i < len(values)-1 {
			block.AddInstruction(func() {
				os.Stdout.Write([]byte(" "))
			})
		}
	}
	return ""
}

//Returns a statement that sends Type 't' over Stream 'c'.
func (l *implementation) Send(c language.Stream, t language.Type) language.Statement {
	panic("Error in "+Name+".Send(Stream, Type): Unimplemented")
	return ""
}

var BufferedStdin = bufio.NewReader(os.Stdin) 

//Returns Type 't' from Stream 'c'.
func (l *implementation) Read(c language.Stream, t language.Type) language.Type {
	if c != nil { 
		panic("Error in "+Name+".Read(Stream, "+t.Name()+"): Unimplemented")
	}
	block := l.loadBlock()
	
	switch value := t.(type) {
		case Symbol:
			
			var s = l.NewString()
			
			var Address = s.Address
			
			var SymbolBlock = value.BlockPointer
			var SymbolAddress = value.Address
			
			if value.IsLiteral {
				
				var literal = value.Literal
				//TODO deal with errors,
				block.AddInstruction(func() {
					result, _ := BufferedStdin.ReadString(byte(literal))
					block.SetString(Address, result)
				})
				
			} else {
				//TODO deal with errors,
				block.AddInstruction(func() {
					result, _ := BufferedStdin.ReadString(byte(SymbolBlock.GetSymbol(SymbolAddress)))
					block.SetString(Address, result)
				})
			}

			return s
			
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
