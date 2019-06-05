package Go

import "github.com/qlova/script/language"

func (implementation Implementation) Head() language.Statement {
	var head language.Statement = "package main\n\n"

	for path := range implementation.Imports {
		head += `import "` + path + `"` + "\n"
	}

	head += "\n"

	return head
}

func (implementation Implementation) Neck() language.Statement {
	return language.Statement(implementation.neck.String())
}

func (implementation Implementation) Body() language.Statement {
	return language.Statement("")
}

func (implementation Implementation) Tail() language.Statement {
	return language.Statement("")
}

func (implementation Implementation) Last() language.Statement {
	return language.Statement("")
}
