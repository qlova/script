package script

type Type struct {
	Ctx     Ctx
	Runtime *func() interface{}
}

type EmptyType interface {
	Value
	typePointer()
}

func (T *Type) typePointer() {}

func NewType(q Ctx, runtime func() interface{}) Type {
	return Type{
		Ctx:     q,
		Runtime: &runtime,
	}
}

func (T Type) Get() interface{} {
	return (*T.Runtime)()
}

func (T Type) T() Type {
	return T
}

func (T Type) Var(name ...string) {
	var q = T.Ctx

	q.Var(name...)
	var old = *T.Runtime
	var variable = q.getVar()
	*T.Runtime = *q.Language.DefineVariable(variable, NewType(T.Ctx, func() interface{} {
		return old()
	}))
}
