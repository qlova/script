package Go

import "reflect"
import "github.com/qlova/script/language"

func (implementation Implementation) Register(register string, value language.Type) language.Statement {
	return language.Statement("var "+register+" = ")+implementation.ExpressionOf(value)+"\n"
}

func (implementation Implementation) Set(register string, value language.Type) language.Statement {
	panic(implementation.Name()+".Set() Unimplemented")
	return language.Statement("")
}

func (implementation Implementation) Get(register string, value language.Type) language.Type {
	var result language.NewType
	result.Expression = language.Statement(register)
	return reflect.ValueOf(result).Convert(reflect.TypeOf(value)).Interface().(language.Type)
}

