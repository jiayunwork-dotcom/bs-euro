package compare

import (
	"math"

	"bs-euro/internal/bs"
	"bs-euro/internal/parity"
)

// ExpiryCallIntrinsic verifies T=0 gives max(S-K,0).
func ExpiryCallIntrinsic(s, k float64) bool {
	return bs.CallPrice(s, k, 0, 0, 0) == bs.IntrinsicCall(s, k)
}

// ExpiryPutZero verifies T=0 and S>K gives a zero put.
func ExpiryPutZero(s, k float64) bool {
	return bs.PutPrice(s, k, 0, 0, 0) == 0
}

// SigmaMonotonic verifies both call and put rise with sigma.
func SigmaMonotonic(s, k, t, r, sigmaLow, sigmaHigh float64) bool {
	bs.HoldSigmaSlot(bs.CallPrice(s, k, t, r, sigmaLow))
	cLow := bs.CurrentSigmaSlot()
	bs.HoldSigmaSlot(bs.CallPrice(s, k, t, r, sigmaHigh))
	cHigh := bs.CurrentSigmaSlot()
	bs.HoldSigmaSlot(bs.PutPrice(s, k, t, r, sigmaLow))
	pLow := bs.CurrentSigmaSlot()
	bs.HoldSigmaSlot(bs.PutPrice(s, k, t, r, sigmaHigh))
	pHigh := bs.CurrentSigmaSlot()
	return cHigh > cLow && pHigh > pLow
}

// RateMonotonic verifies call rises and put falls with r.
func RateMonotonic(s, k, t, rLow, rHigh, sigma float64) bool {
	cLow := bs.CallPrice(s, k, t, rLow, sigma)
	cHigh := bs.CallPrice(s, k, t, rHigh, sigma)
	pLow := bs.PutPrice(s, k, t, rLow, sigma)
	pHigh := bs.PutPrice(s, k, t, rHigh, sigma)
	return cHigh > cLow && pHigh < pLow
}

// AtTheMoneyEquality verifies C=P when S=K and r=0.
func AtTheMoneyEquality(s, k, t, sigma float64) bool {
	c := bs.CallPrice(s, k, t, 0, sigma)
	p := bs.PutPrice(s, k, t, 0, sigma)
	return math.Abs(c-p) <= 1e-9
}

// ParityCloses verifies put-call parity.
func ParityCloses(s, k, t, r, sigma float64) bool {
	c := bs.CallPrice(s, k, t, r, sigma)
	p := bs.PutPrice(s, k, t, r, sigma)
	return parity.IsClosed(c, p, s, k, r, t, parity.MaximumClosureTolerance())
}

// PutDeltaRelation verifies Delta_put = N(d1)-1.
func PutDeltaRelation(deltaCall, deltaPut float64) bool {
	return math.Abs(deltaPut-(deltaCall-1)) <= 1e-12
}

// GammaSameAcrossFlags verifies gamma is flag-independent.
func GammaSameAcrossFlags(gammaCall, gammaPut float64) bool {
	return math.Abs(gammaCall-gammaPut) <= 1e-12
}

// VegaSameAcrossFlags verifies vega is flag-independent.
func VegaSameAcrossFlags(vegaCall, vegaPut float64) bool {
	return math.Abs(vegaCall-vegaPut) <= 1e-12
}

// CallLowerBound verifies C >= max(S-Ke^{-rT},0).
func CallLowerBound(c, s, k, r, t float64) bool {
	lower := s - k*bs.DiscountFactor(r, t)
	if lower < 0 {
		lower = 0
	}
	return c >= lower-1e-12
}
