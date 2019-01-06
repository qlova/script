package Example

import "github.com/qlova/script/language"

func (implementation Implementation) Add(a, b language.Number) language.Number {
	panic(implementation.Name()+".Add() Unimplemented")
	return nil
}

func (implementation Implementation) Sub(a, b language.Number) language.Number {
	panic(implementation.Name()+".Sub() Unimplemented")
	return nil
}

func (implementation Implementation) Mul(a, b language.Number) language.Number {
	panic(implementation.Name()+".Mul() Unimplemented")
	return nil
}

func (implementation Implementation) Div(a, b language.Number) language.Number {
	panic(implementation.Name()+".Div() Unimplemented")
	return nil
}

func (implementation Implementation) Pow(a, b language.Number) language.Number {
	panic(implementation.Name()+".Pow() Unimplemented")
	return nil
}

func (implementation Implementation) Mod(a, b language.Number) language.Number {
	panic(implementation.Name()+".Mod() Unimplemented")
	return nil
}

