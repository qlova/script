package Null

import "github.com/qlova/script/language"


//Returns a Stream at 'path' associated with the given 'protocol'.
func (l *implementation) Open(protocol, path language.String) language.Stream {
	panic("Error in "+Name+".Read(String, String): Unimplemented")
	return nil
}

//Returns a String at 'path' associated with the given 'protocol'.
func (l *implementation) Load(protocol, path language.String) language.String {
	panic("Error in "+Name+".Load(String, String): Unimplemented")
	return nil
}

//Returns a statement that stops Stream 'c'.
func (l *implementation) Stop(c language.Stream) language.Statement {
	panic("Error in "+Name+".Stop(Stream): Unimplemented")
	return ""
}

//Returns a statement that seeks Stream 'c' by 'amount'.
func (l *implementation) Seek(c language.Stream, amount language.Number) language.Statement {
	panic("Error in "+Name+".Seek(Stream, Number): Unimplemented")
	return ""
}

//Returns a String that is the result of a 'query' on Stream 'c'.
func (l *implementation) Info(c language.Stream, query language.String) language.String {
	panic("Error in "+Name+".Info(Stream, String): Unimplemented")
	return nil
}

//Returns a Statement that moves Stream 'c' to 'location'.
func (l *implementation) Move(c language.Stream, location language.String) language.Statement {
	panic("Error in "+Name+".Move(Stream, String): Unimplemented")
	return ""
}
