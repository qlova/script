package interpreter

import "github.com/qlova/script/language"
import "github.com/qlova/script/interpreter/dynamic"

type Buffer struct {
	sister dynamic.BlockPointer
	block  *dynamic.Block
}

func (Buffer) Buffer() {}

func (implementation Implementation) Buffer() language.Buffer {
	var buffer Buffer
	buffer.sister = implementation.active
	buffer.block = new(dynamic.Block)
	buffer.block.ArgumentMapping = make(map[string]int)
	buffer.block.ArgumentTransform = make(map[int]int)

	implementation.buffers = append(implementation.buffers, buffer)

	return buffer
}

func (implementation Implementation) Flush(b language.Buffer) {
	var buffer = b.(Buffer)
	implementation.buffers = implementation.buffers[:len(implementation.buffers)-1]

	for _, instruction := range buffer.block.Instructions {
		implementation.AddInstruction(instruction)
	}

	//This is the out of-order function definition argument access clean-up.

	var CurrentMapping = implementation.Active().ArgumentMapping
	var CurrentTransform = implementation.Active().ArgumentTransform

	var NewMapping = buffer.block.ArgumentMapping

	for register := range CurrentMapping {
		CurrentTransform[CurrentMapping[register]] = NewMapping[register]
	}

	for i := 0; i < buffer.block.Registers; i++ {
		implementation.ReserveRegister()
	}
}
