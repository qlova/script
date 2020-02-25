package script

type Bool struct {
	Type
}

func (q Ctx) Bool(literal bool) Bool {
	return Bool{q.Literal(literal)}
}
