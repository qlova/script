package language

//These are the three options available when configuring a language.
const (
	//Produce deterministic results. Specificiation TBA.
	Deterministic = iota
	
	//Sacrifice cross-language determinism for pretty human readable code.
	Pretty
	
	//Sacrifice cross-language determinism for peformance, language specific calls will be used where possible.
	Fast
)
