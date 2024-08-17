package utils

import "github.com/Jh123x/buildergen/internal/consts"

var kwMap = KeywordMap{}

type KeywordMap struct{}

func NewKeywordMap() *KeywordMap {
	return &kwMap
}

func (k *KeywordMap) IsKeyword(keyword string) bool {
	switch keyword {
	case consts.KEYWORD_GO, consts.KEYWORD_IF, consts.KEYWORD_FOR, consts.KEYWORD_MAP, consts.KEYWORD_VAR, consts.KEYWORD_BREAK, consts.KEYWORD_CASE, consts.KEYWORD_CHAN, consts.KEYWORD_ELSE, consts.KEYWORD_FUNC, consts.KEYWORD_GOTO, consts.KEYWORD_TYPE, consts.KEYWORD_CONST, consts.KEYWORD_DEFER, consts.KEYWORD_RANGE, consts.KEYWORD_RETURN, consts.KEYWORD_SELECT, consts.KEYWORD_STRUCT, consts.KEYWORD_SWITCH, consts.KEYWORD_IMPORT, consts.KEYWORD_DEFAULT, consts.KEYWORD_PACKAGE, consts.KEYWORD_CONTINUE, consts.KEYWORD_INTERFACE, consts.KEYWORD_FALLTHROUGH:
		return true
	default:
		return false
	}
}
