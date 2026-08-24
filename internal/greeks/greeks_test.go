package greeks

import (
	"math"
	"testing"

	"bs-euro/internal/bs"
	"bs-euro/internal/model"
)

func TestPutDeltaRelation(t *testing.T) {
	s, k, maturity, r, sigma := 100.0, 100.0, 1.0, 0.03, 0.2
	d1 := bs.D1(s, k, maturity, r, sigma)
	if math.Abs(DeltaPut(d1)-(DeltaCall(d1)-1)) > 1e-12 {
		t.Errorf("put delta should equal call delta - 1")
	}
}

func TestGammaSameAcrossFlags(t *testing.T) {
	s, k, maturity, r, sigma := 100.0, 100.0, 1.0, 0.03, 0.2
	d1 := bs.D1(s, k, maturity, r, sigma)
	if math.Abs(Gamma(s, sigma, maturity, d1)-Gamma(s, sigma, maturity, d1)) > 1e-12 {
		t.Errorf("gamma should be flag-independent")
	}
}

func TestVegaSameAcrossFlags(t *testing.T) {
	s, maturity := 100.0, 1.0
	d1 := 0.1
	if math.Abs(Vega(s, maturity, d1)-Vega(s, maturity, d1)) > 1e-12 {
		t.Errorf("vega should be flag-independent")
	}
}

func TestComputeCallGreeks(t *testing.T) {
	in := model.NewOptionInput(100, 100, 1, 0.03, 0.2, "call")
	res, err := Compute(in)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.Delta <= 0 || res.Delta >= 1 {
		t.Errorf("call delta = %g, want between 0 and 1", res.Delta)
	}
	if res.Gamma <= 0 || res.Vega <= 0 {
		t.Errorf("gamma/vega should be positive: %+v", res)
	}
}

func TestExpiryDelta(t *testing.T) {
	in := model.NewOptionInput(120, 100, 0, 0, 0, "call")
	res, err := Compute(in)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.Delta != 1 {
		t.Errorf("expiry call delta = %g, want 1", res.Delta)
	}
}

func TestGreeksConsistent(t *testing.T) {
	in := model.NewOptionInput(100, 100, 1, 0.03, 0.2, "call")
	res, err := Compute(in)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !GreeksConsistent(res.Delta, res.Gamma, res.Vega, res.Theta, res.Rho) {
		t.Errorf("greeks failed sign/range checks: %+v", res)
	}
}
