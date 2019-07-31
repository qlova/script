package script

func (q *Script) foreach(collection Type, f func(), names ...string) {

	var i, v string

	if len(names) > 0 {
		i = names[0]
	} else {
		i = Unique()
	}

	if len(names) > 1 {
		v = names[1]
	} else {
		v = Unique()
	}

	q.indent()
	statement, index, value := q.lang.ForEach(i, v, collection.LanguageType())

	q.context.index = q.IntFromLanguageType(index)
	q.context.value = q.ValueFromLanguageType(value)

	q.write(statement)
	q.depth++
	f()
	q.depth--
	q.indent()
	q.write(q.lang.EndForEach())
}

func (q *Script) Index() Int {
	return q.context.index
}

func (q *Script) Value() Value {
	return q.context.value
}
