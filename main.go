package script

//Main is the entrypoint to the script.
func (q Ctx) Main(f func()) {
	q.Language.Main(f)
}
