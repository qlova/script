package internal

import "math/big"

type Block struct {
	//Local Type registers.
	Numbers []*big.Int
	Strings []string
	
	//Instructions in this block.
	Instructions []Instruction
	
	//Table that maps a variable name to a type register.
	Variables map[string]int
	
	//Which parent is our block? We can access their variables.
	Parent *Block
	
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


func (b *Block) GetVariable(name string) (int, *Block) {
	variable, ok := b.Variables[name]
	if !ok {
		
		if b.Parent != nil {
			return b.Parent.GetVariable(name)
		}
		
		panic(name+" is undefined!")
	}
	return variable, b
}

func (b *Block) DeclareVariable(name string, address int) {
	b.Variables[name] = address
}

func (b *Block) AddInstruction(instruction func()) {
	b.Instructions = append(b.Instructions, instruction)
}

func (b *Block) CreateNumber() int {
	b.Numbers = append(b.Numbers, big.NewInt(0))
	return len(b.Numbers)-1
}

func (b *Block) SetNumber(address int, value *big.Int)  {
	b.Numbers[address] = value
}

func (b *Block) GetNumber(address int) *big.Int {
	return b.Numbers[address]
}

func (b *Block) CreateString() int {
	b.Strings = append(b.Strings, "")
	return len(b.Strings)-1
}

func (b *Block) SetString(address int, value string)  {
	b.Strings[address] = value
}

func (b *Block) GetString(address int) string {
	return b.Strings[address]
}
 
