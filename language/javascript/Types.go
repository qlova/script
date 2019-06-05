package Javascript

import "github.com/qlova/script/language"

type Integer language.NewType

func (t Integer) Number()                 {}
func (t Integer) Integer()                {}
func (t Integer) Name() string            { return "integer" }
func (t Integer) Is(b language.Type) bool { _, ok := b.(Integer); return ok }
func (t Integer) Register(name string) language.Type {
	return Integer{Expression: language.Statement(name)}
}
func (t Integer) Raw() language.Statement { return t.Expression }

type Duplex language.NewType

func (t Duplex) Number()                 {}
func (t Duplex) Duplex()                 {}
func (t Duplex) Name() string            { return "duplex" }
func (t Duplex) Is(b language.Type) bool { _, ok := b.(Duplex); return ok }
func (t Duplex) Register(name string) language.Type {
	return Duplex{Expression: language.Statement(name)}
}
func (t Duplex) Raw() language.Statement { return t.Expression }

type Quaternion language.NewType

func (t Quaternion) Number()                 {}
func (t Quaternion) Quaternion()             {}
func (t Quaternion) Name() string            { return "quaternion" }
func (t Quaternion) Is(b language.Type) bool { _, ok := b.(Quaternion); return ok }
func (t Quaternion) Register(name string) language.Type {
	return Quaternion{Expression: language.Statement(name)}
}
func (t Quaternion) Raw() language.Statement { return t.Expression }

type Real language.NewType

func (t Real) Number()                            {}
func (t Real) Real()                              {}
func (t Real) Name() string                       { return "real" }
func (t Real) Is(b language.Type) bool            { _, ok := b.(Real); return ok }
func (t Real) Register(name string) language.Type { return Real{Expression: language.Statement(name)} }
func (t Real) Raw() language.Statement            { return t.Expression }

type Natural language.NewType

func (t Natural) Number()                 {}
func (t Natural) Natural()                {}
func (t Natural) Name() string            { return "natural" }
func (t Natural) Is(b language.Type) bool { _, ok := b.(Natural); return ok }
func (t Natural) Register(name string) language.Type {
	return Natural{Expression: language.Statement(name)}
}
func (t Natural) Raw() language.Statement { return t.Expression }

type Octonion language.NewType

func (t Octonion) Number()                 {}
func (t Octonion) Octonion()               {}
func (t Octonion) Name() string            { return "octonion" }
func (t Octonion) Is(b language.Type) bool { _, ok := b.(Octonion); return ok }
func (t Octonion) Register(name string) language.Type {
	return Octonion{Expression: language.Statement(name)}
}
func (t Octonion) Raw() language.Statement { return t.Expression }

type Sedenion language.NewType

func (t Sedenion) Number()                 {}
func (t Sedenion) Sedenion()               {}
func (t Sedenion) Name() string            { return "sedenion" }
func (t Sedenion) Is(b language.Type) bool { _, ok := b.(Sedenion); return ok }
func (t Sedenion) Register(name string) language.Type {
	return Sedenion{Expression: language.Statement(name)}
}
func (t Sedenion) Raw() language.Statement { return t.Expression }

type Rational language.NewType

func (t Rational) Number()                 {}
func (t Rational) Rational()               {}
func (t Rational) Name() string            { return "rational" }
func (t Rational) Is(b language.Type) bool { _, ok := b.(Rational); return ok }
func (t Rational) Register(name string) language.Type {
	return Rational{Expression: language.Statement(name)}
}
func (t Rational) Raw() language.Statement { return t.Expression }

type Complex language.NewType

func (t Complex) Number()                 {}
func (t Complex) Complex()                {}
func (t Complex) Name() string            { return "complex" }
func (t Complex) Is(b language.Type) bool { _, ok := b.(Complex); return ok }
func (t Complex) Register(name string) language.Type {
	return Complex{Expression: language.Statement(name)}
}
func (t Complex) Raw() language.Statement { return t.Expression }

type Color language.NewType

func (t Color) Number()                            {}
func (t Color) Color()                             {}
func (t Color) Name() string                       { return "color" }
func (t Color) Is(b language.Type) bool            { _, ok := b.(Color); return ok }
func (t Color) Register(name string) language.Type { return Color{Expression: language.Statement(name)} }
func (t Color) Raw() language.Statement            { return t.Expression }

type Stream language.NewType

func (t Stream) Number()                 {}
func (t Stream) Stream()                 {}
func (t Stream) Name() string            { return "stream" }
func (t Stream) Is(b language.Type) bool { _, ok := b.(Stream); return ok }
func (t Stream) Register(name string) language.Type {
	return Stream{Expression: language.Statement(name)}
}
func (t Stream) Raw() language.Statement { return t.Expression }

type Symbol language.NewType

func (t Symbol) Number()                 {}
func (t Symbol) Symbol()                 {}
func (t Symbol) Name() string            { return "symbol" }
func (t Symbol) Is(b language.Type) bool { _, ok := b.(Symbol); return ok }
func (t Symbol) Register(name string) language.Type {
	return Symbol{Expression: language.Statement(name)}
}
func (t Symbol) Raw() language.Statement { return t.Expression }

type Bit language.NewType

func (t Bit) Number()                            {}
func (t Bit) Bit()                               {}
func (t Bit) Name() string                       { return "bit" }
func (t Bit) Is(b language.Type) bool            { _, ok := b.(Bit); return ok }
func (t Bit) Register(name string) language.Type { return Bit{Expression: language.Statement(name)} }
func (t Bit) Raw() language.Statement            { return t.Expression }

type Byte language.NewType

func (t Byte) Number()                            {}
func (t Byte) Byte()                              {}
func (t Byte) Name() string                       { return "byte" }
func (t Byte) Is(b language.Type) bool            { _, ok := b.(Byte); return ok }
func (t Byte) Register(name string) language.Type { return Byte{Expression: language.Statement(name)} }
func (t Byte) Raw() language.Statement            { return t.Expression }

type Video language.NewType

func (t Video) Number()                            {}
func (t Video) Video()                             {}
func (t Video) Name() string                       { return "video" }
func (t Video) Is(b language.Type) bool            { _, ok := b.(Video); return ok }
func (t Video) Register(name string) language.Type { return Video{Expression: language.Statement(name)} }
func (t Video) Raw() language.Statement            { return t.Expression }

type Time language.NewType

func (t Time) Number()                            {}
func (t Time) Time()                              {}
func (t Time) Name() string                       { return "time" }
func (t Time) Is(b language.Type) bool            { _, ok := b.(Time); return ok }
func (t Time) Register(name string) language.Type { return Time{Expression: language.Statement(name)} }
func (t Time) Raw() language.Statement            { return t.Expression }

type String language.NewType

func (t String) Number()                 {}
func (t String) String()                 {}
func (t String) Name() string            { return "string" }
func (t String) Is(b language.Type) bool { _, ok := b.(String); return ok }
func (t String) Register(name string) language.Type {
	return String{Expression: language.Statement(name)}
}
func (t String) Raw() language.Statement { return t.Expression }

type Image language.NewType

func (t Image) Number()                            {}
func (t Image) Image()                             {}
func (t Image) Name() string                       { return "image" }
func (t Image) Is(b language.Type) bool            { _, ok := b.(Image); return ok }
func (t Image) Register(name string) language.Type { return Image{Expression: language.Statement(name)} }
func (t Image) Raw() language.Statement            { return t.Expression }

type Sound language.NewType

func (t Sound) Number()                            {}
func (t Sound) Sound()                             {}
func (t Sound) Name() string                       { return "sound" }
func (t Sound) Is(b language.Type) bool            { _, ok := b.(Sound); return ok }
func (t Sound) Register(name string) language.Type { return Sound{Expression: language.Statement(name)} }
func (t Sound) Raw() language.Statement            { return t.Expression }

type Error language.NewType

func (t Error) Number()                            {}
func (t Error) Error()                             {}
func (t Error) Name() string                       { return "error" }
func (t Error) Is(b language.Type) bool            { _, ok := b.(Error); return ok }
func (t Error) Register(name string) language.Type { return Error{Expression: language.Statement(name)} }
func (t Error) Raw() language.Statement            { return t.Expression }

type Tensor language.NewType

func (t Tensor) Number()                 {}
func (t Tensor) Tensor()                 {}
func (t Tensor) Name() string            { return "tensor" }
func (t Tensor) Is(b language.Type) bool { _, ok := b.(Tensor); return ok }
func (t Tensor) Register(name string) language.Type {
	return Tensor{Expression: language.Statement(name)}
}
func (t Tensor) Raw() language.Statement { return t.Expression }

type Tree language.NewType

func (t Tree) Number()                            {}
func (t Tree) Tree()                              {}
func (t Tree) Name() string                       { return "tree" }
func (t Tree) Is(b language.Type) bool            { _, ok := b.(Tree); return ok }
func (t Tree) Register(name string) language.Type { return Tree{Expression: language.Statement(name)} }
func (t Tree) Raw() language.Statement            { return t.Expression }

type Array language.NewType

func (t Array) Number()                            {}
func (t Array) Array()                             {}
func (t Array) Name() string                       { return "array" }
func (t Array) Is(b language.Type) bool            { _, ok := b.(Array); return ok }
func (t Array) Register(name string) language.Type { return Array{Expression: language.Statement(name)} }
func (t Array) Raw() language.Statement            { return t.Expression }

type List language.NewType

func (t List) Number()                            {}
func (t List) List()                              {}
func (t List) Name() string                       { return "list" }
func (t List) Is(b language.Type) bool            { _, ok := b.(List); return ok }
func (t List) Register(name string) language.Type { return List{Expression: language.Statement(name)} }
func (t List) Raw() language.Statement            { return t.Expression }

type Ring language.NewType

func (t Ring) Number()                            {}
func (t Ring) Ring()                              {}
func (t Ring) Name() string                       { return "ring" }
func (t Ring) Is(b language.Type) bool            { _, ok := b.(Ring); return ok }
func (t Ring) Register(name string) language.Type { return Ring{Expression: language.Statement(name)} }
func (t Ring) Raw() language.Statement            { return t.Expression }

type Graph language.NewType

func (t Graph) Number()                            {}
func (t Graph) Graph()                             {}
func (t Graph) Name() string                       { return "graph" }
func (t Graph) Is(b language.Type) bool            { _, ok := b.(Graph); return ok }
func (t Graph) Register(name string) language.Type { return Graph{Expression: language.Statement(name)} }
func (t Graph) Raw() language.Statement            { return t.Expression }

type Function language.NewType

func (t Function) Number()                 {}
func (t Function) Function()               {}
func (t Function) Name() string            { return "function" }
func (t Function) Is(b language.Type) bool { _, ok := b.(Function); return ok }
func (t Function) Register(name string) language.Type {
	return Function{Expression: language.Statement(name)}
}
func (t Function) Raw() language.Statement { return t.Expression }

type Table language.NewType

func (t Table) Number()                            {}
func (t Table) Table()                             {}
func (t Table) Name() string                       { return "table" }
func (t Table) Is(b language.Type) bool            { _, ok := b.(Table); return ok }
func (t Table) Register(name string) language.Type { return Table{Expression: language.Statement(name)} }
func (t Table) Raw() language.Statement            { return t.Expression }

type Queue language.NewType

func (t Queue) Number()                            {}
func (t Queue) Queue()                             {}
func (t Queue) Name() string                       { return "queue" }
func (t Queue) Is(b language.Type) bool            { _, ok := b.(Queue); return ok }
func (t Queue) Register(name string) language.Type { return Queue{Expression: language.Statement(name)} }
func (t Queue) Raw() language.Statement            { return t.Expression }

type Vector language.NewType

func (t Vector) Number()                 {}
func (t Vector) Vector()                 {}
func (t Vector) Name() string            { return "vector" }
func (t Vector) Is(b language.Type) bool { _, ok := b.(Vector); return ok }
func (t Vector) Register(name string) language.Type {
	return Vector{Expression: language.Statement(name)}
}
func (t Vector) Raw() language.Statement { return t.Expression }

type Set language.NewType

func (t Set) Number()                            {}
func (t Set) Set()                               {}
func (t Set) Name() string                       { return "set" }
func (t Set) Is(b language.Type) bool            { _, ok := b.(Set); return ok }
func (t Set) Register(name string) language.Type { return Set{Expression: language.Statement(name)} }
func (t Set) Raw() language.Statement            { return t.Expression }

type Matrix language.NewType

func (t Matrix) Number()                 {}
func (t Matrix) Matrix()                 {}
func (t Matrix) Name() string            { return "matrix" }
func (t Matrix) Is(b language.Type) bool { _, ok := b.(Matrix); return ok }
func (t Matrix) Register(name string) language.Type {
	return Matrix{Expression: language.Statement(name)}
}
func (t Matrix) Raw() language.Statement { return t.Expression }

type Pointer language.NewType

func (t Pointer) Number()                 {}
func (t Pointer) Pointer()                {}
func (t Pointer) Name() string            { return "pointer" }
func (t Pointer) Is(b language.Type) bool { _, ok := b.(Pointer); return ok }
func (t Pointer) Register(name string) language.Type {
	return Pointer{Expression: language.Statement(name)}
}
func (t Pointer) Raw() language.Statement { return t.Expression }

type Native language.NewType

func (t Native) Native()                 {}
func (t Native) Name() string            { return "native" }
func (t Native) Is(b language.Type) bool { _, ok := b.(Native); return ok }
func (t Native) Register(name string) language.Type {
	return Native{Expression: language.Statement(name)}
}
func (t Native) Raw() language.Statement { return t.Expression }
