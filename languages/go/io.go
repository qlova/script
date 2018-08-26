package Go

import qlova "github.com/qlova/script"

func (l *language) Print(q *qlova.Script, value qlova.String) {		
	l.Import(q, "os")

	q.Body.WriteString("os.Stdout.Write([]byte(")
	q.Body.WriteString(value.String())
	q.Body.WriteString("))\n")
}
