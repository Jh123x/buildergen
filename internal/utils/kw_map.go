package utils

import "github.com/Jh123x/buildergen/internal/consts"

func IsKeyword(s string) bool {
	if len(s) < 2 || len(s) > 11 {
		return false
	}

	hash, i := 0, 0
	for i < len(s) {
		hash += int(s[i])
		i++
	}

	if len(consts.KwHashMap[hash%consts.HASH_IDX_MOD]) != len(s) {
		return false
	}

	return consts.KwHashMap[hash%consts.HASH_IDX_MOD] == s
}
