package runtime

import "github.com/qlova/script"

//Value is a helper function for creating script.Type(s)
func Value(q script.Ctx, f func() interface{}) script.Type {
	return script.Type{
		Ctx:     q,
		Runtime: &f,
	}
}
