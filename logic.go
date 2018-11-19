package script

import "fmt"
import "github.com/qlova/script/language"

type Boolean struct {
	language.Boolean
	EmbeddedScript

	Literal *bool
}

func (Boolean) SameAs(i interface{}) bool { _, ok := i.(Boolean); return ok }

//Converts a Go bool to a language.Boolean.
func (q Script) Boolean(b ...bool) Boolean {
	boolean := false
	if len(b) > 0 {
		boolean = b[0]
	}	
	return Boolean{Literal: &boolean, EmbeddedScript: EmbeddedScript{ q: q }}
}

func (q Script) Equals(a, b Number) Boolean {
	return q.wrap(q.lang.Equals(convert(a), convert(b))).(Boolean)
}

func (q Script) And(a, b Boolean) Boolean {
	return q.wrap(q.lang.And(convert(a).(language.Boolean), convert(b).(language.Boolean))).(Boolean)
}

func (q Script) Or(a, b Boolean) Boolean {
	return q.wrap(q.lang.Or(convert(a).(language.Boolean), convert(b).(language.Boolean))).(Boolean)
}

func (q Script) Not(a Boolean) Boolean {
	return q.wrap(q.lang.Not(convert(a).(language.Boolean))).(Boolean)
}

func (q Script) Xor(a, b Boolean) Boolean {
	return q.And(q.Or(a, b), q.Not(q.And(a, b)))
}


//Magic to make If statements both type and format safe.
func (q Script) If(condition Boolean, block func(Script), ifelsechain ...IfElseChain) {
	
	var chain Chain
	
	if len(ifelsechain) > 0 {
		if elseifchain, ok := ifelsechain[0].(*ElseIfChain); ok {
			chain = elseifchain.Chain
		} else if endchain, ok := ifelsechain[0].(*EndChain); ok {
			chain = endchain.ElseIfChain.Chain
		} else {
			panic("Invalid If statement! "+fmt.Sprint(ifelsechain[0]))
		}
	}

	q.indent()
	q.write(q.lang.If(convert(condition).(language.Boolean)))
	q.depth++
	block(q)
	q.depth--
	for _, elseif := range chain.body {
		elseif(q)
	}
	
	if chain.tail != nil {
		
		q.indent()
		q.write(q.lang.Else())
		q.depth++
		chain.tail(q)
		q.depth--
	}

	q.indent()
	q.write(q.lang.EndIf())
}

type Chain struct {
	head func(Script)
	neck []func(Script)
	body []func(Script)
	tail func(Script)
}

type IfElseChain interface {
	ifElseChain()
}

type ElseIfChain struct {
	Chain
}

func (ElseIfChain) ifElseChain() {}

type EndChain struct {
	ElseIfChain
}

func (EndChain) ifElseChain() {}

func (q Script) ElseIf(condition Boolean, block func(Script)) *ElseIfChain {
	return new(ElseIfChain).ElseIf(condition, block)
}

func (chain *ElseIfChain) ElseIf(condition Boolean, block func(Script)) *ElseIfChain {
	chain.body = append(chain.body, func(q Script) {
		q.indent()
		q.write(q.lang.ElseIf(convert(condition).(language.Boolean)))
		q.depth++
		block(q)
		q.depth--
	})
	
	return chain
}


func (q Script) Else(block func(Script)) *EndChain {
	return new(ElseIfChain).Else(block)
}

func (chain *ElseIfChain) Else(block func(Script)) *EndChain {
	chain.tail = block
	
	return &EndChain{ElseIfChain: *chain}
}
