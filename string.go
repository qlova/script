package script

type String struct {
	Type
}

func (q Ctx) String(literal string) String {
	return String{q.Literal(literal)}
}
