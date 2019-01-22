package interpreter

import "github.com/qlova/script/language"

func (implementation Implementation) PointerOf(t language.Type) language.Pointer {
	panic(implementation.Name()+".PointerOf() Unimplemented")
	return nil
}

func (implementation Implementation) Dereference(p language.Pointer) language.Type {
	panic(implementation.Name()+".Dereference() Unimplemented")
	return nil
}

