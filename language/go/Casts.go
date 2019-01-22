package Go

import "github.com/qlova/script/language"

func (implementation Implementation) Cast(a, b language.Type) language.Type {
	panic(implementation.Name()+".Cast() Unimplemented")
	return nil
}

