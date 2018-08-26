package Go

import qlova "github.com/qlova/script"

func (l *language) ForEachList(i, variable, list string) string {
	if i == "" {
		i = "_"
	}
	
	return "for "+i+", "+variable+" := range "+list+" {\n"
}

func (l *language) EndForEachList() string {
	return "}\n"
}
