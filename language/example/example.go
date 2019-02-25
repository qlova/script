package example

import "github.com/qlova/script/language"

type Implementation struct {}

func (implementation Implementation) Real(r float64) language.Real { panic(`Unimplemented`); return nil }
func (implementation Implementation) Complex() language.Complex { panic(`Unimplemented`); return nil }
func (implementation Implementation) Quaternion() language.Quaternion { panic(`Unimplemented`); return nil }
func (implementation Implementation) Octonion() language.Octonion { panic(`Unimplemented`); return nil }
func (implementation Implementation) Sedenion() language.Sedenion { panic(`Unimplemented`); return nil }
func (implementation Implementation) Rational() language.Rational { panic(`Unimplemented`); return nil }
func (implementation Implementation) Natural(n uint) language.Natural { panic(`Unimplemented`); return nil }
func (implementation Implementation) Integer(i int) language.Integer { panic(`Unimplemented`); return nil }
func (implementation Implementation) Duplex() language.Duplex { panic(`Unimplemented`); return nil }
func (implementation Implementation) Symbol(r rune) language.Symbol { panic(`Unimplemented`); return nil }
func (implementation Implementation) Color() language.Color { panic(`Unimplemented`); return nil }
func (implementation Implementation) Image() language.Image { panic(`Unimplemented`); return nil }
func (implementation Implementation) String(s string) language.String { panic(`Unimplemented`); return nil }
func (implementation Implementation) Bit(b bool) language.Bit { panic(`Unimplemented`); return nil }
func (implementation Implementation) Byte(b byte) language.Byte { panic(`Unimplemented`); return nil }
func (implementation Implementation) Sound() language.Sound { panic(`Unimplemented`); return nil }
func (implementation Implementation) Video() language.Video { panic(`Unimplemented`); return nil }
func (implementation Implementation) Time() language.Time { panic(`Unimplemented`); return nil }
func (implementation Implementation) Stream() language.Stream { panic(`Unimplemented`); return nil }
func (implementation Implementation) Name() string {
	panic(implementation.Name()+".Name() Unimplemented")
	return ""
}

func (implementation Implementation) Init() {
	panic(implementation.Name()+".Init() Unimplemented")
}

func (implementation Implementation) Build(path string) func() {
	panic(implementation.Name()+".Build() Unimplemented")
	return nil
}

