package sel

const (
	tSpace = iota
	tDot   // .a
	tHash  // #a
	tLB    // [
	tRB    // ]
	tColon // :
	// combinator
	tPlus  // +
	tComma // ,
	tGT    // >
	tTilde // ~
	// attr
	tEqual    // =
	tCaret    // ^
	tDollor   // $
	tAsterisk // *
	tHyphen   // -
	tPipe     // |
	// misc
	tLP    // (
	tRP    // )
	tQuote // "
	// ...
	tPipeEqual     // |=
	tTildeEqual    // ~=
	tCaretEqual    // ^=
	tDollorEqual   // $=
	tAsteriskEqual // *=
	tHyphenEqual   // -=

	tText
)

var r2tMap = map[rune]int{
	' ':  tSpace,
	'.':  tDot,
	'#':  tHash,
	',':  tComma,
	'(':  tLP,
	')':  tRP,
	'[':  tLB,
	']':  tRB,
	'\'': tQuote,
	'>':  tGT,
	':':  tColon,
	'+':  tPlus,
	'$':  tDollor,
	'*':  tAsterisk,
	'^':  tCaret,
	'=':  tEqual,
	'-':  tHyphen,
	'|':  tPipe,
	'~':  tTilde,
}
var t2rMap = map[int]rune{
	tSpace:    ' ',
	tDot:      '.',
	tHash:     '#',
	tComma:    ',',
	tLP:       '(',
	tRP:       ')',
	tLB:       '[',
	tRB:       ']',
	tQuote:    '\'',
	tGT:       '>',
	tColon:    ':',
	tPlus:     '+',
	tDollor:   '$',
	tAsterisk: '*',
	tCaret:    '^',
	tEqual:    '=',
	tHyphen:   '-',
	tPipe:     '|',
	tTilde:    '~',
}

type token struct {
	category int
	text     string
}
type tokenList []token

const (
	ctxBlank = iota
	ctxAttr
	ctxDot
	ctxHash
	ctxSpace
)

var oMap = map[rune]int{
	'^': tCaret,
	'$': tDollor,
	'*': tAsterisk,
	'-': tHyphen,
	'|': tPipe,
}

func nextContext(r rune) int {
	switch r {
	case '.':
		return ctxDot
	case '#':
		return ctxHash
	case '[':
		return ctxAttr
	}

	return ctxBlank
}

func parse(s string) tokenList {
	tl := tokenList{}
	ctx := ctxBlank
	buf := []rune{}
	for i, ch := range s {
		if ctx == ctxBlank {
			if nc := nextContext(ch); nc != ctxBlank {
				if len(buf) > 0 {
					tl = append(tl, token{category: tText, text: string(buf)})
					buf = []rune{}
				}

				if ch == '#' {
					tl = append(tl, token{category: tHash})
				} else if ch == '.' {
					tl = append(tl, token{category: tDot})
				} else if ch == '[' {
					ctx = ctxAttr
					tl = append(tl, token{category: tLB})
				}
			} else {
				buf = append(buf, ch)
			}
		} else if ctx == ctxAttr {
			if ch == ']' {
				if len(buf) > 0 {
					tl = append(tl, token{category: tText, text: string(buf)})
					buf = []rune{}
				}
				ctx = ctxBlank
				tl = append(tl, token{category: tRB})
			} else {
				buf = append(buf, ch)
			}
		} else if ch == '.' {
			if len(buf) > 0 {
				tl = append(tl, token{category: tText, text: string(buf)})
				buf = []rune{}
			}
			tl = append(tl, token{category: tDot})
			buf = []rune{}
		} else if ch == '#' {
			if len(buf) > 0 {
				tl = append(tl, token{category: tText, text: string(buf)})
				buf = []rune{}
			}
			tl = append(tl, token{category: tHash})
			buf = []rune{}
		} else if ch == ' ' {
			if len(buf) > 0 {
				tl = append(tl, token{category: tText, text: string(buf)})
				buf = []rune{}
			}
			if i > 1 && tl[i-1].category != tSpace {
				tl = append(tl, token{category: tSpace})
			}
			buf = []rune{}
		} else {
			buf = append(buf, ch)
		}
	}
	if len(buf) > 0 {
		tl = append(tl, token{category: tText, text: string(buf)})
	}

	return tl
}
