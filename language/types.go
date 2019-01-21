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
	Register(name string) Type
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
type Complex interface {
	Type
	Number()
	Complex()
}
type Sedenion interface {
	Type
	Number()
	Sedenion()
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
type String interface {
	Type
	String()
}
type Bit interface {
	Type
	Bit()
}
type Color interface {
	Type
	Color()
}
type Sound interface {
	Type
	Sound()
}
type Stream interface {
	Type
	Stream()
}
type Symbol interface {
	Type
	Symbol()
}
type Byte interface {
	Type
	Byte()
}
type Image interface {
	Type
	Image()
}
type Video interface {
	Type
	Video()
}
type Time interface {
	Type
	Time()
}
type List interface {
	Type
	List()
}
type Set interface {
	Type
	Set()
}
type Vector interface {
	Type
	Vector()
}
type Queue interface {
	Type
	Queue()
}
type Tensor interface {
	Type
	Tensor()
}
type Matrix interface {
	Type
	Matrix()
}
type Error interface {
	Type
	Error()
}
type Array interface {
	Type
	Array()
}
type Ring interface {
	Type
	Ring()
}
type Tree interface {
	Type
	Tree()
}
type Table interface {
	Type
	Table()
}
type Pointer interface {
	Type
	Pointer()
}
type Graph interface {
	Type
	Graph()
}
type Function interface {
	Type
	Function()
}
