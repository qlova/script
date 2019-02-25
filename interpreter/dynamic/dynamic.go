package dynamic

import "strconv"
import "fmt"

const Debug = false

type BlockPointer int 
func (pointer BlockPointer) String() string {
	return strconv.Itoa(int(pointer))
}

//A caller keeps track of the history of the calling frame, where to jump back to etc.
type Caller struct {
	Block BlockPointer
	InstructionCounter int
}

//A thread is capable of running a program.
type Thread struct {
	Program

	InstructionCounter int
	Block BlockPointer
	Registers [][]interface{}
	
	Callers []Caller
}

//Write a value in the thread at the specified location.
func (thread Thread) Set(location int, value interface{}) {
	thread.Registers[len(thread.Registers)-1][location] = value
}

//Read a value from the thread at the specified location.
func (thread Thread) Get(location int) interface{} {
	return thread.Registers[len(thread.Registers)-1][location]
}

//Read a value from the thread at the specified location.
func (thread *Thread) JumpTo(block BlockPointer, arguments ...interface{}) {
	
	//Update callers.
	thread.Callers = append(thread.Callers, Caller{
		Block: thread.Block,
		InstructionCounter: thread.InstructionCounter,
	})
	
	//Initialise registers.
	thread.Registers = append(thread.Registers, make([]interface{}, thread.Program[block].Registers))
	
	//Fill registers with arguments.
	for i := range arguments {
		thread.Set(i, arguments[i])
	}
	
	//Jump.
	thread.Block = block
	thread.InstructionCounter = -1
	
	if Debug {
		println("jumping to a new block (", block,") with")
		println(len(thread.Program[thread.Block].Instructions))
		println("instructions &")
		println(thread.Program[thread.Block].Registers)
		println("registers")
	}
}	


//Return to the caller.
func (thread *Thread) Return() {

	//Return to the caller.
	var Caller = thread.Callers[len(thread.Callers)-1]
	thread.Block = Caller.Block
	thread.InstructionCounter = Caller.InstructionCounter+1
	thread.Callers = thread.Callers[:len(thread.Callers)-1]

	thread.Registers = thread.Registers[:len(thread.Registers)-1]
	
	if Debug {
		println("jumping back to block (", Caller.Block,") with")
		println(len(thread.Program[thread.Block].Instructions))
		println("instructions")
	}
}

//An instruction is a function that takes a thread as an argument.
type Instruction func(thread *Thread)

//A block consists of a number of 'instructions' and a number of registers.
type Block struct {
	Instructions []Instruction
	
	Main bool
	Registers int
}

//A program contains a number of blocks.
type Program []Block

func (program Program) Dump() {
	fmt.Println("Program: ", len(program), " blocks")
	fmt.Println("{")
	for i, block := range program {
		if block.Main {
			fmt.Print("\tmain\t")
		} else {
			fmt.Print("\t\t")
		}
		fmt.Println("block - ", i, " has ", len(block.Instructions), " instructions & ", block.Registers, "registers")
	}
	fmt.Println("}")
}

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
	//Create a new thread to run this program.
	var thread Thread
	thread.Program = program
	
	//Find main.
	for pointer, block := range thread.Program {
		if block.Main {
			thread.Block = BlockPointer(pointer)
		}
	}
	
	//Initisalise the registers.
	thread.Registers = append(thread.Registers, make([]interface{}, program[thread.Block].Registers))
	
	for {
		if int(thread.Block) >= len(program) {
			if Debug {
				println("runtime error: invalid block")
				println(thread.Block)
			}
		}
		
		if thread.InstructionCounter >= len(program[thread.Block].Instructions) {
			if Debug {
				println("runtime error: invalid instructions")
				println(thread.InstructionCounter)
				println(thread.Block)
				println(len(program[thread.Block].Instructions))
				fmt.Println("DUMPING STATE")
				program.Dump()
			}
		}
		
		if Debug {
			println("running instruction ", thread.InstructionCounter, " of block ", thread.Block)
		}
		
		//Process the next instruction.
		program[thread.Block].Instructions[thread.InstructionCounter](&thread)
		thread.InstructionCounter++

		//Have we reached the end of the instructions?
		for thread.InstructionCounter > len(program[thread.Block].Instructions)-1 {
			//If we don't have a caller, the thread is finished.
			if len(thread.Callers) == 0 {
				return
			}
			
			thread.Return()
		}
	}
}
