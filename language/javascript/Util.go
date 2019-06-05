package Javascript

import "github.com/qlova/script/language"

func (implementation Implementation) Copy(t language.Type) language.Type {
	panic(implementation.Name() + ".Copy() Unimplemented")
	return nil
}
