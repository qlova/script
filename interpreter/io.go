package Interpreter

import "github.com/qlova/script/language"
import "os"

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
	
	var PanicName = "Error in "+Name+".Print("
	for i := range values {
		PanicName += values[i].Name()
		if i < len(values)-1 {
			PanicName += ","
		}
	}
	PanicName += ")"
	
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
	
	block.AddInstruction(func() {
		os.Stdout.Write([]byte("\n"))
	})
	
	return ""
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
