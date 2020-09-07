package sel

import (
	"fmt"
	"strings"
	"testing"
)

func TestToChain(t *testing.T) {
	tToChain(t, tokenList{
		token{category: tText, text: "button"},
		token{category: tDot},
		token{category: tText, text: "btn"},
		token{category: tDot},
		token{category: tText, text: "btn-primary"},
	})(chain{
		&selector{category: stTag},
		&selector{category: stClass, value: "btn"},
		&selector{category: stClass, value: "btn-primary"},
	})
}
func tToChain(t *testing.T, tl tokenList) func(c0 chain) {
	return func(c0 chain) {
		t.Helper()
		c1 := toChain(tl)
		if len(c1) != len(c0) {
			t.Errorf("Failed: toChain(%v),\n got %v, \nwant %v", tl, c1, c0)
			return
		}

		for i, c := range c1 {
			if c.category != c0[i].category {
				t.Errorf("Failed: toChain(%v), category, \n got %v, \nwant %v", tl, c.category, c0[i].category)
				return
			}
			if c.name != c0[i].name {
				t.Errorf("Failed: toChain(%v), name, \n got %v, \nwant %v", tl, c.name, c0[i].name)
				return
			}
			if c.value != c0[i].value {
				t.Errorf("Failed: toChain(%v), value, \n got %v, \nwant %v", tl, c.value, c0[i].value)
				return
			}
		}
	}
}
func (c chain) String() string {
	sl := []string{}
	for i, s := range c {
		sl = append(sl, fmt.Sprintf("%d:c=%v,v=%s", i, s.category, s.value))
	}

	return strings.Join(sl, "\\")
}
