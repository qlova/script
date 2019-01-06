package script

import (
	"fmt"
	"bytes"
	"errors"
	//"io"
	
	//"os"
	"runtime/debug"
	
	"math/big"
	"encoding/base64"
)

import "github.com/qlova/script/interpreter"
import "github.com/qlova/script/language"
import "github.com/qlova/script/go"

import (
	Golang "github.com/qlova/script/language/go"
)

var id int64 = 0;
func Unique() Go.String {
	id++
	return base64.RawURLEncoding.EncodeToString(big.NewInt(id).Bytes())
}

type SourceCode struct {
	Error bool
	ErrorMessage Go.String
	Data []byte
}

func (code SourceCode) String() Go.String {
	return Go.String(code.Data)
}

type Program func(Script)

func (program Program) SourceCode(lang language.Interface) (code SourceCode) {
	script := NewScript()
	script.lang = lang

	defer func() {
		if r := recover(); r != nil {
			code.Error = true
			code.ErrorMessage = fmt.Sprint(r, "\n", Go.String(debug.Stack()))
		}
	}()

	program(script)
	
	var buffer bytes.Buffer
	buffer.WriteString(Go.String(script.lang.Head()))
	buffer.WriteString(Go.String(script.lang.Neck()))
	buffer.WriteString(Go.String(script.lang.Body()))
	buffer.WriteString(Go.String(script.lang.Tail()))
	buffer.WriteString(Go.String(script.lang.Last()))

	buffer.Write(script.head.Bytes())
	buffer.Write(script.neck.Bytes())
	buffer.Write(script.body.Bytes())
	buffer.Write(script.tail.Bytes())
	buffer.Write(script.last.Bytes())

	code.Data = buffer.Bytes()

	return
}
//Return the programs SourceCode in Go.
func (program Program) Go() (code SourceCode) {
	return program.SourceCode(Golang.Language())
}

//Starts the program and waits for it to complete.
func (program Program) Run() (err error) {
	script := NewScript()

	script.lang = interpreter.New()

	//Catch errors.
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprint(r, "\n", Go.String(debug.Stack())))
		}
	}()
	
	program(script)
	
	script.lang.(interpreter.Implementation).Start()

	return
}

/*func (p Program) Source(language language.Interface) (source string, err error) {
	script := NewScript()

	script.lang = language
	script.lang.Init()

	//Catch errors.
	/*defer func() {
		if r := recover(); r != nil {
			if message, ok := r.(string); ok {
				err = errors.New(message)
			} else {
				err = errors.New(fmt.Sprint(r))
			}
		}
	}()*/
	
	/*p(script)
	
	script.head.WriteString(string(script.lang.Head()))
	script.neck.WriteString(string(script.lang.Neck()))
	script.body.WriteString(string(script.lang.Body()))
	script.tail.WriteString(string(script.lang.Tail()))
	script.last.WriteString(string(script.lang.Last()))
	
	source = string(script.head.Bytes())+string(script.neck.Bytes())+string(script.body.Bytes())+string(script.tail.Bytes())+string(script.last.Bytes())
	
	return
}

func (p Program) WriteToFile(path string, language language.Interface) (err error) {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	
	script := NewScript()

	script.lang = language
	script.lang.Init()

	//Catch errors.
	/*defer func() {
		if r := recover(); r != nil {
			if message, ok := r.(string); ok {
				err = errors.New(message)
			} else {
				err = errors.New(fmt.Sprint(r))
			}
		}
	}()*/
	
/*	p(script)
	
	script.head.WriteString(string(script.lang.Head()))
	script.neck.WriteString(string(script.lang.Neck()))
	script.body.WriteString(string(script.lang.Body()))
	script.tail.WriteString(string(script.lang.Tail()))
	script.last.WriteString(string(script.lang.Last()))
	
	file.Write(script.head.Bytes())
	file.Write(script.neck.Bytes())
	file.Write(script.body.Bytes())
	file.Write(script.tail.Bytes())
	file.Write(script.last.Bytes())
	
	return
}*/

type Script struct {
	*script
}

type Type interface {
	convert(Script) language.Type
	
	String() string
	Int() int
}

type script struct {
	depth Go.Int
	
	lang language.Interface
	
	head bytes.Buffer
	neck bytes.Buffer
	body bytes.Buffer
	tail bytes.Buffer
	last bytes.Buffer
	
	returns []Type
	
	Optimise Go.Bool
}

func NewScript() Script {
	return Script{script: new(script)}
}

func (q Script) Init() {
	q.lang.Init()
}

func (q Script) Last() {
	q.lang.Last()
}

func (q Script) indent() {
	for i := 0; i < q.depth; i++ {
		q.body.WriteByte('\t')
	}
}

func (q Script) write(s language.Statement) {
	q.body.WriteString(Go.String(s))
}

func (q Script) Raw(language Go.String, statement language.Statement) {
	if q.lang.Name() == language {
		q.write(statement)
	}
}

/*
	Main is the entry point of the program, this will be called when the program is executed.
*/
func (q Script) Main(f func()) {
	q.write(q.lang.Main())
	q.depth++
		f()
	q.depth--
	q.write(q.lang.EndMain())
}
