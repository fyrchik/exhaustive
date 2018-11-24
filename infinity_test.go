package infinity

import (
	"math/big"
	"testing"
)

func alwaysFalse(c Cantor) bool { return false }

func TestTrueAt(t *testing.T) {
	return
}

func TestExists(t *testing.T) {
	var (
		r        Cantor
		p        Predicate
		numTests int64 = 15
	)

	// here we just need to assure that this call take finite time
	r = Exists(alwaysFalse)

	for i := int64(0); i < numTests; i++ {
		p = TrueAt(big.NewInt(i))
		r = Exists(p)
		if !p(r) {
			t.Errorf("expected TrueAt(%s) to find correct point", r.String(numTests))
		}
	}
}
