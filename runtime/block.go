package runtime

//Block is a code block within a script.
type Block struct {
	runtime *Runtime
	Parent  *Block

	Statements []func()

	Variables map[string]interface{}
	Args      map[int]interface{}
}

//NewBlock returns a new block.
func (runtime *Runtime) NewBlock() *Block {
	var block = new(Block)
	block.runtime = runtime
	block.Parent = runtime.Current
	return block
}

func (block *Block) get(variable string) interface{} {
	if value, ok := block.Variables[variable]; ok {
		return value
	}
	/*if value, ok := block.Args[variable]; ok {
		return value
	}*/
	if block.Parent == nil {
		return nil
	}
	return block.Parent.get(variable)
}

func (block *Block) define(variable string, value interface{}) {
	block.Variables[variable] = value
}

func (block *Block) set(variable string, value interface{}) {
	if _, ok := block.Variables[variable]; ok {
		block.Variables[variable] = value
		return
	}
	/*if _, ok := block.Args[variable]; ok {
		block.Args[variable] = value
		return
	}*/
	if block.Parent == nil {
		return
	}
	block.Parent.set(variable, value)
}

//Jump jumps to a block.
func (block *Block) Jump() {
	var runtime = block.runtime

	block.Variables = make(map[string]interface{})

	var caller = runtime.Current
	runtime.Current = block
	for _, statement := range block.Statements {
		if runtime.broken || runtime.returning != nil {
			break
		}
		statement()
	}

	runtime.Current = caller
}

//WriteStatement writes a statement to a block.
func (block *Block) WriteStatement(statement func()) {
	block.Statements = append(block.Statements, statement)
}

//WriteStatement writes a statement to a block.
func (runtime *Runtime) WriteStatement(statement func()) {
	var block = runtime.Current
	block.Statements = append(block.Statements, statement)
}
