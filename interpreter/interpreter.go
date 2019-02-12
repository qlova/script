package interpreter

import "strconv"
import "reflect"

import "github.com/qlova/script/language"
import "github.com/qlova/script/interpreter/dynamic"

type Implementation struct {
	symbols map[string]dynamic.BlockPointer
	program *dynamic.Program
	active *dynamic.BlockPointer
}

func (implementation Implementation) Start() {
	implementation.program.Run()
}

func New() Implementation {
	var implementation Implementation
	implementation.program = new(dynamic.Program)
	implementation.active = new(dynamic.BlockPointer)
	return implementation
}

func (implementation Implementation) String(s string) language.String {
	var register = implementation.ReserveRegister()
	implementation.AddInstruction(func(thread *dynamic.Thread) {
		thread.Set(register, s)
	})
	return String{Expression:language.Statement(strconv.Itoa(register))}
}

func (implementation Implementation) Integer(i int) language.Integer {
	var register = implementation.ReserveRegister()
	implementation.AddInstruction(func(thread *dynamic.Thread) {
		thread.Set(register, i)
	})
	return Integer{Expression:language.Statement(strconv.Itoa(register))}
}

func (implementation Implementation) Symbol(r rune) language.Symbol {
	var register = implementation.ReserveRegister()
	implementation.AddInstruction(func(thread *dynamic.Thread) {
		thread.Set(register, r)
	})
	return Symbol{Expression:language.Statement(strconv.Itoa(register))}
}

func (implementation Implementation) Rational() language.Rational { panic("Not implemented"); return nil }
func (implementation Implementation) Natural(n uint) language.Natural { panic("Not implemented"); return nil }
func (implementation Implementation) Complex() language.Complex { panic("Not implemented"); return nil }
func (implementation Implementation) Real(r float64) language.Real { panic("Not implemented"); return nil }

func (implementation Implementation) Duplex() language.Duplex { panic("Not implemented"); return nil }
func (implementation Implementation) Quaternion() language.Quaternion { panic("Not implemented"); return nil }
func (implementation Implementation) Octonion() language.Octonion { panic("Not implemented"); return nil }
func (implementation Implementation) Sedenion() language.Sedenion { panic("Not implemented"); return nil }
func (implementation Implementation) Byte(b byte) language.Byte { panic("Not implemented"); return nil }
func (implementation Implementation) Image() language.Image { panic("Not implemented"); return nil }
func (implementation Implementation) Sound() language.Sound { panic("Not implemented"); return nil }
func (implementation Implementation) Video() language.Video { panic("Not implemented"); return nil }
func (implementation Implementation) Time() language.Time { panic("Not implemented"); return nil }
func (implementation Implementation) Stream() language.Stream { panic("Not implemented"); return nil }
func (implementation Implementation) Bit(b bool) language.Bit { panic("Not implemented"); return nil }

func (implementation Implementation) Color() language.Color { panic("Not implemented"); return nil }

func (implementation Implementation) Name() string {
	return "interpreter"
}

func (implementation Implementation) Init() {

}

func (implementation Implementation) Build(path string) func() {
	panic(implementation.Name()+".Build() Unimplemented")
	return nil
}

func (implementation Implementation) Literal(t language.Type) interface{} {
	return reflect.ValueOf(t).Convert(reflect.TypeOf(language.NewType{})).Interface().(language.NewType).Literal
}

func (implementation Implementation) AddInstruction(instruction dynamic.Instruction) { 
	implementation.program.WriteTo(*implementation.active, instruction)
}

func (implementation Implementation) ReserveRegister() int {
	return implementation.program.ReserveRegister(*implementation.active)
}

func (implementation Implementation) RegisterOf(value language.Type) int {
	i, _ := strconv.Atoi(string(implementation.ExpressionOf(value)))
	return i
}

func (implementation Implementation) ExpressionOf(t language.Type) language.Statement {
	return reflect.ValueOf(t).Convert(reflect.TypeOf(language.NewType{})).Interface().(language.NewType).Expression
}
