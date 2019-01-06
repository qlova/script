package Go

import "github.com/qlova/script/language"

func (implementation Implementation) DynamicOf(t language.Type) language.Dynamic {
	panic(implementation.Name()+".DynamicOf() Unimplemented")
	return nil
}

func (implementation Implementation) StaticOf(d language.Dynamic) language.Type {
	panic(implementation.Name()+".StaticOf() Unimplemented")
	return nil
}

func (implementation Implementation) MetatypeOf(t language.Type) language.Metatype {
	panic(implementation.Name()+".MetatypeOf() Unimplemented")
	return nil
}

