package dynamic

type BlockPointer int 

type Thread struct {
	Program

	InstructionCounter int
	Block, Caller BlockPointer
	Registers [][]interface{}
}

func (thread Thread) Set(location int, value interface{}) {
	thread.Registers[len(thread.Registers)-1][location] = value
}

func (thread Thread) Get(location int) interface{} {
	return thread.Registers[len(thread.Registers)-1][location]
}

type Instruction func(thread *Thread)
type Block struct {
	Instructions []Instruction
	
	Registers int
}
type Program []Block

func (program *Program) CreateBlock() BlockPointer {
	*program = append(*program, Block{})
	return BlockPointer(len(*program)-1)
}

func (program *Program) WriteTo(block BlockPointer, instruction Instruction) {
	(*program)[block].Instructions = append((*program)[block].Instructions, instruction)
}

func (program *Program) ReserveRegister(block BlockPointer) int {
	(*program)[block].Registers++
	return (*program)[block].Registers-1
}

func (program Program) Run() {
	var thread Thread
	thread.Program = program
	thread.Block = BlockPointer(len(program)-1)
	thread.Caller = -1
	thread.Registers = append(thread.Registers, make([]interface{}, program[thread.Block].Registers))
	
	for {
		program[thread.Block].Instructions[thread.InstructionCounter](&thread)
		thread.InstructionCounter++
		if thread.InstructionCounter > len(program[thread.Block].Instructions)-1 {
			if thread.Caller == -1 {
				break
			}
			thread.Block = thread.Caller
			thread.Registers = thread.Registers[:len(thread.Registers)-1]
		}
	}
}
