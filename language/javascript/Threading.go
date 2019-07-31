package Javascript

import "github.com/qlova/script/language"

func (implementation Implementation) Thread(name string, distance int, arguments []language.Type) language.Stream {
	panic(implementation.Name() + ".Thread() Unimplemented")
	return nil
}
