package Interpreter

import "github.com/qlova/script/language"
import "github.com/qlova/script/interpreter/internal"
import "os/exec"

const Name = "Interpreter"

type implementation struct {
	Blocks map[string]*internal.Block
	count int
	
	BlockPointer *internal.Block
	InstructionPointer int
	BreakPoint int //This is where we go when we hit a break.
	
	pushedBlocks []*internal.Block
	pushedAddresses [][]internal.Address
	pushedPointers []int
}


func (l *implementation) Name() string { return Name }

func New() *implementation {
	var i implementation 
	i.Init()
	i.Blocks["main"] = new(internal.Block)
	i.BlockPointer = i.Blocks["main"]
	return &i
}

func (c *implementation) loadBlock() *internal.Block {
	
	if c.BlockPointer == nil {
		panic("BlockPointer is nil!")
	}
	
	return c.BlockPointer
}

func (c *implementation) pushBlock(b *internal.Block) {
	c.pushedBlocks = append(c.pushedBlocks, b)
}

func (c *implementation) pushAddresses(a ...internal.Address) {
	c.pushedAddresses = append(c.pushedAddresses, a)
}

func (c *implementation) pushPointer() {
	c.pushedPointers = append(c.pushedPointers, c.InstructionPointer)
}


func (c *implementation) addAddress(a internal.Address) {
	c.pushedAddresses[len(c.pushedAddresses)-1] = append(c.pushedAddresses[len(c.pushedAddresses)-1], a)
}

func (b *implementation) popBlock() (p *internal.Block) {
	p = b.pushedBlocks[len(b.pushedBlocks)-1]
	b.pushedBlocks = b.pushedBlocks[:len(b.pushedBlocks)-1]
	return
}

func (b *implementation) popPointer() {
	b.InstructionPointer = b.pushedPointers[len(b.pushedPointers)-1]
	b.pushedPointers = b.pushedPointers[:len(b.pushedPointers)-1]
	
	b.InstructionPointer++
	return
}

func (b *implementation) popAddresses() (p []internal.Address) {
	p = b.pushedAddresses[len(b.pushedAddresses)-1]
	b.pushedAddresses = b.pushedAddresses[:len(b.pushedAddresses)-1]
	return
}


func (c *implementation) JumpTo(name string) {
	c.BlockPointer = c.Blocks[name]
	c.InstructionPointer = 0
}

func (c *implementation) JumpToBlock(target *internal.Block) {
	c.pushBlock(c.BlockPointer)
	c.pushPointer()
	
	c.InstructionPointer = -1
	c.BlockPointer = target
}


func (c *implementation) CreateBlock(name string) {
	c.Blocks[name] = &internal.Block{Variables: make(map[string]int)}
}

func (c *implementation) SetBlock(name string) {
	c.BlockPointer = c.Blocks[name]
}

func (c *implementation) Start() {
	
	if len(c.pushedBlocks) > 0 {
		panic("Houston, we have a problem")
	}
	
	c.JumpTo("main")
	
	for {
		c.BlockPointer.Instructions[c.InstructionPointer]()
		c.InstructionPointer++
		for c.InstructionPointer > len(c.BlockPointer.Instructions)-1 {
			if len(c.pushedBlocks) > 0 {
				//Jump back to parent block!
				c.BlockPointer = c.popBlock()
				c.popPointer()
				
			} else {
				return
			}
		}
	}
}

var _ = language.Interface(New())

func (c *implementation) Init() {
	c.Blocks = make(map[string]*internal.Block)
}
func (l *implementation) Head() language.Statement { return "" }
func (l *implementation) Neck() language.Statement { return "" }
func (l *implementation) Body() language.Statement { return "" }
func (l *implementation) Tail() language.Statement { return "" }
func (l *implementation) Last() language.Statement { return "" }

//Returns a Statement that begins the main entry point to the program.
func (l *implementation) Build(path string) *exec.Cmd {
	return nil
}

//Returns a Statement that begins the main entry point to the program.
func (l *implementation) Main() language.Statement {
	l.CreateBlock("main")
	l.SetBlock("main")
	return ""
}

//Returns a Statement that exits the program.
func (l *implementation) Exit() language.Statement {
	panic("Error in "+Name+".Main(): Unimplemented")
	return ""
}

//Returns a Statement that ends the main entry point to the program.
func (l *implementation) EndMain() language.Statement {
	l.SetBlock("")
	return ""
}
