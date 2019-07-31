package Javascript

import "github.com/qlova/script/language"

func (implementation Implementation) Cast(a, b language.Type) language.Type {

	switch a.(type) {
	case Integer:
		switch b.(type) {
		case String:
			return String{Expression: "(" + implementation.ExpressionOf(a) + ").toString()"}
		}
	}

	panic(implementation.Name() + ".Cast() Unimplemented")
	return nil
}
