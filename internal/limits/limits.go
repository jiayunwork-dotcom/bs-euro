package limits

import (
	"bs-euro/internal/bs"
)

// ExpiryCall returns the intrinsic value at expiry.
func ExpiryCall(s, k float64) float64 {
	return bs.IntrinsicCall(s, k)
}

// ExpiryPut returns the intrinsic value at expiry.
func ExpiryPut(s, k float64) float64 {
	return bs.IntrinsicPut(s, k)
}

// DeepITMCall approaches S-K*e^{-rT} for large volatility.
func DeepITMCall(s, k, t, r float64) float64 {
	return bs.CallPrice(s, k, t, r, 1e-6)
}

// DeepOTMPut approaches zero for high spot.
func DeepOTMPut(s, k, t, r, sigma float64) float64 {
	return bs.PutPrice(s, k, t, r, sigma)
}

// UpperBoundCall reports C <= S.
func UpperBoundCall(c, s float64) bool {
	return c <= s+1e-12
}

// UpperBoundPut reports P <= K*e^{-rT}.
func UpperBoundPut(p, k, r, t float64) bool {
	return p <= k*bs.DiscountFactor(r, t)+1e-12
}

// NonNegativePrice reports price >= 0.
func NonNegativePrice(price float64) bool {
	return price >= 0
}

// TimeDecayReportsNegativeTheta verifies theta <= 0 for calls.
func TimeDecayReportsNegativeTheta(thetaCall float64) bool {
	return thetaCall <= 0
}

// BoundPrice clamps a price to non-negative.
func BoundPrice(price float64) float64 {
	if price < 0 {
		return 0
	}
	return price
}

// IsAtExpiry reports T=0.
func IsAtExpiry(t float64) bool {
	return t == 0
}

// Moneyness returns S/K.
func Moneyness(s, k float64) float64 {
	if k == 0 {
		return 0
	}
	return s / k
}
