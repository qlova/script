package compiler

type Expression struct {
	Name   Translatable
	OnScan func(*Compiler) Type

	Detect func(*Compiler) Type
}

func (c *Compiler) Expression() Type {
	var token = c.Scan()
	for token == "\n" {
		token = c.Scan()
	}

	for _, expression := range c.Expressions {
		if expression.Name[c.Language] == token {
			return expression.OnScan(c)
		}
	}

	for _, expression := range c.Expressions {
		if expression.Detect != nil {
			if t := expression.Detect(c); t != nil {

				if variable, ok := t.(Variable); ok {
					if variable.Type == nil {
						continue
					} else {
						return variable.Type
					}
				}

				return t
			}
		}
	}

	c.RaiseError(Translatable{
		English: "Unknown Expression: " + token,
	})

	return nil
}

func (c *Compiler) ScanExpression() Type {
	var result = c.Shunt(c.Expression(), 0)

	if result == nil {
		c.RaiseError(Translatable{
			English: "Invalid Expression ",
		})
	}

	return result
}
