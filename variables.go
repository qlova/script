package script

func (q Script) Var(T Type, name ...string) Value {
	var register = unique(name)
	
	q.indent()
	statement, variable := q.lang.Register(register, T.LanguageType())
	q.write(statement)
	
	return Value{
		script: q,
		internal: variable,
	}
}
