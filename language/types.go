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
type Quaternion interface {
	Type
	Number()
	Quaternion()
}
type Real interface {
	Type
	Number()
	Real()
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
type Octonion interface {
	Type
	Number()
	Octonion()
}
type Sedenion interface {
	Type
	Number()
	Sedenion()
}
type Bit interface {
	Type
	Bit()
}
type Byte interface {
	Type
	Byte()
}
type Sound interface {
	Type
	Sound()
}
type Time interface {
	Type
	Time()
}
type Stream interface {
	Type
	Stream()
}
type Symbol interface {
	Type
	Symbol()
}
type String interface {
	Type
	String()
}
type Color interface {
	Type
	Color()
}
type Image interface {
	Type
	Image()
}
type Video interface {
	Type
	Video()
}
type Array interface {
	Type
	Array()
}
type Vector interface {
	Type
	Vector()
}
type Table interface {
	Type
	Table()
}
type Tensor interface {
	Type
	Tensor()
}
type Matrix interface {
	Type
	Matrix()
}
type Ring interface {
	Type
	Ring()
}
type Tree interface {
	Type
	Tree()
}
type Function interface {
	Type
	Function()
}
type Error interface {
	Type
	Error()
}
type List interface {
	Type
	List()
}
type Set interface {
	Type
	Set()
}
type Queue interface {
	Type
	Queue()
}
type Pointer interface {
	Type
	Pointer()
}
type Graph interface {
	Type
	Graph()
}
