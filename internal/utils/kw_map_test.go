package utils

import (
	"strings"
	"testing"

	"github.com/Jh123x/buildergen/internal/consts"
	"github.com/stretchr/testify/assert"
)

type kwTest struct {
	keyword     string
	expectedRes bool
}

func TestKeywordMap(t *testing.T) {
	tests := map[string]kwTest{
		"not key word": {
			keyword:     "notkw",
			expectedRes: false,
		},
	}

	for i := 0; i < 25; i++ {
		kw := consts.Keywords[i]
		tests["Keyword: "+kw] = kwTest{
			keyword:     kw,
			expectedRes: true,
		}
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expectedRes, IsKeyword(tc.keyword))
		})
	}
}

type algo func(string) bool

func BenchmarkKeywordLookup(b *testing.B) {
	testAlgorithms := map[string]algo{
		"Loop":           naiveIteration(),
		"Map":            naiveMap(),
		"Current":        IsKeyword,
		"Switch":         naiveSwitch(),
		"Optimization 1": attempt1(),
		"attempt 2":      attempt2(),
	}

	tests := make([]kwTest, 0, 50)

	for i := 0; i < 25; i++ {
		kw := consts.Keywords[i]
		kw1 := strings.Repeat("a", i)
		tests = append(
			tests,
			kwTest{keyword: kw, expectedRes: true},
			kwTest{keyword: kw1, expectedRes: false},
		)
	}

	for name, algorithm := range testAlgorithms {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, t := range tests {
					assert.Equal(b, t.expectedRes, algorithm(t.keyword), t.keyword)
				}
			}
		})
	}
}

func naiveIteration() algo {
	return func(s string) bool {
		for _, kw := range consts.Keywords {
			if kw == s {
				return true
			}
		}
		return false
	}
}

func naiveMap() algo {
	type empty struct{}
	mp := make(map[string]empty, 25)
	for i := 0; i < 25; i++ {
		mp[consts.Keywords[i]] = empty{}
	}

	return func(s string) bool {
		_, ok := mp[s]
		return ok
	}
}

func naiveSwitch() algo {
	return func(s string) bool {
		switch s {
		case consts.KEYWORD_GO, consts.KEYWORD_IF, consts.KEYWORD_FOR, consts.KEYWORD_MAP, consts.KEYWORD_VAR, consts.KEYWORD_BREAK, consts.KEYWORD_CASE, consts.KEYWORD_CHAN, consts.KEYWORD_ELSE, consts.KEYWORD_FUNC, consts.KEYWORD_GOTO, consts.KEYWORD_TYPE, consts.KEYWORD_CONST, consts.KEYWORD_DEFER, consts.KEYWORD_RANGE, consts.KEYWORD_RETURN, consts.KEYWORD_SELECT, consts.KEYWORD_STRUCT, consts.KEYWORD_SWITCH, consts.KEYWORD_IMPORT, consts.KEYWORD_DEFAULT, consts.KEYWORD_PACKAGE, consts.KEYWORD_CONTINUE, consts.KEYWORD_INTERFACE, consts.KEYWORD_FALLTHROUGH:
			return true
		default:
			return false
		}
	}
}

func attempt1() algo {
	buckets := [...][2]int{
		{0, 2},   // 2
		{2, 5},   // 3
		{5, 11},  // 4
		{11, 15}, // 5
		{15, 20}, // 6
		{20, 22}, // 7
		{22, 23}, // 8
		{23, 24}, // 9
		{0, 0},   // 10
		{24, 25}, // 11
	}
	return func(s string) bool {
		if len(s) < 2 || len(s) > 11 {
			return false
		}

		rangeVal := buckets[len(s)-2]
		for i := rangeVal[0]; i < rangeVal[1]; i++ {
			if consts.Keywords[i] == s {
				return true
			}
		}

		return false
	}
}

func attempt2() algo {
	KwHashMap := [73]string{"", "switch", "", "goto", "", "", "break", "defer", "", "", "import", "default", "type", "", "range", "return", "fallthrough", "", "", "", "struct", "", "", "", "", "", "map", "", "", "", "", "", "", "", "", "for", "", "var", "", "", "const", "", "", "", "", "chan", "", "case", "", "", "", "", "", "", "", "", "select", "", "", "package", "else", "if", "", "func", "", "", "continue", "", "go", "interface", "", "", ""}
	return func(s string) bool {
		if len(s) < 2 || len(s) > 11 {
			return false
		}
		hash, i := 0, 0
		for i < len(s) {
			hash += int(s[i])
			i++
		}

		if len(KwHashMap[hash%73]) != len(s) {
			return false
		}

		return KwHashMap[hash%73] == s
	}
}
