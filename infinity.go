package infinity

import (
	"math/big"
)

// strictC is an infinite 0-1 sequence.
type strictC func(*big.Int) bool

// Cantor is lazy equivalent of strictC.
type Cantor func() strictC

// Predicate is ANY predicate on Cantor set.
type Predicate func(Cantor) bool

func toChar(b bool) byte {
	if b {
		return '1'
	}
	return '0'
}

// String returns n first digits of c.
func (c Cantor) String(n int64) string {
	var (
		f = c()
		s = make([]byte, 0, n+3)
	)

	for i := int64(0); i < n; i++ {
		s = append(s, toChar(f(big.NewInt(i))))
	}
	return string(append(s, '.', '.', '.'))
}

func makeLazy(c strictC) Cantor {
	return func() strictC { return c }
}

// TrueAt returns true for all infinite sequences which have
// true at their n-th position.
func TrueAt(n *big.Int) Predicate {
	return func(c Cantor) bool {
		return c()(n) == true
	}
}

// cons transforms infinite sequence
// a0,a1,a2,... -> f,a0,a1,a2,...
func cons(f bool, c Cantor) Cantor {
	return makeLazy(func(n *big.Int) bool {
		if n.IsUint64() && n.Uint64() == 0 {
			return f
		}
		return c()(new(big.Int).Sub(n, big.NewInt(1)))
	})
}

// withPrefix returns new Predicate which returns
// result of applying p to a sequence with added prefix.
func withPrefix(f bool, p Predicate) Predicate {
	return func(a Cantor) bool { return p(cons(f, a)) }
}

// Exists checks if p is true on some point and returns it.
// If p is always false, result is undefined.
func Exists(p Predicate) Cantor {
	t := func() strictC { return Exists(withPrefix(false, p))() }
	if p(cons(false, t)) {
		return cons(false, t)
	}
	return cons(true, func() strictC { return Exists(withPrefix(true, p))() })
}
