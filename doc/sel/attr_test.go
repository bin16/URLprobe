package sel

import "testing"

func TestParseAttr(t *testing.T) {
	tParseAttr(t, "[name]")(&selector{
		category: stAttr,
		name:     "name",
		operator: opExists,
	})
	tParseAttr(t, "[name=content]")(&selector{
		category: stAttr,
		name:     "name",
		value:    "content",
		operator: opFullMatch,
	})
	tParseAttr(t, "[name^=content]")(&selector{
		category: stAttr,
		name:     "name",
		value:    "content",
		operator: opStartsWith,
	})
	tParseAttr(t, "[name$=content]")(&selector{
		category: stAttr,
		name:     "name",
		value:    "content",
		operator: opEndsWith,
	})
	tParseAttr(t, "[name*=content]")(&selector{
		category: stAttr,
		name:     "name",
		value:    "content",
		operator: opContains,
	})
	tParseAttr(t, "[name|=content]")(&selector{
		category: stAttr,
		name:     "name",
		value:    "content",
		operator: opListOf,
	})
}
func tParseAttr(t *testing.T, s string) func(s0 *selector) {
	return func(s0 *selector) {
		t.Helper()
		s1 := parseAttrSelector(s)
		if s1.category != s0.category {
			t.Errorf("Failed: parseAttrSelector(%s), category\n got %v,\nwant %v.", s, s1.category, s0.category)
		}
		if s1.name != s0.name {
			t.Errorf("Failed: parseAttrSelector(%s), name\n got %v,\nwant %v.", s, s1.name, s0.name)
		}
		if s1.value != s0.value {
			t.Errorf("Failed: parseAttrSelector(%s), value\n got %v,\nwant %v.", s, s1.value, s0.value)
		}
		if s1.operator != s0.operator {
			t.Errorf("Failed: parseAttrSelector(%s), operator\n got %v,\nwant %v.", s, s1.operator, s0.operator)

		}
	}
}
