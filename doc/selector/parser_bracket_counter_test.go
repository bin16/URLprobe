package selector

import (
	"testing"
)

func TestBracketCounter(t *testing.T) {
	tBCounter(t, bracketCounter{
		pLB, pRB,
	})(bCtxBlank)
	tBCounter(t, bracketCounter{
		pQuote, pLP,
	})(bCtxQuote)
	tBCounter(t, bracketCounter{
		pLB, pQuote, pQuote, pRB,
	})(bCtxBlank)
	tBCounter(t, bracketCounter{
		pLB, pQuote, pQuote,
	})(bCtxBracket)
	tBCounter(t, bracketCounter{
		pLB, pQuote,
	})(bCtxQuote)
}
func tBCounter(t *testing.T, bc bracketCounter) func(r0 int) {
	return func(r0 int) {
		t.Helper()
		if r1 := bc.context(); r1 != r0 {
			t.Errorf("Failed: bc.context(); \n got %d; \nwant %d; \n  bc is %v", r1, r0, bc)
		}
	}
}
