package interpreter

import (
	"bufio"
	"os"
)

import "strconv"
import "github.com/qlova/script/language"

import "github.com/qlova/script/interpreter/dynamic"

func (implementation Implementation) Open(protocol string, path language.String) language.Type {
	panic(implementation.Name() + ".Open() Unimplemented")
	return nil
}

func (implementation Implementation) Load(protocol string, path language.String) language.Type {
	panic(implementation.Name() + ".Load() Unimplemented")
	return nil
}

var stdin = bufio.NewReader(os.Stdin)

func ReadString(char rune) string {
	text, _ := stdin.ReadString(byte(char))
	return text[:len(text)-1]
}

func (implementation Implementation) Read(stream language.Stream, mode language.Type) language.Type {

	if stream == nil {
		switch mode.(type) {
		case Symbol:
			var register = implementation.ReserveRegister()
			var char = implementation.RegisterOf(mode)
			implementation.AddInstruction(func(thread *dynamic.Thread) {
				thread.Set(register, ReadString(thread.Get(char).(rune)))
			})

			return String{
				Expression: language.Statement(strconv.Itoa(register)),
			}
		}
	}

	panic(implementation.Name() + ".Read() Unimplemented")
	return nil
}

func (implementation Implementation) Stop(stream language.Stream) language.Statement {
	panic(implementation.Name() + ".Stop() Unimplemented")
	return language.Statement("")
}

func (implementation Implementation) Seek(stream language.Stream, amount language.Integer) language.Statement {
	panic(implementation.Name() + ".Seek() Unimplemented")
	return language.Statement("")
}

func (implementation Implementation) Info(stream language.Stream, query language.String) language.String {
	panic(implementation.Name() + ".Info() Unimplemented")
	return nil
}

func (implementation Implementation) Move(stream language.Stream, location language.String) language.Statement {
	panic(implementation.Name() + ".Move() Unimplemented")
	return language.Statement("")
}
