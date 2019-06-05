package Go

import "github.com/qlova/script/language"

func (implementation Implementation) Copy(t language.Type) language.Type {

	if _, ok := t.(List); !ok {
		panic("Unimplemented: Copy(): non-list")
	}

	return t.Register("append(" + t.Name() + "{}, " + string(t.Raw()) + "...)")
}
