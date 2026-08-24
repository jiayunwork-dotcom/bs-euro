package compare

import "testing"

func TestExpiryCallIntrinsicInvariant(t *testing.T) {
	if !ExpiryCallIntrinsic(120, 100) {
		t.Errorf("expiry call should equal intrinsic")
	}
}

func TestExpiryPutZeroInvariant(t *testing.T) {
	if !ExpiryPutZero(120, 100) {
		t.Errorf("expiry put should be zero when S>K")
	}
}

func TestSigmaMonotonicInvariant(t *testing.T) {
	if !SigmaMonotonic(100, 100, 1, 0.03, 0.15, 0.3) {
		t.Errorf("call and put should both rise with sigma")
	}
}

func TestRateMonotonicInvariant(t *testing.T) {
	if !RateMonotonic(100, 100, 1, 0.01, 0.05, 0.2) {
		t.Errorf("call should rise and put should fall with r")
	}
}

func TestAtTheMoneyEqualityInvariant(t *testing.T) {
	if !AtTheMoneyEquality(100, 100, 1, 0.2) {
		t.Errorf("ATM with r=0 should have C=P")
	}
}

func TestParityClosesInvariant(t *testing.T) {
	if !ParityCloses(100, 100, 1, 0.03, 0.2) {
		t.Errorf("put-call parity should close")
	}
}

func TestPutDeltaRelationInvariant(t *testing.T) {
	if !PutDeltaRelation(0.6, -0.4) {
		t.Errorf("put delta should be call delta - 1")
	}
}

func TestGammaAndVegaSame(t *testing.T) {
	if !GammaSameAcrossFlags(0.05, 0.05) || !VegaSameAcrossFlags(30, 30) {
		t.Errorf("gamma/vega should not depend on flag")
	}
}

func TestCallLowerBoundInvariant(t *testing.T) {
	if !CallLowerBound(8, 100, 100, 0.03, 1) {
		t.Errorf("call should respect lower bound")
	}
}
