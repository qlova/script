package Go

import "fmt"
import "strconv"
import "reflect"
import "github.com/qlova/script/language"

type Implementation struct {
	Imports map[language.Statement]struct{}
}

func (implementation Implementation) String(s string) language.String {
	return String{
		Expression: language.Statement(strconv.Quote(s)),
	}
}

func (implementation Implementation) Integer(i int) language.Integer { 
	return Integer {
		Expression: language.Statement(fmt.Sprint(i)),
	}
}

func (implementation Implementation) Real(r float64) language.Real { return nil }
func (implementation Implementation) Rational() language.Rational { return nil }
func (implementation Implementation) Natural(n uint) language.Natural { return nil }

func (implementation Implementation) Quaternion() language.Quaternion { return nil }
func (implementation Implementation) Sedenion() language.Sedenion { return nil }
func (implementation Implementation) Duplex() language.Duplex { return nil }
func (implementation Implementation) Complex() language.Complex { return nil }
func (implementation Implementation) Octonion() language.Octonion { return nil }
func (implementation Implementation) Color() language.Color { return nil }
func (implementation Implementation) Video() language.Video { return nil }
func (implementation Implementation) Stream() language.Stream { return nil }
func (implementation Implementation) Symbol(r rune) language.Symbol { return nil }
func (implementation Implementation) Byte(b byte) language.Byte { return nil }
func (implementation Implementation) Image() language.Image { return nil }
func (implementation Implementation) Sound() language.Sound { return nil }
func (implementation Implementation) Time() language.Time { return nil }
func (implementation Implementation) Bit(b bool) language.Bit { return nil }

func (implementation Implementation) Name() string {
	return "Go"
}

func Language() Implementation {
	var implementation Implementation
	implementation.Imports = make(map[language.Statement]struct{})
	return implementation
}

func (implementation Implementation) Init() {
	panic(implementation.Name()+".Init() Unimplemented")
}

func (implementation Implementation) Build(path string) func() {
	panic(implementation.Name()+".Build() Unimplemented")
	return nil
}

func (implementation Implementation) Import(path language.Statement) {
	implementation.Imports[path] = struct{}{}
}

func (implementation Implementation) ExpressionOf(t language.Type) language.Statement {
	return reflect.ValueOf(t).Convert(reflect.TypeOf(language.NewType{})).Interface().(language.NewType).Expression
}
