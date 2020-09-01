package selector

import (
	"fmt"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	checklist := map[string]([]part){
		".btn": []part{
			part{pType: pDot},
			part{pType: pText, text: "btn"},
		},
		".btn a": []part{
			part{pType: pDot},
			part{pType: pText, text: "btn"},
			part{pType: pSpace},
			part{pType: pText, text: "a"},
		},
		".btn a, span": []part{
			part{pType: pDot},
			part{pType: pText, text: "btn"},
			part{pType: pSpace},
			part{pType: pText, text: "a"},
			part{pType: pComma},
			part{pType: pSpace},
			part{pType: pText, text: "span"},
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
			t.Errorf("Failed: parser(%s),\n got %s, want %s", s, plText(r1), plText(r0))
			return
		}
		for i, p := range r0 {
			if r1[i] != p {
				t.Errorf("Failed: parser(%s),\n got %s, want %s", s, plText(r1), plText(r0))
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
