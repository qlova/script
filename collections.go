package script

func (q Script) Len(t Type) Int {
	return q.IntFromLanguageType(q.lang.Length(t.LanguageType()))
}
