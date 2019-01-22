package script

func (q Script) Trace(line int, file string) {
	q.write(q.lang.Trace(line, file))
}
