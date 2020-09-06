package sel

import (
	"strings"
	"unicode/utf8"
)

func parseAttrSelector(s string) *selector {
	sel := &selector{
		category: stAttr,
	}
	sl := strings.Split(unpack(s), "=")
	if len(sl) == 2 {
		l := sl[0]
		sel.value = sl[1]
		r, _ := utf8.DecodeLastRuneInString(sl[0])
		if strings.ContainsRune("^$*|", r) {
			sel.name = l[:len(l)-1]
			sel.operator = attrOperator([]rune{r, '='})
		} else {
			sel.name = l
			sel.operator = opFullMatch
		}
	} else {
		sel.name = sl[0]
		sel.operator = opExists
	}

	return sel
}

func unpack(s string) string {
	if strings.HasPrefix(s, "[") {
		s = s[1:]
	}
	if strings.HasSuffix(s, "]") {
		s = s[:len(s)-1]
	}

	return s
}

func attrOperator(rl []rune) int {
	switch string(rl) {
	case "^=":
		return opStartsWith
	case "$=":
		return opEndsWith
	case "*=":
		return opContains
	case "|=":
		return opListOf
	case "=":
		return opStartsWith
	}

	return 0
}
