package compiler

import "io"
import "os"
import "text/scanner"

import qlova "github.com/qlova/script"

import "reflect"
import "fmt"

type Type = qlova.Type

type Compiler struct {
	Syntax
	*qlova.Script
	
	CurrentFunction *Function
	
	Language Language
	
	GlobalScope Scope
	Scope []Scope
	
	Header io.Writer
	Output io.Writer

	StdErr []io.Writer
	
	DisableOutput bool
	
	
	Scanners []*scanner.Scanner

	CurrentLines []string
	LastLine string

	LineOffset int
	NextToken string
	DoneToken bool
	token string
	
	Errors bool
}

func New() *Compiler {
	var c = new(Compiler)
	c.Script = qlova.NewScript()
	c.StdErr = append(c.StdErr, os.Stdout)
	return c
}

func (c *Compiler) SwapOutput() {
	c.Header, c.Output = c.Output, c.Header
}

func (c *Compiler) Debug(T Type) {
	fmt.Println(reflect.TypeOf(T))
}

func (c *Compiler) Token() string {
	return c.token
}

func (c *Compiler) Peek() string {
	if c.NextToken != "" && !c.DoneToken {
		return c.NextToken
	}
	
	var token = c.Scanners[len(c.Scanners)-1].TokenText()
	c.NextToken = c.Scan()
	c.token = token
	return c.NextToken
}

func (c *Compiler) scan() string {
	if c.NextToken != "" && !c.DoneToken  {
		c.DoneToken = true
		c.token = c.NextToken
		return c.NextToken
	}
	if c.DoneToken {
		c.NextToken = ""
		c.DoneToken = false
	}
	
	tok := c.Scanners[len(c.Scanners)-1].Scan()
	
	for c.Scanners[len(c.Scanners)-1].TokenText() == " " || c.Scanners[len(c.Scanners)-1].TokenText() == "\t" {
		if c.Scanners[len(c.Scanners)-1].TokenText() != "\t" {
			c.CurrentLines[len(c.Scanners)-1] += c.Scanners[len(c.Scanners)-1].TokenText()
		}
		tok = c.Scanners[len(c.Scanners)-1].Scan()
	}
	
	if tok == scanner.EOF {
		c.CurrentLines = c.CurrentLines[:len(c.Scanners)-1]
		c.Scanners = c.Scanners[:len(c.Scanners)-1]
		return ""
	}
	
	c.CurrentLines[len(c.Scanners)-1] += c.Scanners[len(c.Scanners)-1].TokenText()
	
	if c.Scanners[len(c.Scanners)-1].TokenText() == "\n" {
		c.LastLine = c.CurrentLines[len(c.Scanners)-1][:len(c.CurrentLines[len(c.Scanners)-1])-1]
		c.CurrentLines[len(c.Scanners)-1] = ""
	}
	
	c.token = c.Scanners[len(c.Scanners)-1].TokenText()
	return c.Scanners[len(c.Scanners)-1].TokenText()
}

func (c *Compiler) Scan() string {
	var token = c.scan()
	
	if alias, ok := c.Aliases[token]; ok {
		return alias
	}
	
	return token
}

func (c *Compiler) ScanType(T Type) Type {
	var expression = c.ScanExpression()
	
	if !expression.SameAs(T) {
		
		fmt.Println(reflect.TypeOf(T), reflect.TypeOf(expression))
		c.ExpectingTypeName(T.Name(), expression.Name())
	}
	
	return expression
}

func (c *Compiler) ScanIf(test string) bool {
	if c.Peek() == test {
		c.Scan()
		return true
	}
	return false
}

func (c *Compiler) AddInput(input io.Reader) {
	var s = new(scanner.Scanner)
	s.Init(input)
	s.Whitespace = 0
	c.Scanners = append(c.Scanners, s)
	c.CurrentLines = append(c.CurrentLines, "")
}

func (c *Compiler) CompileBlock(first, last string) {	
	if first != "" {
		c.Expecting(first)
	}
	
	if len(c.Scanners) == 0 {
		return
	}
	
	for {
		if c.Peek() == last {
			c.Scan()
			return
		}
		
		c.ScanStatement()
		if len(c.Scanners) == 0 {
			return
		}
	}
}

func (c *Compiler) Compile() qlova.Program {
	
	return qlova.NewProgram(func(q *qlova.Script) {
		c.Script = q
		
		c.Script.Init()
		
		if c.GlobalScope.Variables == nil {
			c.GlobalScope = NewScope()
		}
		
		if len(c.Scanners) == 0 {
			return
		}
		
		for {
			c.ScanStatement()
			if len(c.Scanners) == 0 {
				c.Script.Last()
				return
			}
		}
	})
}
