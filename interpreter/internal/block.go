package internal

import "math/big"
import "github.com/qlova/script/language"

type Block struct {
	//Local Type registers.
	Numbers []*big.Int
	Strings []string
	Booleans []bool
	Symbols []rune
	Functions []*Block

	Pointers []int
	
	//Instructions in this block.
	Instructions []Instruction
	
	//Table that maps a variable name to a type register.
	Variables map[string]int
	
	//Pointers to local Instruction space.
	pointers []int
}

func (b *Block) PushPointer() {
	b.AddInstruction(nil)
	b.pointers = append(b.pointers, len(b.Instructions)-1)
}

func (b *Block) PopPointer() (p int) {
	p = b.pointers[len(b.pointers)-1]
	b.pointers = b.pointers[:len(b.pointers)-1]
	return
}


func (b *Block) DeclareVariable(name string, address int) {
	b.Variables[name] = address
}

func (b *Block) AddInstruction(instruction func()) {
	b.Instructions = append(b.Instructions, instruction)
}

type NumberAddress int
func (address NumberAddress) Address() int { return int(address) }

	func (b *Block) CreateNumber() NumberAddress {
		b.Numbers = append(b.Numbers, big.NewInt(0))
		return NumberAddress(len(b.Numbers)-1)
	}

	func (b *Block) SetNumber(address NumberAddress, value *big.Int)  {
		b.Numbers[address] = value
	}

	func (b *Block) GetNumber(address NumberAddress) *big.Int {
		return b.Numbers[address]
	}

type StringAddress int
func (address StringAddress) Address() int { return int(address) }
	
	func (b *Block) CreateString() StringAddress {
		b.Strings = append(b.Strings, "")
		return StringAddress(len(b.Strings)-1)
	}

	func (b *Block) SetString(address StringAddress, value string)  {
		b.Strings[address] = value
	}

	func (b *Block) GetString(address StringAddress) string {
		return b.Strings[address]
	}
 
 
type BooleanAddress int
func (address BooleanAddress) Address() int { return int(address) }

	//Booleans
	func (b *Block) CreateBoolean() BooleanAddress {
		b.Booleans = append(b.Booleans, false)
		return BooleanAddress(len(b.Booleans)-1)
	}

	func (b *Block) SetBoolean(address BooleanAddress, value bool)  {
		b.Booleans[address] = value
	}

	func (b *Block) GetBoolean(address BooleanAddress) bool {
		return b.Booleans[address]
	}

type SymbolAddress int
func (address SymbolAddress) Address() int { return int(address) }

	//Symbols
	func (b *Block) CreateSymbol() SymbolAddress {
		b.Symbols = append(b.Symbols, 0)
		return SymbolAddress(len(b.Symbols)-1)
	}

	func (b *Block) SetSymbol(address SymbolAddress, value rune)  {
		b.Symbols[address] = value
	}

	func (b *Block) GetSymbol(address SymbolAddress) rune {
		return b.Symbols[address]
	}

type FunctionAddress int
func (address FunctionAddress) Address() int { return int(address) }
	
	func (b *Block) CreateFunction() FunctionAddress {
		b.Functions = append(b.Functions, nil)
		return FunctionAddress(len(b.Functions)-1)
	}

	func (b *Block) SetFunction(address FunctionAddress, value *Block)  {
		b.Functions[address] = value
	}

	func (b *Block) GetFunction(address FunctionAddress) *Block {
		return b.Functions[address]
	}
	
type ArrayAddress int
func (address ArrayAddress) Address() int { return int(address) }
		
	//Arrays //TODO put arrays in their own block so they don't clash.
	func (b *Block) CreateArray(T language.Array) ArrayAddress {
		
		var Address Address
		var Length = T.Length()
		switch T.SubType().(type) {
			case language.Number:
				Address = NumberAddress(len(b.Numbers))
				b.Numbers = append(b.Numbers, make([]*big.Int, Length)...)
				b.AddInstruction(func() {
					var z big.Int
					for i := Address.Address(); i < Length; i++ {
						b.SetNumber(NumberAddress(Address.Address() + i), &z)
					}
				})
				
			default:
				panic("Error in Block.CreateArray("+T.Name()+"): Unsupported subtype for Array: "+T.SubType().Name())
		}
		
		b.Pointers = append(b.Pointers, 0)
		return ArrayAddress(len(b.Pointers)-1)
	}

	func (b *Block) SetArray(address ArrayAddress, value int)  {
		b.Pointers[address] = value
	}

	func (b *Block) GetArray(address ArrayAddress) int {
		return b.Pointers[address]
	}
