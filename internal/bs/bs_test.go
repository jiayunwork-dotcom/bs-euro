package bs

import (
	"math"
	"testing"
)

func TestCallIntrinsicAtExpiry(t *testing.T) {
	got := CallPrice(120, 100, 0, 0, 0)
	if got != 20 {
		t.Errorf("call = %g, want 20", got)
	}
}

func TestPutZeroAtExpiryWhenInMoneyCall(t *testing.T) {
	got := PutPrice(120, 100, 0, 0, 0)
	if got != 0 {
		t.Errorf("put = %g, want 0", got)
	}
}

func TestD2Relation(t *testing.T) {
	s, k, maturity, r, sigma := 100.0, 100.0, 1.0, 0.03, 0.2
	d1 := D1(s, k, maturity, r, sigma)
	d2 := D2(s, k, maturity, r, sigma)
	want := d1 - sigma*math.Sqrt(maturity)
	if math.Abs(d2-want) > 1e-12 {
		t.Errorf("d2 = %g, want %g", d2, want)
	}
}

func TestParityClosure(t *testing.T) {
	s, k, maturity, r, sigma := 100.0, 100.0, 1.0, 0.03, 0.2
	c := CallPrice(s, k, maturity, r, sigma)
	p := PutPrice(s, k, maturity, r, sigma)
	diff := (c - p) - (s - k*DiscountFactor(r, maturity))
	if math.Abs(diff) > 1e-9 {
		t.Errorf("parity gap = %g, want 0", diff)
	}
}

func TestATMOneYearCallNearApproximation(t *testing.T) {
	got := CallPrice(100, 100, 1, 0.03, 0.2)
	approx := 0.4 * 100 * 0.2 * math.Sqrt(1)
	if math.Abs(got-approx) > 2.5 {
		t.Errorf("call = %g, expected near %g", got, approx)
	}
}

func TestSigmaZeroCallLimit(t *testing.T) {
	got := CallPrice(100, 100, 1, 0.03, 0)
	want := math.Max(0, 100-100*math.Exp(-0.03))
	if math.Abs(got-want) > 1e-9 {
		t.Errorf("zero-vol call = %g, want %g", got, want)
	}
}

func TestImpliedVolatilityReproducesPrice(t *testing.T) {
	target := CallPrice(100, 100, 1, 0.03, 0.2)
	vol, err := ImpliedVolatility(target, 100, 100, 1, 0.03, "call", 1e-10, 120)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !RepriceVolatility(vol, target, 100, 100, 1, 0.03, "call") {
		t.Errorf("solved vol %g does not reproduce price", vol)
	}
}

func TestIsWithinPriceBounds(t *testing.T) {
	if !IsWithinPriceBounds(10, 100, 100, 1, 0.03, "call") {
		t.Errorf("ATM price should be within the vol range")
	}
}
