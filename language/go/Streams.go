package Go

import "github.com/qlova/script/language"

func (implementation Implementation) Open(protocol string, path language.String) language.Type {
	panic(implementation.Name()+".Open() Unimplemented")
	return nil
}

func (implementation Implementation) Load(protocol string, path language.String) language.Type {
	panic(implementation.Name()+".Load() Unimplemented")
	return nil
}

func (implementation Implementation) Stop(stream language.Stream) language.Statement {
	panic(implementation.Name()+".Stop() Unimplemented")
	return language.Statement("")
}

func (implementation Implementation) Seek(stream language.Stream, amount language.Integer) language.Statement {
	panic(implementation.Name()+".Seek() Unimplemented")
	return language.Statement("")
}

func (implementation Implementation) Info(stream language.Stream, query language.String) language.String {
	panic(implementation.Name()+".Info() Unimplemented")
	return nil
}

func (implementation Implementation) Move(stream language.Stream, location language.String) language.Statement {
	panic(implementation.Name()+".Move() Unimplemented")
	return language.Statement("")
}

