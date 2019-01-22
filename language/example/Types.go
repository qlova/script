package Example

import "github.com/qlova/script/language"

type Quaternion language.NewType
func (t Quaternion) Number() {}
func (t Quaternion) Quaternion() {}
func (t Quaternion) Name() string { return "quaternion" }
func (t Quaternion) Is(b language.Type) bool { _, ok := b.(Quaternion); return ok }
func (t Quaternion) Register(name string) langauge.Type { return Quaternion{Expression: name} }

type Octonion language.NewType
func (t Octonion) Number() {}
func (t Octonion) Octonion() {}
func (t Octonion) Name() string { return "octonion" }
func (t Octonion) Is(b language.Type) bool { _, ok := b.(Octonion); return ok }
func (t Octonion) Register(name string) langauge.Type { return Octonion{Expression: name} }

type Natural language.NewType
func (t Natural) Number() {}
func (t Natural) Natural() {}
func (t Natural) Name() string { return "natural" }
func (t Natural) Is(b language.Type) bool { _, ok := b.(Natural); return ok }
func (t Natural) Register(name string) langauge.Type { return Natural{Expression: name} }

type Integer language.NewType
func (t Integer) Number() {}
func (t Integer) Integer() {}
func (t Integer) Name() string { return "integer" }
func (t Integer) Is(b language.Type) bool { _, ok := b.(Integer); return ok }
func (t Integer) Register(name string) langauge.Type { return Integer{Expression: name} }

type Duplex language.NewType
func (t Duplex) Number() {}
func (t Duplex) Duplex() {}
func (t Duplex) Name() string { return "duplex" }
func (t Duplex) Is(b language.Type) bool { _, ok := b.(Duplex); return ok }
func (t Duplex) Register(name string) langauge.Type { return Duplex{Expression: name} }

type Complex language.NewType
func (t Complex) Number() {}
func (t Complex) Complex() {}
func (t Complex) Name() string { return "complex" }
func (t Complex) Is(b language.Type) bool { _, ok := b.(Complex); return ok }
func (t Complex) Register(name string) langauge.Type { return Complex{Expression: name} }

type Sedenion language.NewType
func (t Sedenion) Number() {}
func (t Sedenion) Sedenion() {}
func (t Sedenion) Name() string { return "sedenion" }
func (t Sedenion) Is(b language.Type) bool { _, ok := b.(Sedenion); return ok }
func (t Sedenion) Register(name string) langauge.Type { return Sedenion{Expression: name} }

type Real language.NewType
func (t Real) Number() {}
func (t Real) Real() {}
func (t Real) Name() string { return "real" }
func (t Real) Is(b language.Type) bool { _, ok := b.(Real); return ok }
func (t Real) Register(name string) langauge.Type { return Real{Expression: name} }

type Rational language.NewType
func (t Rational) Number() {}
func (t Rational) Rational() {}
func (t Rational) Name() string { return "rational" }
func (t Rational) Is(b language.Type) bool { _, ok := b.(Rational); return ok }
func (t Rational) Register(name string) langauge.Type { return Rational{Expression: name} }

type Sound language.NewType
func (t Sound) Number() {}
func (t Sound) Sound() {}
func (t Sound) Name() string { return "sound" }
func (t Sound) Is(b language.Type) bool { _, ok := b.(Sound); return ok }
func (t Sound) Register(name string) langauge.Type { return Sound{Expression: name} }

type Stream language.NewType
func (t Stream) Number() {}
func (t Stream) Stream() {}
func (t Stream) Name() string { return "stream" }
func (t Stream) Is(b language.Type) bool { _, ok := b.(Stream); return ok }
func (t Stream) Register(name string) langauge.Type { return Stream{Expression: name} }

type String language.NewType
func (t String) Number() {}
func (t String) String() {}
func (t String) Name() string { return "string" }
func (t String) Is(b language.Type) bool { _, ok := b.(String); return ok }
func (t String) Register(name string) langauge.Type { return String{Expression: name} }

type Bit language.NewType
func (t Bit) Number() {}
func (t Bit) Bit() {}
func (t Bit) Name() string { return "bit" }
func (t Bit) Is(b language.Type) bool { _, ok := b.(Bit); return ok }
func (t Bit) Register(name string) langauge.Type { return Bit{Expression: name} }

type Color language.NewType
func (t Color) Number() {}
func (t Color) Color() {}
func (t Color) Name() string { return "color" }
func (t Color) Is(b language.Type) bool { _, ok := b.(Color); return ok }
func (t Color) Register(name string) langauge.Type { return Color{Expression: name} }

type Video language.NewType
func (t Video) Number() {}
func (t Video) Video() {}
func (t Video) Name() string { return "video" }
func (t Video) Is(b language.Type) bool { _, ok := b.(Video); return ok }
func (t Video) Register(name string) langauge.Type { return Video{Expression: name} }

type Time language.NewType
func (t Time) Number() {}
func (t Time) Time() {}
func (t Time) Name() string { return "time" }
func (t Time) Is(b language.Type) bool { _, ok := b.(Time); return ok }
func (t Time) Register(name string) langauge.Type { return Time{Expression: name} }

type Symbol language.NewType
func (t Symbol) Number() {}
func (t Symbol) Symbol() {}
func (t Symbol) Name() string { return "symbol" }
func (t Symbol) Is(b language.Type) bool { _, ok := b.(Symbol); return ok }
func (t Symbol) Register(name string) langauge.Type { return Symbol{Expression: name} }

type Byte language.NewType
func (t Byte) Number() {}
func (t Byte) Byte() {}
func (t Byte) Name() string { return "byte" }
func (t Byte) Is(b language.Type) bool { _, ok := b.(Byte); return ok }
func (t Byte) Register(name string) langauge.Type { return Byte{Expression: name} }

type Image language.NewType
func (t Image) Number() {}
func (t Image) Image() {}
func (t Image) Name() string { return "image" }
func (t Image) Is(b language.Type) bool { _, ok := b.(Image); return ok }
func (t Image) Register(name string) langauge.Type { return Image{Expression: name} }

type Error language.NewType
func (t Error) Number() {}
func (t Error) Error() {}
func (t Error) Name() string { return "error" }
func (t Error) Is(b language.Type) bool { _, ok := b.(Error); return ok }
func (t Error) Register(name string) langauge.Type { return Error{Expression: name} }

type Array language.NewType
func (t Array) Number() {}
func (t Array) Array() {}
func (t Array) Name() string { return "array" }
func (t Array) Is(b language.Type) bool { _, ok := b.(Array); return ok }
func (t Array) Register(name string) langauge.Type { return Array{Expression: name} }

type Ring language.NewType
func (t Ring) Number() {}
func (t Ring) Ring() {}
func (t Ring) Name() string { return "ring" }
func (t Ring) Is(b language.Type) bool { _, ok := b.(Ring); return ok }
func (t Ring) Register(name string) langauge.Type { return Ring{Expression: name} }

type Tree language.NewType
func (t Tree) Number() {}
func (t Tree) Tree() {}
func (t Tree) Name() string { return "tree" }
func (t Tree) Is(b language.Type) bool { _, ok := b.(Tree); return ok }
func (t Tree) Register(name string) langauge.Type { return Tree{Expression: name} }

type Table language.NewType
func (t Table) Number() {}
func (t Table) Table() {}
func (t Table) Name() string { return "table" }
func (t Table) Is(b language.Type) bool { _, ok := b.(Table); return ok }
func (t Table) Register(name string) langauge.Type { return Table{Expression: name} }

type Pointer language.NewType
func (t Pointer) Number() {}
func (t Pointer) Pointer() {}
func (t Pointer) Name() string { return "pointer" }
func (t Pointer) Is(b language.Type) bool { _, ok := b.(Pointer); return ok }
func (t Pointer) Register(name string) langauge.Type { return Pointer{Expression: name} }

type Graph language.NewType
func (t Graph) Number() {}
func (t Graph) Graph() {}
func (t Graph) Name() string { return "graph" }
func (t Graph) Is(b language.Type) bool { _, ok := b.(Graph); return ok }
func (t Graph) Register(name string) langauge.Type { return Graph{Expression: name} }

type Function language.NewType
func (t Function) Number() {}
func (t Function) Function() {}
func (t Function) Name() string { return "function" }
func (t Function) Is(b language.Type) bool { _, ok := b.(Function); return ok }
func (t Function) Register(name string) langauge.Type { return Function{Expression: name} }

type List language.NewType
func (t List) Number() {}
func (t List) List() {}
func (t List) Name() string { return "list" }
func (t List) Is(b language.Type) bool { _, ok := b.(List); return ok }
func (t List) Register(name string) langauge.Type { return List{Expression: name} }

type Set language.NewType
func (t Set) Number() {}
func (t Set) Set() {}
func (t Set) Name() string { return "set" }
func (t Set) Is(b language.Type) bool { _, ok := b.(Set); return ok }
func (t Set) Register(name string) langauge.Type { return Set{Expression: name} }

type Vector language.NewType
func (t Vector) Number() {}
func (t Vector) Vector() {}
func (t Vector) Name() string { return "vector" }
func (t Vector) Is(b language.Type) bool { _, ok := b.(Vector); return ok }
func (t Vector) Register(name string) langauge.Type { return Vector{Expression: name} }

type Queue language.NewType
func (t Queue) Number() {}
func (t Queue) Queue() {}
func (t Queue) Name() string { return "queue" }
func (t Queue) Is(b language.Type) bool { _, ok := b.(Queue); return ok }
func (t Queue) Register(name string) langauge.Type { return Queue{Expression: name} }

type Tensor language.NewType
func (t Tensor) Number() {}
func (t Tensor) Tensor() {}
func (t Tensor) Name() string { return "tensor" }
func (t Tensor) Is(b language.Type) bool { _, ok := b.(Tensor); return ok }
func (t Tensor) Register(name string) langauge.Type { return Tensor{Expression: name} }

type Matrix language.NewType
func (t Matrix) Number() {}
func (t Matrix) Matrix() {}
func (t Matrix) Name() string { return "matrix" }
func (t Matrix) Is(b language.Type) bool { _, ok := b.(Matrix); return ok }
func (t Matrix) Register(name string) langauge.Type { return Matrix{Expression: name} }

