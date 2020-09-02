package selector

import (
	"fmt"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	checklist := map[string]([]part){
		".btn a, span": []part{
			part{pType: pDot},
			part{pType: pText, text: "btn"},
			part{pType: pSpace},
			part{pType: pText, text: "a"},
			part{pType: pComma},
			part{pType: pSpace},
			part{pType: pText, text: "span"},
		},
		`.card-header + [name^="form_"] *`: []part{
			part{pType: pDot},
			part{pType: pText, text: "card-header"},
			part{pType: pSpace},
			part{pType: pPlus},
			part{pType: pSpace},

			part{pType: pLB},
			part{pType: pText, text: "name"},
			part{pType: pCaret},
			part{pType: pEqual},
			part{pType: pQuote},
			part{pType: pText, text: "form_"},
			part{pType: pQuote},
			part{pType: pRB},

			part{pType: pSpace},
			part{pType: pAsterisk},
		},
		"#preview .container button.btn-primary, a.btn[disabled]": []part{
			part{pType: pHash},
			part{pType: pText, text: "preview"},
			part{pType: pSpace},

			part{pType: pDot},
			part{pType: pText, text: "container"},
			part{pType: pSpace},

			part{pType: pText, text: "button"},
			part{pType: pDot},
			part{pType: pText, text: "btn-primary"},

			part{pType: pComma},
			part{pType: pSpace},

			part{pType: pText, text: "a"},
			part{pType: pDot},
			part{pType: pText, text: "btn"},
			part{pType: pLB},
			part{pType: pText, text: "disabled"},
			part{pType: pRB},
		},
	}

	for k, v := range checklist {
		tParse(t, k)(v)
	}
}
func tParse(t *testing.T, s string) func(r0 []part) {
	return func(r0 []part) {
		t.Helper()
		r1 := parse(s)
		if len(r0) != len(r1) {
			t.Errorf("Failed: parse(%s),\n got %s, \nwant %s", s, plText(r1), plText(r0))
			return
		}
		for i, p := range r0 {
			if r1[i] != p {
				t.Errorf("Failed: parse(%s),\n got %s, \nwant %s", s, plText(r1), plText(r0))
				return
			}
		}
	}
}
func plText(pl []part) string {
	sl := []string{}
	for _, p := range pl {
		if p.pType == pText {
			sl = append(sl, fmt.Sprintf("TEXT=%s", p.text))
		} else if p.pType == pSpace {
			sl = append(sl, "SPACE")
		} else {
			s := pTypeName(p.pType)
			sl = append(sl, s)
		}
	}
	return "{" + strings.Join(sl, "|") + "}"
}
