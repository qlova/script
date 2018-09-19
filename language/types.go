package language

import "math/big"

type BooleanType struct {
	Expression string
	Literal *bool
}
	func (BooleanType) Name() string { return "boolean" }
	func (BooleanType) SameAs(i interface{}) bool { _, ok := i.(BooleanType); return ok }
	func (BooleanType) Boolean() {}

type NumberType struct {
	Expression string
	Literal *big.Int
}
	func (NumberType) Name() string { return "number" }
	func (NumberType) SameAs(i interface{}) bool { _, ok := i.(NumberType); return ok }
	func (NumberType) Number() {}
	func (NumberType) Numeric() {}

type StringType struct {
	Expression string
	Literal string
}
	func (StringType) Name() string { return "string" }
	func (StringType) SameAs(i interface{}) bool { _, ok := i.(StringType); return ok }
	func (StringType) String() {}
	func (StringType) Numeric() {}

type ArrayType struct {
	Expression string

	Subtype Type
	Size int
	
	Full bool
}
	func (ArrayType) Name() string { return "array" }
	func (ArrayType) SameAs(i interface{}) bool { _, ok := i.(ArrayType); return ok }
	func (ArrayType) Array() {}
	func (a ArrayType) Length() int { return a.Size }
	func (a ArrayType) SubType() Type { return a.Subtype }
	
type FunctionType struct {
	Expression string

	Args []Type
	Rets Type
}
	func (FunctionType) Name() string { return "function" }
	func (FunctionType) SameAs(i interface{}) bool { _, ok := i.(FunctionType); return ok }
	func (FunctionType) Function() {}
	func (a FunctionType) Arguments() []Type { return a.Args }
	func (a FunctionType) Returns() Type { return a.Rets }
	
type ErrorType struct {
	Expression string

	Message string
	Code int
}
	func (ErrorType) Name() string { return "error" }
	func (ErrorType) SameAs(i interface{}) bool { _, ok := i.(ErrorType); return ok }
	func (ErrorType) Error() {}
