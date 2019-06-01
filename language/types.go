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

type Statement = string

type Buffer interface {
	Buffer()
}

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
type Octonion interface {
	Type
	Number()
	Octonion()
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
type Sedenion interface {
	Type
	Number()
	Sedenion()
}
type Time interface {
	Type
	Time()
}
type Symbol interface {
	Type
	Symbol()
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
type Image interface {
	Type
	Image()
}
type Video interface {
	Type
	Video()
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
type Graph interface {
	Type
	Graph()
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
type Pointer interface {
	Type
	Pointer()
}
type Set interface {
	Type
	Set()
}
type Table interface {
	Type
	Table()
}
type Tree interface {
	Type
	Tree()
}
type Vector interface {
	Type
	Vector()
}
type Array interface {
	Type
	Array()
}
type Ring interface {
	Type
	Ring()
}
type Native interface {
	Type
	Native()
}
type Native interface {
	Type
	Native()
}