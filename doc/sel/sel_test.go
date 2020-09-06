package sel

import (
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	tParse(t, ".btn.btn-primary")(tokenList{
		token{category: tDot},
		token{category: tText, text: "btn"},
		token{category: tDot},
		token{category: tText, text: "btn-primary"},
	})
	tParse(t, "[name]")(tokenList{
		token{category: tLB},
		token{category: tText, text: "name"},
		token{category: tRB},
	})
	tParse(t, "[name=content]")(tokenList{
		token{category: tLB},
		token{category: tText, text: "name=content"},
		token{category: tRB},
	})
	tParse(t, "textarea[name=content]")(tokenList{
		token{category: tText, text: "textarea"},
		token{category: tLB},
		token{category: tText, text: "name=content"},
		token{category: tRB},
	})
	tParse(t, "textarea.editor.active[name=content]")(tokenList{
		token{category: tText, text: "textarea"},
		token{category: tDot},
		token{category: tText, text: "editor"},
		token{category: tDot},
		token{category: tText, text: "active"},
		token{category: tLB},
		token{category: tText, text: "name=content"},
		token{category: tRB},
	})
}
func tParse(t *testing.T, s string) func(t0 tokenList) {
	return func(t0 tokenList) {
		t.Helper()
		t1 := parse(s)
		if len(t1) != len(t0) {
			t.Errorf("Failed: parse(%s),\n got %v, \nwant %v", s, t1, t0)
			return
		}

		for i, p := range t0 {
			if t1[i] != p {
				t.Errorf("Failed: parse(%s),\n got %v, \nwant %v", s, t1, t0)
				return
			}
		}
	}
}
func (tl tokenList) String() string {
	s := []string{}
	for _, t := range tl {
		if ch, ok := t2rMap[t.category]; ok {
			s = append(s, string([]rune{ch}))
		} else {
			s = append(s, t.text)

		}
	}

	return strings.Join(s, "\\")
}
