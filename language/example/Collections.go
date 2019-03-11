package example

import "github.com/qlova/script/language"

func (implementation Implementation) ArrayOf(t language.Type, length int) language.Array {
	panic(implementation.Name()+".ArrayOf() Unimplemented")
	return nil
}

func (implementation Implementation) Join(a, b language.Type) language.Type {
	panic(implementation.Name()+".Join() Unimplemented")
	return nil
}

func (implementation Implementation) Length(t language.Type) language.Integer {
	panic(implementation.Name()+".Length() Unimplemented")
	return nil
}

func (implementation Implementation) Push(value language.Type, list language.Type) language.Statement {
	panic(implementation.Name()+".Push() Unimplemented")
	return language.Statement("")
}

func (implementation Implementation) Pop(list language.Type) language.Type {
	panic(implementation.Name()+".Pop() Unimplemented")
	return nil
}

func (implementation Implementation) TableOf(t language.Type) language.Table {
	panic(implementation.Name()+".TableOf() Unimplemented")
	return nil
}

func (implementation Implementation) ListOf(t language.Type) language.List {
	panic(implementation.Name()+".ListOf() Unimplemented")
	return nil
}

func (implementation Implementation) List(t ...language.Type) language.List {
	panic(implementation.Name()+".List() Unimplemented")
	return nil
}

func (implementation Implementation) Array(t ...language.Type) language.Array {
	panic(implementation.Name()+".Array() Unimplemented")
	return nil
}

func (implementation Implementation) Table(index language.String, value language.Type) language.Table {
	panic(implementation.Name()+".Table() Unimplemented")
	return nil
}

