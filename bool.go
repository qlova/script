package script

type Bool struct {
	Type
}

//AnyBool is anything that can retrieve a script.Bool.
type AnyBool interface {
	BoolFromCtx(AnyCtx) Bool
}

//BoolFromCtx implements AnyBool.
func (b Bool) BoolFromCtx(AnyCtx) Bool {
	return b
}

func (q Ctx) Bool(literal bool) Bool {
	return Bool{q.Literal(literal)}
}

func (q Ctx) Not(b Bool) Bool {
	return q.Language.Not(b)
}
