package compiler

type Shunt func(*Compiler, string, Type, Type) Type

func (c *Compiler) Shunt(t Type, precedence int) Type {

shunting:
	for peek := c.Peek(); c.GetOperator(peek).Precedence >= precedence; {

		if c.GetOperator(c.Peek()).Precedence == -1 {
			break
		}
		op := c.GetOperator(peek)
		c.Scan()

		rhs := c.Expression()
		peek = c.Peek()
		for c.GetOperator(peek).Precedence > op.Precedence {
			rhs = c.Shunt(rhs, c.GetOperator(peek).Precedence)
			peek = c.Peek()
		}

		for i := range c.Shunts {
			if result := c.Shunts[i](c, op.Symbol, t, rhs); result != nil {
				t = result
				continue shunting
			}
		}

		c.RaiseError(Translatable{
			English: "Operator " + op.Symbol + " does not apply to " + t.LanguageType().Name(),
		})
	}
	return t
}
