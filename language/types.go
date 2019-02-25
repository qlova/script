package language

type Buffer interface {
	Buffer()
}

type NewType struct {
	Custom string
	Subtype Type
	Expression Statement
	Literal interface{}
	Length int
}

func (t NewType) Name() string { return t.Custom }
func (t NewType) Is(b Type) bool { c, ok := b.(NewType); return ok && c.Custom == t.Custom }
func (t NewType) Register(name string) Type { return NewType{Expression: Statement(name)} }
func (t NewType) Raw() Statement { return t.Expression }

type Statement string

type Type interface { 
	Name() string
	Is(Type) bool
	Register(name string) Type
	Raw() Statement
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

type Real interface {
	Type
	Number()
	Real()
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
type Rational interface {
	Type
	Number()
	Rational()
}
type Video interface {
	Type
	Video()
}
type Time interface {
	Type
	Time()
}
type Stream interface {
	Type
	Stream()
}
type String interface {
	Type
	String()
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
type Symbol interface {
	Type
	Symbol()
}
type Color interface {
	Type
	Color()
}
type Image interface {
	Type
	Image()
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
type Function interface {
	Type
	Function()
}
type Array interface {
	Type
	Array()
}
type Table interface {
	Type
	Table()
}
type Vector interface {
	Type
	Vector()
}
type Graph interface {
	Type
	Graph()
}
type Error interface {
	Type
	Error()
}
type Set interface {
	Type
	Set()
}
type Pointer interface {
	Type
	Pointer()
}
type Tree interface {
	Type
	Tree()
}
type List interface {
	Type
	List()
}
type Queue interface {
	Type
	Queue()
}
