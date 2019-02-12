package language
	
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

type Integer interface {
	Type
	Number()
	Integer()
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
type Natural interface {
	Type
	Number()
	Natural()
}
type Duplex interface {
	Type
	Number()
	Duplex()
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
type Sound interface {
	Type
	Sound()
}
type Video interface {
	Type
	Video()
}
type Symbol interface {
	Type
	Symbol()
}
type Queue interface {
	Type
	Queue()
}
type Tensor interface {
	Type
	Tensor()
}
type Tree interface {
	Type
	Tree()
}
type Error interface {
	Type
	Error()
}
type Array interface {
	Type
	Array()
}
type List interface {
	Type
	List()
}
type Table interface {
	Type
	Table()
}
type Vector interface {
	Type
	Vector()
}
type Matrix interface {
	Type
	Matrix()
}
type Set interface {
	Type
	Set()
}
type Ring interface {
	Type
	Ring()
}
type Function interface {
	Type
	Function()
}
type Pointer interface {
	Type
	Pointer()
}
type Graph interface {
	Type
	Graph()
}
