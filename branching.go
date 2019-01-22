package script

import "github.com/qlova/script/language"

type ElseIfChain struct {
	q Script

	ElseIf []func()
	Else func()
}

func (chain *ElseIfChain) eif(condition Bool, block func()) *ElseIfChain {
	var q = chain.q
	chain.ElseIf = append(chain.ElseIf, func() {
		q.indent()
		q.write(q.lang.ElseIf(condition.LanguageType().(language.Bit)))
		q.depth++
		block()
		q.depth--
	})
	return chain
}

func (chain *ElseIfChain) e(block func()) *ElseIfChain {
	chain.Else = block
	return chain
}

func (q Script) ElseIf(condition Bool, block func()) *ElseIfChain {
	var chain = new(ElseIfChain)
		chain.q = q

	return chain.eif(condition, block)
}

func (q Script) Else(block func()) *ElseIfChain {
	var chain = new(ElseIfChain)
		chain.q = q

	return chain.e(block)
}

func (q Script) If(condition Bool, block func(), elseifchain ...*ElseIfChain) {
	q.indent()
	q.write(q.lang.If(condition.LanguageType().(language.Bit)))
	q.depth++
	block()
	q.depth--
	
	var chain *ElseIfChain
	
	if len(elseifchain) > 0 {
		chain = elseifchain[0]
	}
	
	if chain != nil {
		for _, Elseif := range chain.ElseIf {
			Elseif()
		}
		
		if chain.Else != nil {
			q.indent()
			q.write(q.lang.Else())
			q.depth++
			chain.Else()
			q.depth--
		}
	}

	q.indent()
	q.write(q.lang.EndIf())
}
