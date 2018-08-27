package Null

import . "github.com/qlova/script/language"


//Returns a Statement that prints a Strings to os.Stdout with a newline.
func (l *language) Print(...String) Statement {	
	panic("Error in "+Name+".Print(...String): Unimplemented")
	return ""
}

//Returns a Statement that writes a String to Stream (or Stdout) without a newline.
func (l *language) WriteString(Stream, String) Statement {
	panic("Error in "+Name+".WriteString(Stream, String): Unimplemented")
	return ""
}

//Returns a Statement that writes the contents of Array to a Stream (or Stdout) without a newline.
func (l *language) WriteArray(Stream, Array) Statement {
	panic("Error in "+Name+".WriteArray(Stream, String): Unimplemented")
	return ""
}

//Returns a statement that sends Type 't' over Stream 'c'.
func (l *language) Send(c Stream, t Type) Statement {
	panic("Error in "+Name+".Send(Stream, Type): Unimplemented")
	return ""
}

//Returns Type 't' from Stream 'c'.
func (l *language) Read(c Stream, t Type) Type {
	panic("Error in "+Name+".Read(Stream, Type): Unimplemented")
	return nil
}

//Reads Symbols from Stream (or Stdin) until Symbol is reached, returns a String of all Symbols up until Symbol.
func (l *language) ReadSymbol(Stream, Symbol) String {
	panic("Error in "+Name+".ReadSymbol(Stream, Symbol): Unimplemented")
	return nil
}

//Reads 'amount' bytes from Stream (or Stdin), returns Array of all Bytes up until 'amount'. 
func (l *language) ReadNumber(s Stream, amount Number) Array {
	panic("Error in "+Name+".ReadNumber(Stream, Number): Unimplemented")
	return nil
}

//Returns a Statement that Reads bytes from Stream (or Stdin) and fills Array. 
func (l *language) ReadArray(s Stream, fill Array) Statement {
	panic("Error in "+Name+".ReadArray(Stream, Array): Unimplemented")
	return ""
}
