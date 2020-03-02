package language

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/qlova/script"
)

//Javascript returns a script formatted as Javascript source code.
func Javascript(f func(q script.Ctx)) []byte {
	var q = script.NewCtx()
	var language = new(javascript)
	language.imports = make(map[string]bool)
	q.Language = language
	f(q)
	return append(language.neck.Bytes(), language.Bytes()...)
}

type javascript struct {
	tabs int
	bytes.Buffer
	neck    bytes.Buffer
	imports map[string]bool
}

func (out *javascript) indent() {
	fmt.Fprint(out, strings.Repeat("\t", out.tabs))
}

func (out *javascript) Main(f func()) {
	f()
}

func (out *javascript) Raw(value script.Value) string {
	return out.valueOf(value)
}

func (out *javascript) valueOf(value script.Value) string {
	switch runtime := value.T().Get().(type) {
	case string:
		return strconv.Quote(runtime)
	case bool, int, expression, script.Variable:
		return fmt.Sprint(runtime)
	case map[string]interface{}: //Struct
		if reflect.TypeOf(value).Kind() == reflect.Ptr {
			return fmt.Sprintf("%v{}", reflect.TypeOf(value).Elem().Name())
		}
		return fmt.Sprintf("%v{}", reflect.TypeOf(value).Name())
	default:

		return fmt.Sprintf("%#v", runtime)

		//panic("go.valueOf: invalid type: " + reflect.TypeOf(runtime).String())
	}
}

func (out *javascript) typeOf(value script.Value) string {
	if value == nil {
		return ""
	}

	switch runtime := value.(type) {
	case script.String:
		return "string"
	case script.Int:
		return "int"
	default:
		return reflect.TypeOf(runtime).Name()
		panic("go.typeOf: invalid type: " + reflect.TypeOf(runtime).String())
	}
}

func (out *javascript) Set(a, b script.Value) {
	out.indent()
	fmt.Fprintf(out, "%v = %v;\n", out.valueOf(a), out.valueOf(b))
}

func (out *javascript) Field(structure script.Value, name string) script.Result {
	f := func() interface{} {
		return expression(fmt.Sprintf("%v.%v", out.valueOf(structure), name))
	}
	return &f
}

func (out *javascript) Index(a script.Value, b script.Int) script.Result {
	f := func() interface{} {
		return expression(fmt.Sprintf("%v[%v]", out.valueOf(a), out.valueOf(b)))
	}
	return &f
}

func (out *javascript) Mutate(a script.Value, b script.Int, c script.Value) {
	out.indent()
	fmt.Fprintf(out, "%v[%v] = %v;\n", out.valueOf(a), out.valueOf(b), out.valueOf(c))
}

func (out *javascript) Lookup(a script.Value, b script.String) script.Result {
	f := func() interface{} {
		return expression(fmt.Sprintf("%v[%v]", out.valueOf(a), out.valueOf(b)))
	}
	return &f
}

func (out *javascript) Insert(a script.Value, b script.String, c script.Value) {
	out.indent()
	fmt.Fprintf(out, "%v[%v] = %v;\n", out.valueOf(a), out.valueOf(b), out.valueOf(c))
}

func (out *javascript) DefineVariable(name string, value script.Value) script.Result {
	out.indent()
	fmt.Fprintf(out, "let %v = %v;\n", name, out.valueOf(value))
	f := func() interface{} {
		return expression(name)
	}
	return &f
}

func (out *javascript) DefineStruct(def script.Struct) {
	out.indent()
	fmt.Fprintf(out, "type %v struct {\n", def.Name)
	out.tabs++
	for _, field := range def.Fields {
		out.indent()
		fmt.Fprintf(out, "%v %v\n", field.Name, out.typeOf(field.Type))
	}
	out.tabs--
	fmt.Fprint(out, "}\n\n")
	for _, method := range def.Methods {
		out.indent()
		fmt.Fprintf(out, "func (%v %v) %v(", def.Reciever, def.Name, method.Name)
		if len(method.Args) > 0 {
			fmt.Fprintf(out, "%v %v", method.Args[0].Name, out.typeOf(method.Args[0].Value))
			for _, arg := range method.Args[1:] {
				fmt.Fprintf(out, ",%v %v", arg.Name, out.typeOf(arg.Value))
			}
		}
		fmt.Fprintf(out, ") %v {\n", out.typeOf(method.Returns))
		out.tabs++
		method.Block()
		out.tabs--
		fmt.Fprint(out, "}\n\n")
	}
}

func (out *javascript) Argument(name string, nth int) script.Result {
	f := func() interface{} {
		return expression(name)
	}
	return &f
}

func (out *javascript) If(first script.If, chain []script.If, last func()) {
	out.indent()
	fmt.Fprintf(out, "if (%v) {\n", out.valueOf(first.Condition))
	out.tabs++
	first.Block()
	out.tabs--
	out.indent()
	fmt.Fprintf(out, "}")
	if len(chain) > 0 {
		for _, link := range chain {
			fmt.Fprintf(out, " else if (%v) {\n", out.valueOf(link.Condition))
			out.tabs++
			link.Block()
			out.tabs--
			out.indent()
			fmt.Fprint(out, "}")
		}
	}
	if last != nil {
		fmt.Fprint(out, " else {\n")
		out.tabs++
		last()
		out.tabs--
		out.indent()
		fmt.Fprint(out, "}\n")
	}
}

func (out *javascript) For(set script.Int,
	condition script.ForLoopCondition,
	action script.ForLoopAction,
	f func(script.Int)) {

	out.indent()
	var last string
	if action.Operator != script.Plus1 && action.Operator != script.Minus1 {
		last = out.valueOf(action.Subject)
	}
	fmt.Fprintf(out, "for (let i = %v; i %v %v; i %v %v) {\n",
		out.valueOf(set),
		condition.Operator,
		out.valueOf(condition.Subject),
		action.Operator,
		last,
	)
	out.tabs++
	f(script.Int{Type: Expression(set.Ctx, "i")})
	out.tabs--
	out.indent()
	fmt.Fprintln(out, "}")
}

func (out *javascript) While(condition script.Bool, f func()) {
	out.indent()
	fmt.Fprintf(out, "while (%v) {\n", out.valueOf(condition))
	out.tabs++
	f()
	out.tabs--
	out.indent()
	fmt.Fprintln(out, "}")
}

func (out *javascript) Loop(f func()) {
	out.indent()
	fmt.Fprint(out, "while true {\n")
	out.tabs++
	f()
	out.tabs--
	out.indent()
	fmt.Fprintln(out, "}")
}

func (out *javascript) Break() {
	out.indent()
	fmt.Fprint(out, "break;\n")
}

func (out *javascript) Print(values script.Values) {
	out.indent()
	fmt.Fprint(out, "console.log(")
	if len(values) > 0 {
		fmt.Fprint(out, out.valueOf(values[0]))
		for _, value := range values[1:] {
			fmt.Fprintf(out, ",%v", out.valueOf(value))
		}
	}
	fmt.Fprintln(out, ")")
}

func (out *javascript) DefineFunction(function script.Function) {

	var name = function.Name
	var args = function.Args
	var returns = function.Returns
	var f = function.Block

	out.indent()
	fmt.Fprintf(out, "function %v(", name)
	if len(args) > 0 {
		fmt.Fprintf(out, "%v %v", args[0].Name, out.typeOf(args[0].Value))
		for _, arg := range args[1:] {
			fmt.Fprintf(out, ",%v %v", arg.Name, out.typeOf(arg.Value))
		}
	}
	fmt.Fprintf(out, ") %v {\n", out.typeOf(returns))
	out.tabs++
	f()
	out.tabs--
	fmt.Fprint(out, "}\n\n")
}

func (out *javascript) CallFunction(name string, args []script.Value) script.Result {
	var call = name + "("
	if len(args) > 0 {
		call += out.valueOf(args[0])
		for _, arg := range args[1:] {
			call += "," + out.valueOf(arg)
		}
	}
	call += ")"

	f := func() interface{} {
		return expression(call)
	}
	return &f
}

func (out *javascript) RunFunction(name string, args []script.Value) {
	out.indent()
	fmt.Fprintf(out, "%v(", name)
	if len(args) > 0 {
		fmt.Fprintf(out, "%v", out.valueOf(args[0]))
		for _, arg := range args[1:] {
			fmt.Fprintf(out, ",%v", out.valueOf(arg))
		}
	}
	fmt.Fprint(out, ");")
}

func (out *javascript) CallMethod(structure script.Value, name string, args []script.Value) script.Result {
	var call = fmt.Sprintf("%v.%v(", out.valueOf(structure), name)
	if len(args) > 0 {
		call += out.valueOf(args[0])
		for _, arg := range args[1:] {
			call += "," + out.valueOf(arg)
		}
	}
	call += ")"

	f := func() interface{} {
		return expression(call)
	}
	return &f
}

func (out *javascript) RunMethod(structure script.Value, name string, args []script.Value) {
	out.indent()
	fmt.Fprintf(out, "%v.%v(", out.valueOf(structure), name)
	if len(args) > 0 {
		fmt.Fprintf(out, "%v", out.valueOf(args[0]))
		for _, arg := range args[1:] {
			fmt.Fprintf(out, ",%v", out.valueOf(arg))
		}
	}
	fmt.Fprint(out, ");")
}

func (out *javascript) Return(v script.Value) {
	out.indent()
	if v == nil {
		fmt.Fprint(out, "return;\n")
		return
	}
	fmt.Fprintf(out, "return %v;\n", out.valueOf(v))
}

func (out *javascript) Plus(a, b script.Int) script.Int {
	return script.Int{
		Type: Expression(a.Ctx, fmt.Sprintf("(%v + %v)", out.valueOf(a), out.valueOf(b))),
	}
}

func (out *javascript) Not(b script.Bool) script.Bool {
	return script.Bool{
		Type: Expression(b.Ctx, fmt.Sprintf("(!%v)", out.valueOf(b))),
	}
}

func (out *javascript) Same(a, b script.Int) script.Bool {
	return script.Bool{
		Type: Expression(b.Ctx, fmt.Sprintf("(%v === %v)", out.valueOf(a), out.valueOf(b))),
	}
}

func (out *javascript) Join(a, b script.String) script.String {
	return script.String{
		Type: Expression(b.Ctx, fmt.Sprintf("(%v +%v)", out.valueOf(a), out.valueOf(b))),
	}
}
