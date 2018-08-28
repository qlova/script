package script

import (
	"fmt"
	"bytes"
	"errors"
	//"io"
)

import "github.com/qlova/script/interpreter"
import "github.com/qlova/script/language"

type Program struct {
	program func(*Script)
}

func NewProgram(program func(*Script)) Program {
	return Program{program:program}
}

//Starts the program and waits for it to complete.
func (p Program) Run() (err error) {
	script := NewScript()
	
	interpreter := Interpreter.New()

	script.lang = interpreter

	//Catch errors.
	defer func() {
		if r := recover(); r != nil {
			if message, ok := r.(string); ok {
				err = errors.New(message)
			} else {
				err = errors.New(fmt.Sprint(r))
			}
		}
	}()
	
	p.program(script)
	
	interpreter.Start()

	return
}

func (p Program) Source(language language.Interface) (source string, err error) {
	script := NewScript()

	script.lang = language
	script.lang.Init()

	//Catch errors.
	defer func() {
		if r := recover(); r != nil {
			if message, ok := r.(string); ok {
				err = errors.New(message)
			} else {
				err = errors.New(fmt.Sprint(r))
			}
		}
	}()
	
	p.program(script)
	
	script.head.WriteString(string(script.lang.Head()))
	script.neck.WriteString(string(script.lang.Neck()))
	script.body.WriteString(string(script.lang.Body()))
	script.tail.WriteString(string(script.lang.Tail()))
	script.last.WriteString(string(script.lang.Last()))
	
	source = string(script.head.Bytes())+string(script.neck.Bytes())+string(script.body.Bytes())+string(script.tail.Bytes())+string(script.last.Bytes())
	
	return
}

type Script struct {
	depth int
	
	lang language.Interface
	
	head bytes.Buffer
	neck bytes.Buffer
	body bytes.Buffer
	tail bytes.Buffer
	last bytes.Buffer
}

func NewScript() *Script {
	return new(Script)
}

func (q *Script) indent() {
	for i := 0; i < q.depth; i++ {
		q.body.WriteByte('\t')
	}
}

func (q *Script) write(s language.Statement) {
	q.body.WriteString(string(s))
}

func (q *Script) Main(f func(*Script)) {
	q.write(q.lang.Main())
	q.depth++
		f(q)
	q.depth--
	q.write(q.lang.EndMain())
}
