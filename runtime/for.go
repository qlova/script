package runtime

import (
	"github.com/qlova/script"
)

//For implements script.Language.For
func (runtime *Runtime) For(set script.Int,
	condition script.ForLoopCondition,
	action script.ForLoopAction,
	f func(script.Int)) {

	var backup = runtime.Current
	var block = runtime.NewBlock()
	runtime.Current = block

	f(script.Int{Type: Value(set.Ctx, func() interface{} {
		return block.get("i")
	})})

	runtime.Current = backup

	var a, b = *set.T().Runtime,
		*condition.Subject.T().Runtime

	var c func() interface{}
	if action.Subject.T().Runtime != nil {
		c = *action.Subject.T().Runtime
	}

	switch condition.Operator {
	case script.LessThan:
		switch action.Operator {
		case script.Plus1:
			runtime.WriteStatement(func() {
				for i := a().(int); !runtime.broken && i < b().(int); i++ {
					runtime.Current.define("i", i)
					block.Jump()
				}
				runtime.broken = false
			})
		case script.Minus1:
			runtime.WriteStatement(func() {
				for i := a().(int); !runtime.broken && i < b().(int); i-- {
					runtime.Current.define("i", i)
					block.Jump()
				}
				runtime.broken = false
			})
		case script.PlusEquals:
			runtime.WriteStatement(func() {
				for i := a().(int); !runtime.broken && i < b().(int); i += c().(int) {
					runtime.Current.define("i", i)
					block.Jump()
				}
				runtime.broken = false
			})
		case script.MinusEquals:
			runtime.WriteStatement(func() {
				for i := a().(int); !runtime.broken && i < b().(int); i -= c().(int) {
					runtime.Current.define("i", i)
					block.Jump()
				}
				runtime.broken = false
			})
		case script.OverEquals:
			runtime.WriteStatement(func() {
				for i := a().(int); !runtime.broken && i < b().(int); i /= c().(int) {
					runtime.Current.define("i", i)
					block.Jump()
				}
				runtime.broken = false
			})
		case script.LeftOverEquals:
			runtime.WriteStatement(func() {
				for i := a().(int); !runtime.broken && i < b().(int); i %= c().(int) {
					runtime.Current.define("i", i)
					block.Jump()
				}
				runtime.broken = false
			})
		case script.TimesEquals:
			runtime.WriteStatement(func() {
				for i := a().(int); !runtime.broken && i < b().(int); i *= c().(int) {
					runtime.Current.define("i", i)
					block.Jump()
				}
				runtime.broken = false
			})
		}
	case script.MoreThan:
		switch action.Operator {
		case script.Plus1:
			runtime.WriteStatement(func() {
				for i := a().(int); !runtime.broken && i > b().(int); i++ {
					runtime.Current.define("i", i)
					block.Jump()
				}
				runtime.broken = false
			})
		case script.Minus1:
			runtime.WriteStatement(func() {

				for i := a().(int); !runtime.broken && i > b().(int); i-- {
					runtime.Current.define("i", i)
					block.Jump()
				}
				runtime.broken = false
			})
		case script.PlusEquals:
			runtime.WriteStatement(func() {
				for i := a().(int); !runtime.broken && i > b().(int); i += c().(int) {
					runtime.Current.define("i", i)
					block.Jump()
				}
				runtime.broken = false
			})
		case script.MinusEquals:
			runtime.WriteStatement(func() {
				for i := a().(int); !runtime.broken && i > b().(int); i -= c().(int) {
					runtime.Current.define("i", i)
					block.Jump()
				}
				runtime.broken = false
			})
		case script.OverEquals:
			runtime.WriteStatement(func() {
				for i := a().(int); !runtime.broken && i > b().(int); i /= c().(int) {
					runtime.Current.define("i", i)
					block.Jump()
				}
				runtime.broken = false
			})
		case script.LeftOverEquals:
			runtime.WriteStatement(func() {
				for i := a().(int); !runtime.broken && i > b().(int); i %= c().(int) {
					runtime.Current.define("i", i)
					block.Jump()
				}
				runtime.broken = false
			})
		case script.TimesEquals:
			runtime.WriteStatement(func() {
				for i := a().(int); !runtime.broken && i > b().(int); i *= c().(int) {
					runtime.Current.define("i", i)
					block.Jump()
				}
				runtime.broken = false
			})
		}
	case script.Equals:
		switch action.Operator {
		case script.Plus1:
			runtime.WriteStatement(func() {
				for i := a().(int); !runtime.broken && i == b().(int); i++ {
					runtime.Current.define("i", i)
					block.Jump()
				}
				runtime.broken = false
			})
		case script.Minus1:
			runtime.WriteStatement(func() {
				for i := a().(int); !runtime.broken && i == b().(int); i-- {
					runtime.Current.define("i", i)
					block.Jump()
				}
				runtime.broken = false
			})
		case script.PlusEquals:
			runtime.WriteStatement(func() {
				for i := a().(int); !runtime.broken && i == b().(int); i += c().(int) {
					runtime.Current.define("i", i)
					block.Jump()
				}
				runtime.broken = false
			})
		case script.MinusEquals:
			runtime.WriteStatement(func() {
				for i := a().(int); !runtime.broken && i == b().(int); i -= c().(int) {
					runtime.Current.define("i", i)
					block.Jump()
				}
				runtime.broken = false
			})
		case script.OverEquals:
			runtime.WriteStatement(func() {
				for i := a().(int); !runtime.broken && i == b().(int); i /= c().(int) {
					runtime.Current.define("i", i)
					block.Jump()
				}
				runtime.broken = false
			})
		case script.LeftOverEquals:
			runtime.WriteStatement(func() {
				for i := a().(int); !runtime.broken && i == b().(int); i %= c().(int) {
					runtime.Current.define("i", i)
					block.Jump()
				}
				runtime.broken = false
			})
		case script.TimesEquals:
			runtime.WriteStatement(func() {
				for i := a().(int); !runtime.broken && i == b().(int); i *= c().(int) {
					runtime.Current.define("i", i)
					block.Jump()
				}
				runtime.broken = false
			})
		}

	case script.LessThanEquals:
		switch action.Operator {
		case script.Plus1:
			runtime.WriteStatement(func() {
				for i := a().(int); !runtime.broken && i <= b().(int); i++ {
					runtime.Current.define("i", i)
					block.Jump()
				}
				runtime.broken = false
			})
		case script.Minus1:
			runtime.WriteStatement(func() {
				for i := a().(int); !runtime.broken && i <= b().(int); i-- {
					runtime.Current.define("i", i)
					block.Jump()
				}
				runtime.broken = false
			})
		case script.PlusEquals:
			runtime.WriteStatement(func() {
				for i := a().(int); !runtime.broken && i <= b().(int); i += c().(int) {
					runtime.Current.define("i", i)
					block.Jump()
				}
				runtime.broken = false
			})
		case script.MinusEquals:
			runtime.WriteStatement(func() {
				for i := a().(int); !runtime.broken && i <= b().(int); i -= c().(int) {
					runtime.Current.define("i", i)
					block.Jump()
				}
				runtime.broken = false
			})
		case script.OverEquals:
			runtime.WriteStatement(func() {
				for i := a().(int); !runtime.broken && i <= b().(int); i /= c().(int) {
					runtime.Current.define("i", i)
					block.Jump()
				}
				runtime.broken = false
			})
		case script.LeftOverEquals:
			runtime.WriteStatement(func() {
				for i := a().(int); !runtime.broken && i <= b().(int); i %= c().(int) {
					runtime.Current.define("i", i)
					block.Jump()
				}
				runtime.broken = false
			})
		case script.TimesEquals:
			runtime.WriteStatement(func() {
				for i := a().(int); !runtime.broken && i <= b().(int); i *= c().(int) {
					runtime.Current.define("i", i)
					block.Jump()
				}
				runtime.broken = false
			})
		}

	case script.MoreThanEquals:
		switch action.Operator {
		case script.Plus1:
			runtime.WriteStatement(func() {
				for i := a().(int); !runtime.broken && i >= b().(int); i++ {
					runtime.Current.define("i", i)
					block.Jump()
				}
				runtime.broken = false
			})
		case script.Minus1:
			runtime.WriteStatement(func() {
				for i := a().(int); !runtime.broken && i >= b().(int); i-- {
					runtime.Current.define("i", i)
					block.Jump()
				}
				runtime.broken = false
			})
		case script.PlusEquals:
			runtime.WriteStatement(func() {
				for i := a().(int); !runtime.broken && i >= b().(int); i += c().(int) {
					runtime.Current.define("i", i)
					block.Jump()
				}
				runtime.broken = false
			})
		case script.MinusEquals:
			runtime.WriteStatement(func() {
				for i := a().(int); !runtime.broken && i >= b().(int); i -= c().(int) {
					runtime.Current.define("i", i)
					block.Jump()
				}
				runtime.broken = false
			})
		case script.OverEquals:
			runtime.WriteStatement(func() {
				for i := a().(int); !runtime.broken && i >= b().(int); i /= c().(int) {
					runtime.Current.define("i", i)
					block.Jump()
				}
				runtime.broken = false
			})
		case script.LeftOverEquals:
			runtime.WriteStatement(func() {
				for i := a().(int); !runtime.broken && i >= b().(int); i %= c().(int) {
					runtime.Current.define("i", i)
					block.Jump()
				}
				runtime.broken = false
			})
		case script.TimesEquals:
			runtime.WriteStatement(func() {
				for i := a().(int); !runtime.broken && i >= b().(int); i *= c().(int) {
					runtime.Current.define("i", i)
					block.Jump()
				}
				runtime.broken = false
			})
		}
	}
}
