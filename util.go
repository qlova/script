package script

func unique(name []string) string {
	var register string
	if len(name) > 0 {
		register = name[0]
	} else {
		register = Unique()
	}
	return register
}
