package consts

type Mode string

const (
	MODE_AST  Mode = "DEFAULT"
	MODE_FAST Mode = "FAST"
)

var (
	ALL_MODES = []Mode{
		MODE_FAST,
		MODE_AST,
	}
)
