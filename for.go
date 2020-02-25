package script

//For loop constants.
const (
	Equals         = "=="
	LessThan       = "<"
	MoreThan       = ">"
	LessThanEquals = "<="
	MoreThanEquals = ">="

	Plus1  = "++"
	Minus1 = "--"

	PlusEquals     = "+="
	MinusEquals    = "-="
	OverEquals     = "/="
	LeftOverEquals = "%="
	TimesEquals    = "*="
)

type ForLoopCondition struct {
	Operator string
	Subject  Int
}

func (q Ctx) LessThan(other Int) ForLoopCondition {
	return ForLoopCondition{
		Operator: LessThan,
		Subject:  other,
	}
}

//LessThanL is the literal version of LessThan.
func (q Ctx) LessThanL(other int) ForLoopCondition {
	return q.LessThan(q.Int(other))
}

func (q Ctx) MoreThan(other Int) ForLoopCondition {
	return ForLoopCondition{
		Operator: MoreThan,
		Subject:  other,
	}
}

func (q Ctx) MoreThanL(other int) ForLoopCondition {
	return q.MoreThan(q.Int(other))
}

func (q Ctx) Equals(other Int) ForLoopCondition {
	return ForLoopCondition{
		Operator: Equals,
		Subject:  other,
	}
}

func (q Ctx) EqualsL(other int) ForLoopCondition {
	return q.Equals(q.Int(other))
}

func (q Ctx) LessThanEquals(other Int) ForLoopCondition {
	return ForLoopCondition{
		Operator: LessThanEquals,
		Subject:  other,
	}
}

func (q Ctx) LessThanEqualsL(other int) ForLoopCondition {
	return q.LessThanEquals(q.Int(other))
}

func (q Ctx) MoreThanEquals(other Int) ForLoopCondition {
	return ForLoopCondition{
		Operator: MoreThanEquals,
		Subject:  other,
	}
}

func (q Ctx) MoreThanEqualsL(other int) ForLoopCondition {
	return q.MoreThanEquals(q.Int(other))
}

type ForLoopAction struct {
	Operator string
	Subject  Int
}

func (q Ctx) Plus1() ForLoopAction {
	return ForLoopAction{
		Operator: Plus1,
	}
}

func (q Ctx) Minus1() ForLoopAction {
	return ForLoopAction{
		Operator: Minus1,
	}
}

//PlusEquals forloop action +=
func (q Ctx) PlusEquals(other Int) ForLoopAction {
	return ForLoopAction{
		Operator: PlusEquals,
		Subject:  other,
	}
}

//PlusEqualsL is the literal version of PlusEquals.
func (q Ctx) PlusEqualsL(other int) ForLoopAction {
	return q.PlusEquals(q.Int(other))
}

//MinusEquals forloop action -=
func (q Ctx) MinusEquals(other Int) ForLoopAction {
	return ForLoopAction{
		Operator: MinusEquals,
		Subject:  other,
	}
}

//MinusEqualsL is the literal version of MinusEquals.
func (q Ctx) MinusEqualsL(other int) ForLoopAction {
	return q.MinusEquals(q.Int(other))
}

//TimesEquals forloop action *=
func (q Ctx) TimesEquals(other Int) ForLoopAction {
	return ForLoopAction{
		Operator: TimesEquals,
		Subject:  other,
	}
}

//TimesEqualsL is the literal version of TimesEquals.
func (q Ctx) TimesEqualsL(other int) ForLoopAction {
	return q.MinusEquals(q.Int(other))
}

//OverEquals forloop action /=
func (q Ctx) OverEquals(other Int) ForLoopAction {
	return ForLoopAction{
		Operator: OverEquals,
		Subject:  other,
	}
}

//OverEqualsL is the literal version of OverEquals.
func (q Ctx) OverEqualsL(other int) ForLoopAction {
	return q.OverEquals(q.Int(other))
}

//LeftOverEquals forloop action %=
func (q Ctx) LeftOverEquals(other Int) ForLoopAction {
	return ForLoopAction{
		Operator: LeftOverEquals,
		Subject:  other,
	}
}

//LeftOverEqualsL is the literal version of LeftOverEquals.
func (q Ctx) LeftOverEqualsL(other int) ForLoopAction {
	return q.LeftOverEquals(q.Int(other))
}

//For is a traditional for loop.
func (q Ctx) For(set Int, condition ForLoopCondition, action ForLoopAction, f func(Int)) {
	q.Language.For(set, condition, action, f)
}

//ForL is the literal version of For.
func (q Ctx) ForL(set int, condition ForLoopCondition, action ForLoopAction, f func(Int)) {
	q.For(q.Int(set), condition, action, f)
}
