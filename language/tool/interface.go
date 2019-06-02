package main

import "os"
import "fmt"
import "strings"
import "path"

var LanguageInterface = [...]string{
	"Name() string",
	"Init()",
	"Build(path string) func()",
	
	"//Hooks",
	"Head() language.Statement",
	"Neck() language.Statement",
	"Body() language.Statement",
	"Tail() language.Statement",
	"Last() language.Statement",
	
	"//Context",
	"Buffer() language.Buffer",
	"Flush(buffer language.Buffer)",

	"//Variables",
	"Register(register string, value language.Type) (language.Statement, language.Type)",
	"Set(variable, value language.Type) language.Statement",

	"//Structures",
	"Index(structure, index language.Type) language.Type",
	"Modify(structure, index, value language.Type) language.Statement",
	
	"//Branching",
	"If(condition language.Bit) language.Statement",
	"ElseIf(condition language.Bit) language.Statement",
	"Else() language.Statement",
	"EndIf() language.Statement",
	
	"//Logic",
	"And(a, b language.Bit) language.Bit",
	"Or(a, b language.Bit) language.Bit",
	"Not(b language.Bit) language.Bit",
	"Equals(a, b language.Type) language.Bit",
	"Smaller(a, b language.Type) language.Bit",
	"Greater(a, b language.Type) language.Bit",
	
	"//Loops",
	"Loop() language.Statement",
	"EndLoop() language.Statement",
	"Break() language.Statement",
	"While(condition language.Bit) language.Statement",
	"EndWhile() language.Statement",
	"ForRange(i string, a, b language.Number) (language.Statement, language.Type)",
	"EndForRange() language.Statement",
	"ForEach(i, v string, list language.Type) (language.Statement, language.Type, language.Type)",
	"EndForEach() language.Statement",
	"For(i string, condition language.Bit, action language.Statement) (language.Statement, language.Type)",
	"EndFor() language.Statement",
	
	"//Entrypoint",
	"Main() language.Statement",
	"EndMain() language.Statement",
	"Exit() language.Statement",
	
	"//Function",
	"Function(name string, registers []string, arguments []language.Type, returns language.Type) (language.Statement, language.Function)",
	"EndFunction() language.Statement",
	"Call(f language.Function, arguments []language.Type) language.Type",
	"Run(f language.Function, arguments []language.Type) language.Statement",
	"Return(value language.Type) language.Statement",

	"//Threading",
	"Thread(name string, distance int, arguments []language.Type) language.Stream",
	
	"//IO",
	"Print(values ...language.Type) language.Statement",
	"Write(stream language.Stream, values ...language.Type) language.Statement",
	
	"//Streams",
	"Open(protocol string, path language.String) language.Type",
	"Load(protocol string, path language.String) language.Type",
	
	"Read(stream language.Stream, mode language.Type) language.Type",
	"Stop(stream language.Stream) language.Statement",
	"Seek(stream language.Stream, amount language.Integer) language.Statement",
	"Info(stream language.Stream, query language.String) language.String",
	"Move(stream language.Stream, location language.String) language.Statement",
	
	"//Errors",
	"Error() language.Error",
	"Trace(line int, file string) language.Statement",
	"Throw(code language.Number, message language.String) language.Statement",
	"Catch() language.Bit",
	
	"//Operations",
	"Add(a, b language.Number) language.Number",
	"Sub(a, b language.Number) language.Number",
	"Mul(a, b language.Number) language.Number",
	"Div(a, b language.Number) language.Number",
	"Pow(a, b language.Number) language.Number",
	"Mod(a, b language.Number) language.Number",
	
	"//Collections",
	"ArrayOf(t language.Type, length int) language.Array",
	"Join(a, b language.Type) language.Type",
	"Length(t language.Type) language.Integer",
	"Push(value language.Type, list language.Type) language.Statement",
	"Pop(list language.Type) language.Type",
	"TableOf(t language.Type) language.Table",
	"ListOf(t language.Type) language.List",
	
	"List(t ...language.Type) language.List",
	"Array(t ...language.Type) language.Array",
	"Table(index language.String, value language.Type) language.Table",
	
	"//Pointers",
	"PointerOf(t language.Type) language.Pointer",
	"Dereference(p language.Pointer) language.Type",
	
	"//Reflection",
	"DynamicOf(t language.Type) language.Dynamic",
	"StaticOf(d language.Dynamic) language.Type",
	"MetatypeOf(t language.Type) language.Metatype",
	
	"//Casts",
	"Cast(a, b language.Type) language.Type",
	
	"//Custom",
	"Type(name string, registers []string, elements []language.Type) language.Statement",
	"Method(t string, name string, registers []string, arguments []language.Type, returns language.Type) language.Statement",
	"This() language.Type",
	"New(name string) language.Type",	
	"Invoke(t language.Type, method string, arguments []language.Type) language.Type",
	"Execute(t language.Type, method string, arguments []language.Type) language.Statement",
	"EndMethod() language.Statement",
	
	"//Util",
	"Copy(t language.Type) language.Type",
}

func GenerateLanguageInterface() {
	file, err := os.Create(path.Dir(os.Args[0])+"/../interface.go")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	
	fmt.Fprintln(file, "package language\n\ntype Interface interface {")
	
	Literal := func(name string, sig string) {
		var title = strings.Title(name)
		fmt.Fprintln(file, "\t"+title+"("+sig+") "+title)
	}
	
	for a, b := range NumberTypes {
		Literal(a, b)
	}
	for a, b := range HumanTypes {
		Literal(a, b)
	}
	
	for _, line := range LanguageInterface {
		line = strings.Replace(line, "language.", "", -1)
		fmt.Fprintln(file, "\t"+line)
	}
	fmt.Fprintln(file, "}")
}

func GenerateLanguageTemplate(name string) {
	lower := strings.ToLower(name)
	
	os.Mkdir(path.Dir(os.Args[0])+"/../"+lower, 0755)
	
	
	
	file, err := os.Create(path.Dir(os.Args[0])+"/../"+lower+"/"+name+".go")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintln(file, "package", name)
	fmt.Fprintln(file)
	fmt.Fprintln(file, `import "github.com/qlova/script/language"`)
	fmt.Fprintln(file)
	fmt.Fprintln(file, "type Implementation struct {}")
	fmt.Fprintln(file)
	
	Literal := func(name string, sig string) {
		var title = strings.Title(name)
		fmt.Fprintln(file, "func (implementation Implementation) "+title+"("+sig+") language."+title+" { panic(`Unimplemented`); return nil }")
	}
	
	for a, b := range NumberTypes {
		Literal(a, b)
	}
	for a, b := range HumanTypes {
		Literal(a, b)
	}
	
	for _, line := range LanguageInterface {
		
		//Create a new file.
		if line[0] == '/' {
			if file != nil {
				file.Close()
			}
			
			file, err = os.Create(path.Dir(os.Args[0])+"/../"+lower+"/"+line[2:]+".go")
			if err != nil {
				fmt.Println(err)
				return
			}
			
			fmt.Fprintln(file, "package", name)
			fmt.Fprintln(file)
			fmt.Fprintln(file, `import "github.com/qlova/script/language"`)
			fmt.Fprintln(file)
			
			continue
		}
		
		
		//Add methods.
		fmt.Fprintln(file, "func (implementation Implementation)", line, "{")
		
		//Figure out return value
		kind := strings.TrimSpace(strings.Split(line, ")")[1])
		fname := strings.TrimSpace(strings.Split(line, "(")[0])
		
		fmt.Fprintln(file, "\t"+`panic(implementation.Name()+".`+fname+`() Unimplemented")`)
		
		if fname == "Register" {
			fmt.Fprintln(file, "\t"+`return language.Statement(""), nil`)
		} else if fname == "Function" {
			fmt.Fprintln(file, "\t"+`return language.Statement(""), nil`)
		} else if fname == "ForRange" {
			fmt.Fprintln(file, "\t"+`return language.Statement(""), nil`)
		} else if fname == "ForEach" {
			fmt.Fprintln(file, "\t"+`return language.Statement(""), nil, nil`)
		} else if fname == "For" {
			fmt.Fprintln(file, "\t"+`return language.Statement(""), nil`)
		} else {
		
			switch kind {
				
				case "language.Statement":
					fmt.Fprintln(file, "\t"+`return language.Statement("")`)
				
				case "string":
					fmt.Fprintln(file, "\t"+`return ""`)
				
				case "":
				
				default:
					fmt.Fprintln(file, "\t"+`return nil`)
			}
		}

		fmt.Fprintln(file, "}")
		fmt.Fprintln(file)
	}
	
	if file != nil {
		file.Close()
	}
	
	file, err = os.Create(path.Dir(os.Args[0])+"/../"+lower+"/Types.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintln(file, "package", name)
	fmt.Fprintln(file)
	fmt.Fprintln(file, `import "github.com/qlova/script/language"`)
	fmt.Fprintln(file)

	for t := range NumberTypes {
		var title = strings.Title(t)
		fmt.Fprintln(file, "type "+title+ " language.NewType")
		fmt.Fprintln(file, "func (t "+title+") Number() {}")
		fmt.Fprintln(file, "func (t "+title+") "+title+"() {}")
		fmt.Fprintln(file, "func (t "+title+") Name() string { return \""+t+"\" }")
		fmt.Fprintln(file, "func (t "+title+") Is(b language.Type) bool { _, ok := b.("+title+"); return ok }")
		fmt.Fprintln(file, "func (t "+title+") Register(name string) language.Type { var result = t; result.Expression = language.Statement(name); return result }")
		fmt.Fprintln(file, "func (t "+title+") Raw() language.Statement { return t.Expression }")
		fmt.Fprintln(file)
	}
	for t := range HumanTypes {
		var title = strings.Title(t)
		fmt.Fprintln(file, "type "+title+ " language.NewType")
		fmt.Fprintln(file, "func (t "+title+") Number() {}")
		fmt.Fprintln(file, "func (t "+title+") "+title+"() {}")
		fmt.Fprintln(file, "func (t "+title+") Name() string { return \""+t+"\" }")
		fmt.Fprintln(file, "func (t "+title+") Is(b language.Type) bool { _, ok := b.("+title+"); return ok }")
		fmt.Fprintln(file, "func (t "+title+") Register(name string) language.Type { var result = t; result.Expression = language.Statement(name); return result }")
		fmt.Fprintln(file, "func (t "+title+") Raw() language.Statement { return t.Expression }")
		fmt.Fprintln(file)
	}
	for t := range StructureTypes {
		var title = strings.Title(t)
		fmt.Fprintln(file, "type "+title+ " language.NewType")
		fmt.Fprintln(file, "func (t "+title+") Number() {}")
		fmt.Fprintln(file, "func (t "+title+") "+title+"() {}")
		fmt.Fprintln(file, "func (t "+title+") Name() string { return \""+t+"\" }")
		fmt.Fprintln(file, "func (t "+title+") Is(b language.Type) bool { _, ok := b.("+title+"); return ok }")
		fmt.Fprintln(file, "func (t "+title+") Register(name string) language.Type { var result = t; result.Expression = language.Statement(name); return result }")
		fmt.Fprintln(file, "func (t "+title+") Raw() language.Statement { return t.Expression }")
		fmt.Fprintln(file)
	}
}
