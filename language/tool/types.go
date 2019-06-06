package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

var NumberTypes = map[string]string{
	"real":       "r float64",
	"rational":   "",
	"natural":    "n uint",
	"integer":    "i int",
	"duplex":     "",
	"complex":    "",
	"quaternion": "",
	"octonion":   "",
	"sedenion":   "",
}

var HumanTypes = map[string]string{
	"symbol": "r rune",
	"string": "s string",

	"bit":  "b bool",
	"byte": "b byte",

	"color": "",
	"image": "",
	"sound": "",
	"video": "",

	"time": "",

	"stream": "",
}

var StructureTypes = map[string]string{
	"error": "",

	"array": "",
	"list":  "",

	"set":   "",
	"table": "",

	"queue": "",

	"tensor": "",
	"vector": "",
	"matrix": "",

	"pointer": "",
	"ring":    "",
	"tree":    "",
	"graph":   "",

	"function": "",
	
	"native": "",
}

func GenerateLanguageTypes() {
	file, err := os.Create(path.Dir(os.Args[0]) + "/../types.go")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	fmt.Fprintln(file, `package language
	
type NewType struct {
	Custom string
	Subtype Type
	Expression Statement
	Literal interface{}
	Length int
}

func (t NewType) Name() string { return t.Custom }
func (t NewType) Is(b Type) bool { c, ok := b.(NewType); return ok && c.Custom == t.Custom }
func (t NewType) Register(name string) Type { return NewType{Expression: Statement(name)} }
func (t NewType) Raw() Statement { return t.Expression }

type Statement = string

type Buffer interface {
	Buffer()
}

type Type interface { 
	Name() string
	Is(Type) bool
	Register(name string) Type
	Raw() Statement
}

type Number interface {
	Type
	Number()
}

type Dynamic interface {
	Type
	Dynamic()
}

type Metatype interface {
	Type
	Metatype()
}
`)

	for t := range NumberTypes {
		var title = strings.Title(t)
		fmt.Fprintln(file, "type "+title+" interface {")
		fmt.Fprintln(file, "\tType")
		fmt.Fprintln(file, "\tNumber()")
		fmt.Fprintln(file, "\t"+title+"()")
		fmt.Fprintln(file, "}")
	}
	for t := range HumanTypes {
		var title = strings.Title(t)
		fmt.Fprintln(file, "type "+title+" interface {")
		fmt.Fprintln(file, "\tType")
		fmt.Fprintln(file, "\t"+title+"()")
		fmt.Fprintln(file, "}")
	}
	for t := range StructureTypes {
		var title = strings.Title(t)
		fmt.Fprintln(file, "type "+title+" interface {")
		fmt.Fprintln(file, "\tType")
		fmt.Fprintln(file, "\t"+title+"()")
		fmt.Fprintln(file, "}")
	}
}
