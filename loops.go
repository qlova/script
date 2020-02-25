package script

func (q Ctx) While(condition Bool, f func()) {
	q.Language.While(condition, f)
}

func (q Ctx) WhileL(condition bool, f func()) {
	q.While(q.Bool(condition), f)
}

func (q Ctx) Loop(f func()) {
	q.Language.Loop(f)
}

func (q Ctx) Break() {
	q.Language.Break()
}
