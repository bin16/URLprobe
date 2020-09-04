package selector

const (
	pText = iota
	pSpace
	pDot      // .
	pHash     // #
	pComma    // ,
	pLP       // (
	pRP       // )
	pLB       // [
	pRB       // ]
	pColon    // :
	pQuote    // "
	pPlus     // +
	pGT       // >
	pCaret    // ^
	pDollor   // $
	pAsterisk // *
	pEqual    // =
	pHyphen   // -
	pPipe     // |
	pTilde    // ~

	pPipeEqual     // |=
	pTildeEqual    // ~=
	pCaretEqual    // ^=
	pDollorEqual   // $=
	pAsteriskEqual // *=
	pHyphenEqual   // -=
)

var pTypeMap = map[string]int{
	" ":  pSpace,
	".":  pDot,
	"#":  pHash,
	",":  pComma,
	"(":  pLP,
	")":  pRP,
	"[":  pLB,
	"]":  pRB,
	`"`:  pQuote,
	">":  pGT,
	":":  pColon,
	"-=": pHyphenEqual,
	"|=": pPipeEqual,
	"~=": pTildeEqual,
	"^=": pCaretEqual,
	"$=": pDollorEqual,
	"*=": pAsteriskEqual,
	"+":  pPlus,
	"$":  pDollor,
	"*":  pAsterisk,
	"^":  pCaret,
	"=":  pEqual,
	"-":  pHyphen,
	"|":  pPipe,
	"~":  pTilde,
}

type part struct {
	pType int
	text  string
}

func cut(s string) []string {
	units := []string{}
	buf := []rune{}
	rl := []rune(s)
	ppt := -1
	for _, ch := range rl {
		pt := pGetType([]rune{ch})
		if pt != ppt {
			if len(buf) > 0 {
				units = append(units, string(buf))
			}
			buf = []rune{ch}
		} else if pt == pSpace && len(buf) > 0 {
			continue
		} else {
			buf = append(buf, ch)
		}
		ppt = pt
	}
	if len(buf) > 0 {
		units = append(units, string(buf))
	}

	return units
}

func parseV2(s string) []part {
	pl := []part{}
	sl := cut(s)
	for _, s1 := range sl {
		pt := pGetType([]rune(s1))
		if pt != pText {
			pl = append(pl, part{pType: pt})
		} else {
			pl = append(pl, part{pType: pt, text: s1})
		}
	}

	return pl
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
