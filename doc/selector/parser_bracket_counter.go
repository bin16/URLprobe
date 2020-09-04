package selector

const (
	bCtxBlank       = pText
	bCtxBracket     = pLB
	bCtxParenthesis = pLP
	bCtxQuote       = pQuote
)

type bracketCounter []int

func (bc bracketCounter) push(b int) {
	if b == pLB || b == pRB {
		bc = append(bc, b)
	} else if b == pLP || b == pRP {
		bc = append(bc, b)
	} else if b == pQuote {
		bc = append(bc, b)
	}
}
func (bc bracketCounter) quoted() bool {
	quoted := false
	for _, p := range bc {
		if p == pQuote {
			quoted = !quoted
		}
	}

	return quoted
}

func (bc bracketCounter) context() int {
	var lb, lp int
	var quoted = false
	for _, p := range bc {
		switch p {
		case pQuote:
			quoted = !quoted
		case pLB:
			lb++
		case pRB:
			lb--
		case pLP:
			lp++
		case pRP:
			lp--
		}
	}

	if quoted {
		return pQuote
	} else if lb > 0 {
		return pLB
	} else if lp > 0 {
		return pLP
	}

	return 0
}
