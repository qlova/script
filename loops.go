 package script
 
 import "github.com/qlova/script/language"
 
 func (q *Script) ForRange(i string, from language.Number, to language.Number, block func(language.Number, *Script)) {
	variable, statement := q.lang.ForRange(i, from, to)
	
	q.indent()
	q.write(statement)
	q.depth++
		block(variable, q)
	q.depth--
	q.indent()
	q.write(q.lang.EndForRange())
 }
