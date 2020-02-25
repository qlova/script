package script

type If struct {
	Condition Bool
	Block     func()
}

func (q Ctx) If(condition Bool, f func()) *ElseIfChain {
	var chain ElseIfChain
	chain.q = q
	chain.If.Condition = condition
	chain.If.Block = f
	return &chain
}

func (q Ctx) IfL(condition bool, f func()) *ElseIfChain {
	return q.If(q.Bool(condition), f)
}

type ElseIfChain struct {
	q Ctx

	If
	Chain []If
}

type ElseIfEnder struct {
	ElseIfChain ElseIfChain
	Else        func()
}

func (ender ElseIfEnder) End() {
	var chain = ender.ElseIfChain
	var q = chain.q
	q.Language.If(chain.If, chain.Chain, ender.Else)
}

func (chain *ElseIfChain) ElseIf(condition Bool, f func()) *ElseIfChain {
	chain.Chain = append(chain.Chain, If{
		Condition: condition,
		Block:     f,
	})
	return chain
}

func (chain *ElseIfChain) ElseIfL(condition bool, f func()) *ElseIfChain {
	return chain.ElseIf(chain.q.Bool(condition), f)
}

func (chain *ElseIfChain) Else(f func()) ElseIfEnder {
	return ElseIfEnder{*chain, f}
}

func (chain *ElseIfChain) End() {

}
