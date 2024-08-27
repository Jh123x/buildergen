package consts

const (
	KEYWORD_GO          = "go"
	KEYWORD_IF          = "if"
	KEYWORD_FOR         = "for"
	KEYWORD_MAP         = "map"
	KEYWORD_VAR         = "var"
	KEYWORD_CASE        = "case"
	KEYWORD_CHAN        = "chan"
	KEYWORD_ELSE        = "else"
	KEYWORD_FUNC        = "func"
	KEYWORD_GOTO        = "goto"
	KEYWORD_TYPE        = "type"
	KEYWORD_BREAK       = "break"
	KEYWORD_CONST       = "const"
	KEYWORD_DEFER       = "defer"
	KEYWORD_RANGE       = "range"
	KEYWORD_RETURN      = "return"
	KEYWORD_SELECT      = "select"
	KEYWORD_STRUCT      = "struct"
	KEYWORD_SWITCH      = "switch"
	KEYWORD_IMPORT      = "import"
	KEYWORD_DEFAULT     = "default"
	KEYWORD_PACKAGE     = "package"
	KEYWORD_CONTINUE    = "continue"
	KEYWORD_INTERFACE   = "interface"
	KEYWORD_FALLTHROUGH = "fallthrough"

	HASH_IDX_MOD = 73
)

var (
	Keywords = [25]string{
		KEYWORD_GO,
		KEYWORD_IF,
		KEYWORD_FOR,
		KEYWORD_MAP,
		KEYWORD_VAR,
		KEYWORD_CASE,
		KEYWORD_CHAN,
		KEYWORD_ELSE,
		KEYWORD_FUNC,
		KEYWORD_GOTO,
		KEYWORD_TYPE,
		KEYWORD_BREAK,
		KEYWORD_CONST,
		KEYWORD_DEFER,
		KEYWORD_RANGE,
		KEYWORD_RETURN,
		KEYWORD_SELECT,
		KEYWORD_STRUCT,
		KEYWORD_SWITCH,
		KEYWORD_IMPORT,
		KEYWORD_DEFAULT,
		KEYWORD_PACKAGE,
		KEYWORD_CONTINUE,
		KEYWORD_INTERFACE,
		KEYWORD_FALLTHROUGH,
	}

	KwHashMap = [HASH_IDX_MOD]string{"", "switch", "", "goto", "", "", "break", "defer", "", "", "import", "default", "type", "", "range", "return", "fallthrough", "", "", "", "struct", "", "", "", "", "", "map", "", "", "", "", "", "", "", "", "for", "", "var", "", "", "const", "", "", "", "", "chan", "", "case", "", "", "", "", "", "", "", "", "select", "", "", "package", "else", "if", "", "func", "", "", "continue", "", "go", "interface", "", "", ""}
)
