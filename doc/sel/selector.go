package sel

const (
	stClass = iota
	stID
	stAttr
	stTag
	stGTComb    // a>b
	stSpaceComb // a b
	stPlusComb  // a+b
	stTildeComb // a~b

	opExists     // [a]
	opFullMatch  // =
	opContains   // *=
	opStartsWith // ^=
	opEndsWith   // $=
	opListOf     // |=
)

type selector struct {
	category int
	name     string
	value    string
	operator int
}

type chain []*selector

func toChain(tl tokenList) chain {
	sl := chain{}
	i := 0
	for i < len(tl) {
		t := tl[i]
		nt := tl[i+1]
		if t.category == tText {
			sl = append(sl, &selector{
				category: stTag,
				value:    nt.text,
			})
			i++
			continue
		} else if t.category == tDot {
			sl = append(sl, &selector{
				category: stClass,
				value:    nt.text,
			})
			i += 2
			continue
		} else if t.category == tHash {
			sl = append(sl, &selector{
				category: stID,
				value:    nt.text,
			})
			i += 2
			continue
		} else if t.category == tLB {
			sl = append(sl, parseAttrSelector(nt.text))
			i += 3
			continue
		}
	}

	return sl
}
