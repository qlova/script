package script

func (q Script) Var(T Type, name ...string) Value {
	var register string
	if len(name) > 0 {
		register = name[0]
	} else {
		register = Unique()
	}
	
	q.indent()
	statement, variable := q.lang.Register(register, T.LanguageType())
	q.write(statement)
	
	return Value{
		script: q,
		internal: variable,
	}
}
