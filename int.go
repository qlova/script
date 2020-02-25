package script

//Int is the default integer type for the language.
type Int struct {
	Type
}

//Int returns a new Int from a Go int.
func (q Ctx) Int(literal int) Int {
	return Int{q.Literal(literal)}
}

//Set sets the integer.
func (a Int) Set(b Int) {
	a.Ctx.Set(a, b)
}

//Setc sets the integer.
func (a Int) SetL(b int) {
	a.Ctx.Set(a, a.Ctx.Int(b))
}

//IntList is an array of ints.
type IntList struct {
	List
	L []Int

	Set func(IntList)

	Index  func(Int) Int
	Mutate func(Int, Int)
}

//MutateL mutates the list with literal values.
func (array IntList) MutateL(index, mutation int) {
	array.Mutate(array.T().Ctx.Int(index), array.T().Ctx.Int(mutation))
}

//IndexL returns the value at the given literal index in the list.
func (array IntList) IndexL(index int) Int {
	return array.Index(array.T().Ctx.Int(index))
}

//IntList returns a new array of ints from given Ints.
func (q Ctx) IntList(elements ...Int) IntList {
	var array = IntList{L: elements}
	q.Make(&array, len(elements))
	return array
}

//IntListL returns a new array of ints from given literal elements.
func (q Ctx) IntListL(elements ...int) IntList {
	var values = make([]Int, len(elements))
	for i, element := range elements {
		values[i] = q.Int(element)
	}
	return q.IntList(values...)
}

//IntTable is a table of ints.
type IntTable struct {
	Table
	L map[String]Int

	Set func(IntTable)

	Lookup func(String) Int
	Insert func(String, Int)
}

//InsertL mutates the list with literal values.
func (table IntTable) InsertL(index string, mutation int) {
	table.Insert(table.T().Ctx.String(index), table.T().Ctx.Int(mutation))
}

//LookupL returns the value at the given literal index in the list.
func (table IntTable) LookupL(index string) Int {
	return table.Lookup(table.T().Ctx.String(index))
}

//IntTable returns a new table of ints from given Ints.
func (q Ctx) IntTable(elements map[String]Int) IntTable {
	var table = IntTable{L: elements}
	q.Make(&table)
	return table
}

//IntTableL returns a new table of ints from given literal elements.
func (q Ctx) IntTableL(elements map[string]int) IntTable {
	var values = make(map[String]Int, len(elements))
	for key, element := range elements {
		values[q.String(key)] = q.Int(element)
	}
	var table = IntTable{L: values}
	q.Make(&table)
	return table
}

func (a Int) Plus(b Int) Int {
	return a.Ctx.Language.Plus(a, b)
}

func (a Int) Return(result Int) {
	result.Ctx.Return(result)
}
