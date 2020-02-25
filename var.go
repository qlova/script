package script

//Literal returns the script value of the given Go value.
func (q Ctx) Literal(value interface{}) Type {
	if q.defining {

		var variable = q.getVar()

		if q.Language != nil {
			return Type{
				Ctx: q,
				Runtime: q.Language.DefineVariable(variable, NewType(q, func() interface{} {
					return value
				})),
			}
		}
		panic("unimplemented")
	}

	return NewType(q, func() interface{} {
		return value
	})
}

//Var returns a variable with optional name.
func (q Ctx) Var(name ...string) Ctx {
	q.defining = true
	if len(name) > 0 {
		q.variables = append(q.variables, name[0])
	} else {
		q.variables = append(q.variables, q.ID("var"))
	}
	return q
}

func (q Ctx) getVar() (variable string) {
	variable = q.variables[len(q.variables)-1]
	q.variables = q.variables[:len(q.variables)-1]

	if len(q.variables) == 0 {
		defer func() {
			q.defining = false
			q.variables = nil
		}()
	}
	return
}

//Set sets a value 'a' to be set to value 'b'.
func (q Ctx) Set(a Value, b Value) {
	if q.Language != nil {
		q.Language.Set(a, b)
		return
	}
}
