package Null

import . "github.com/qlova/script/language"


//Returns a Stream at 'path' associated with the given 'protocol'.
func (l *language) Open(protocol String, path String) Stream {
	panic("Error in "+Name+".Read(String, String): Unimplemented")
	return nil
}

//Returns a String at 'path' associated with the given 'protocol'.
func (l *language) Load(protocol String, path String) String {
	panic("Error in "+Name+".Load(String, String): Unimplemented")
	return nil
}

//Returns a statement that stops Stream 'c'.
func (l *language) Stop(c Stream) Statement {
	panic("Error in "+Name+".Stop(Stream): Unimplemented")
	return ""
}

//Returns a statement that seeks Stream 'c' by 'amount'.
func (l *language) Seek(c Stream, amount Number) Statement {
	panic("Error in "+Name+".Seek(Stream, Number): Unimplemented")
	return ""
}

//Returns a String that is the result of a 'query' on Stream 'c'.
func (l *language) Info(c Stream, query String) String {
	panic("Error in "+Name+".Info(Stream, String): Unimplemented")
	return nil
}

//Returns a Statement that moves Stream 'c' to 'location'.
func (l *language) Move(c Stream, location String) Statement {
	panic("Error in "+Name+".Move(Stream, String): Unimplemented")
	return ""
}
