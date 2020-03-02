package script

//String is a unicode string.
type String struct {
	Type
}

func (a String) Join(b String) String {
	return a.Ctx.Language.Join(a, b)
}

//StringFromCtx implements AnyString.
func (a String) StringFromCtx(AnyCtx) String {
	return a
}

//AnyString needs to return a String from this package.
type AnyString interface {
	StringFromCtx(AnyCtx) String
}

func (q Ctx) String(literal ...string) String {
	if len(literal) > 0 {
		return String{q.Literal(literal[0])}
	}
	return String{q.Literal("")}
}
