package Interpreter

import "github.com/qlova/script/language"
import "github.com/qlova/script/interpreter/internal"

const Name = "Interpreter"

type implementation struct {
	Blocks map[string]*internal.Block
	
	BlockPointer *internal.Block
	InstructionPointer int
	BreakPoint int //This is where we go when we hit a break.
	
	currentBlock string //Used when compiling.
}

func New() *implementation {
	var i implementation 
	i.Init()
	return &i
}

func (c *implementation) loadBlock() *internal.Block {
	return c.Blocks[c.currentBlock]
}

func (c *implementation) JumpTo(name string) {
	c.BlockPointer = c.Blocks[name]
}

func (c *implementation) CreateBlock(name string) {
	c.Blocks[name] = &internal.Block{Variables: make(map[string]int)}
}

func (c *implementation) SetBlock(name string) {
	c.currentBlock = name
}

func (c *implementation) Start() {
	
	c.JumpTo("main")
	
	for {
		c.BlockPointer.Instructions[c.InstructionPointer]()
		c.InstructionPointer++
		if c.InstructionPointer > len(c.BlockPointer.Instructions)-1 {
			return
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
