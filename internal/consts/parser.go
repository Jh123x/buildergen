package consts

type Mode string

const (
	MODE_FAST Mode = "FAST"
	MODE_AST       = "AST"
)

var (
	ALL_MODES = []Mode{
		MODE_FAST,
		MODE_AST,
	}
)
