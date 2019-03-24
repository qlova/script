package interpreter

import "strconv"
import "reflect"

import "github.com/qlova/script/language"
import "github.com/qlova/script/interpreter/dynamic"

type implementation struct {
	symbols map[string]dynamic.BlockPointer
	
	program dynamic.Program

	//Trackers for program blocks.
	active dynamic.BlockPointer 		//This is the code block we are writing to.
	inactive []dynamic.BlockPointer	   //These are the code blocks we were writing to.
	
	//In order to support out-of order function definitions, we need to use Buffers.
	buffers []Buffer
	
	//Stores instruction-count positions of loops.
	loops []int
}

type Implementation struct {
	*implementation
}

func (implementation Implementation) Start() {
	implementation.program.Run()
}

func New() Implementation {
	var implementation implementation
	implementation.program = make(dynamic.Program, 0)

	implementation.inactive = make([]dynamic.BlockPointer, 0)
	implementation.buffers = make([]Buffer, 0)
	
	return Implementation{&implementation}
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

func (implementation Implementation) Active() *dynamic.Block {
	if len(implementation.buffers) > 0 {
		var buffer = implementation.buffers[len(implementation.buffers)-1]
		
		if buffer.sister == implementation.active {
			return implementation.buffers[len(implementation.buffers)-1].block
		}
	}
	return &implementation.program[implementation.active]
}

//Return the current instruction index.
func (implementation Implementation) Instruction() int {
	var active = implementation.Active()
	return len(active.Instructions)
}

func (implementation Implementation) SetInstruction(index int, instruction dynamic.Instruction) {
	var active = implementation.Active()
	active.Instructions[index] = instruction
}


func (implementation Implementation) AddInstruction(instruction dynamic.Instruction) {
	var active = implementation.Active()
	active.Instructions = append(active.Instructions, instruction)

}

func (implementation Implementation) ReserveRegister() int {
	var block = implementation.Active()
	block.Registers++
	return block.Registers-1
}

func (implementation Implementation) RegisterOf(value language.Type) int {
	var expression = string(implementation.ExpressionOf(value))
	
	i, err := strconv.Atoi(expression)
	if err != nil && implementation.ExpressionOf(value)[0] == '$' {

		//This is a function variable.
		i, _ = strconv.Atoi(expression[1:])
		
	} else if err != nil {
		
		//This must be an argument, has it been defined yet?
		var arguments = implementation.Active().Arguments
		
		if table, ok := arguments[expression]; ok {
			return table[1]
		} else {
			//The argument register is not defined yet, so we will create a mapping for it.
			arguments[expression] = [2]int{-1, len(arguments)}
			return len(arguments)-1
		}
	}
	return i
}

func (implementation Implementation) BlockOf(value language.Type) dynamic.BlockPointer {
	i, _ := strconv.Atoi(string(implementation.ExpressionOf(value)))
	return dynamic.BlockPointer(i)
}

func (implementation Implementation) ExpressionOf(t language.Type) language.Statement {
	return reflect.ValueOf(t).Convert(reflect.TypeOf(language.NewType{})).Interface().(language.NewType).Expression
}

func (implementation Implementation) CreateBlock() dynamic.BlockPointer {
	var pointer = implementation.program.CreateBlock()
	implementation.active = pointer
	return pointer
}

func (implementation Implementation) Activate(pointer dynamic.BlockPointer) {
	implementation.inactive = append(implementation.inactive, implementation.active)
	implementation.active = pointer
}

func (implementation Implementation) Deactivate() {
	implementation.active = implementation.inactive[len(implementation.inactive)-1]
	implementation.inactive = implementation.inactive[:len(implementation.inactive)-1]
}

func GoTypeOf(t interface{}) reflect.Type {
	switch t.(type) {
		case String:
			 return reflect.TypeOf("")
		case Integer:
			 return reflect.TypeOf(0)
		default:
			panic("Unimplemented")
	}
	return nil
}
