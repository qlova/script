package example

import "github.com/qlova/script/language"

func (implementation Implementation) Buffer() language.Buffer {
	panic(implementation.Name() + ".Buffer() Unimplemented")
	return nil
}

func (implementation Implementation) Flush(buffer language.Buffer) {
	panic(implementation.Name() + ".Flush() Unimplemented")
}
