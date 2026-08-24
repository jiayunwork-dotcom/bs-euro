package parity

import (
	"math"
	"testing"

	"bs-euro/internal/bs"
)

func TestClosureErrorZeroForConsistentPrices(t *testing.T) {
	s, k, maturity, r, sigma := 100.0, 100.0, 1.0, 0.03, 0.2
	c := bs.CallPrice(s, k, maturity, r, sigma)
	p := bs.PutPrice(s, k, maturity, r, sigma)
	if !IsClosed(c, p, s, k, r, maturity, MaximumClosureTolerance()) {
		t.Errorf("parity should close")
	}
}

func TestSymmetricAtMoney(t *testing.T) {
	s, k, maturity, sigma := 100.0, 100.0, 1.0, 0.2
	c := bs.CallPrice(s, k, maturity, 0, sigma)
	p := bs.PutPrice(s, k, maturity, 0, sigma)
	if !SymmetricAtMoney(c, p) {
		t.Errorf("ATM with r=0 should give C=P")
	}
}

func TestCallFromPutInverts(t *testing.T) {
	s, k, r, maturity := 100.0, 100.0, 0.03, 1.0
	c := bs.CallPrice(s, k, maturity, r, 0.2)
	p := bs.PutPrice(s, k, maturity, r, 0.2)
	if math.Abs(CallFromPut(p, s, k, r, maturity)-c) > 1e-9 {
		t.Errorf("derived call should match")
	}
}

func TestStrikeFromParity(t *testing.T) {
	c, p := 10.0, 6.0
	s, r, maturity := 100.0, 0.03, 1.0
	k := StrikeFromParity(c, p, s, r, maturity)
	if k <= 0 {
		t.Errorf("strike = %g, want positive", k)
	}
}

func TestRelativeClosureError(t *testing.T) {
	err := RelativeClosureError(10, 6, 100, 100, 0.03, 1)
	if err < 0 {
		t.Errorf("relative error = %g, want non-negative", err)
	}
}
