package script

func unique(name []string) string {
	var register string
	if len(name) > 0 {
		register = name[0]
	} else {
		register = Unique()
	}

	if register[0] >= '0' && register[0] <= '9' {
		register = "_" + register
	}

	return register
}
