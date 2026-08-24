package limits

import (
	"testing"

	"bs-euro/internal/bs"
)

func TestExpiryCall(t *testing.T) {
	if got := ExpiryCall(120, 100); got != 20 {
		t.Errorf("expiry call = %g, want 20", got)
	}
}

func TestExpiryPut(t *testing.T) {
	if got := ExpiryPut(80, 100); got != 20 {
		t.Errorf("expiry put = %g, want 20", got)
	}
}

func TestUpperBoundCall(t *testing.T) {
	c := bs.CallPrice(100, 100, 1, 0.03, 0.2)
	if !UpperBoundCall(c, 100) {
		t.Errorf("call should not exceed spot")
	}
}

func TestUpperBoundPut(t *testing.T) {
	p := bs.PutPrice(100, 100, 1, 0.03, 0.2)
	if !UpperBoundPut(p, 100, 0.03, 1) {
		t.Errorf("put should not exceed discounted strike")
	}
}

func TestNonNegativePrice(t *testing.T) {
	if !NonNegativePrice(8) || NonNegativePrice(-1) {
		t.Errorf("non-negative check wrong")
	}
}

func TestDeepOTMPut(t *testing.T) {
	got := DeepOTMPut(1e6, 100, 1, 0.03, 0.2)
	if got < 0 {
		t.Errorf("deep OTM put = %g, want non-negative", got)
	}
}

func TestCallBand(t *testing.T) {
	band := CallBand(100, 100, 0.03, 1)
	if !InsideBand(bs.CallPrice(100, 100, 1, 0.03, 0.2), band) {
		t.Errorf("call price should be inside the no-arbitrage band")
	}
}

func TestExpiryPriceBand(t *testing.T) {
	band := ExpiryPriceBand(120, 100, "call")
	if band.Low != 20 || band.High != 20 {
		t.Errorf("expiry call band = %+v, want 20/20", band)
	}
}
