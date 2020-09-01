package selector

const (
	pText = iota
	pSpace
	pDot   // .
	pHash  // #
	pComma // ,
	pLP    // (
	pRP    // )
	pLB    // [
	pRB    // ]
	pColon // :
	pQuote // "
	pPlus  // +
	pGT    // >
)

var pTypeMap = map[string]int{
	" ": pSpace,
	".": pDot,
	"#": pHash,
	",": pComma,
	"(": pLP,
	")": pRP,
	"[": pLB,
	"]": pRB,
	":": pColon,
	`"`: pQuote,
	"+": pPlus,
	">": pGT,
}

type part struct {
	pType int
	text  string
}

func parse(s string) []part {
	pl := []part{}
	rl := []rune(s)
	buf := []rune{}
	ppt := -1
	for _, ch := range rl {
		pt := pGetType([]rune{ch})
		if pt == pText {
			buf = append(buf, ch)
		} else if pt == pSpace {
			if ppt == pSpace {
				continue
			} else if ppt == pText {
				pl = append(pl, part{pType: pText, text: string(buf)})
				buf = []rune{}
			}
			pl = append(pl, part{pType: pSpace})
		} else {
			if ppt == pText {
				pl = append(pl, part{pType: pText, text: string(buf)})
				buf = []rune{}
			}
			pl = append(pl, part{pType: pt})
		}

		ppt = pt
	}
	if len(buf) > 0 {
		pl = append(pl, part{pType: pText, text: string(buf)})
	}

	return pl
}

func pTypeName(i int) string {
	for pt, pn := range pTypeMap {
		if pn == i {
			return pt
		}
	}

	return "{TEXT}"
}

func pGetType(rl []rune) int {
	s := string(rl)
	pt, ok := pTypeMap[s]
	if !ok {
		return pText
	}

	return pt
}
