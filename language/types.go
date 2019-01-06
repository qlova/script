package language
	
type NewType struct {
	Subtype Type
	Expression Statement
	Literal interface{}
}

type Statement string

type Type interface { 
	Name() string
	Is(Type) bool
}

type Number interface {
	Type
	Number()
}

type Dynamic interface {
	Type
	Dynamic()
}

type Metatype interface {
	Type
	Metatype()
}

type Quaternion interface {
	Type
	Number()
	Quaternion()
}
type Octonion interface {
	Type
	Number()
	Octonion()
}
type Real interface {
	Type
	Number()
	Real()
}
type Rational interface {
	Type
	Number()
	Rational()
}
type Natural interface {
	Type
	Number()
	Natural()
}
type Complex interface {
	Type
	Number()
	Complex()
}
type Integer interface {
	Type
	Number()
	Integer()
}
type Duplex interface {
	Type
	Number()
	Duplex()
}
type Sedenion interface {
	Type
	Number()
	Sedenion()
}
type String interface {
	Type
	String()
}
type Sound interface {
	Type
	Sound()
}
type Video interface {
	Type
	Video()
}
type Time interface {
	Type
	Time()
}
type Symbol interface {
	Type
	Symbol()
}
type Bit interface {
	Type
	Bit()
}
type Byte interface {
	Type
	Byte()
}
type Color interface {
	Type
	Color()
}
type Image interface {
	Type
	Image()
}
type Stream interface {
	Type
	Stream()
}
type Graph interface {
	Type
	Graph()
}
type Matrix interface {
	Type
	Matrix()
}
type Pointer interface {
	Type
	Pointer()
}
type Vector interface {
	Type
	Vector()
}
type Tree interface {
	Type
	Tree()
}
type Table interface {
	Type
	Table()
}
type Tensor interface {
	Type
	Tensor()
}
type List interface {
	Type
	List()
}
type Set interface {
	Type
	Set()
}
type Error interface {
	Type
	Error()
}
type Array interface {
	Type
	Array()
}
type Function interface {
	Type
	Function()
}
type Queue interface {
	Type
	Queue()
}
type Ring interface {
	Type
	Ring()
}
