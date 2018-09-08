 package script
 
 import "github.com/qlova/script/language"
 
 func (q *Script) ForRange(i string, from Number, to Number, block func(Number, *Script)) {
	variable, statement := q.lang.ForRange(i, convert(from).(language.Number), convert(to).(language.Number))
	
	q.indent()
	q.write(statement)
	q.depth++
		block(q.wrap(variable).(Number), q)
	q.depth--
	q.indent()
	q.write(q.lang.EndForRange())
 }
