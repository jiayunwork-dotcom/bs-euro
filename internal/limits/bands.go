package limits

import (
	"bs-euro/internal/bs"
)

// PriceBand is a low/high price envelope for an option.
type PriceBand struct {
	Low  float64
	High float64
}

// CallBand returns the lower and upper no-arbitrage bounds.
func CallBand(s, k, r, t float64) PriceBand {
	lower := s - k*bs.DiscountFactor(r, t)
	if lower < 0 {
		lower = 0
	}
	return PriceBand{Low: lower, High: s}
}

// PutBand returns the lower and upper no-arbitrage bounds.
func PutBand(s, k, r, t float64) PriceBand {
	lower := k*bs.DiscountFactor(r, t) - s
	if lower < 0 {
		lower = 0
	}
	return PriceBand{Low: lower, High: k * bs.DiscountFactor(r, t)}
}

// InsideBand reports whether a price is within the no-arbitrage envelope.
func InsideBand(price float64, band PriceBand) bool {
	return price >= band.Low-1e-12 && price <= band.High+1e-12
}

// VolatilityBand returns a price band across two volatilities.
func VolatilityBand(s, k, t, r, sigmaLow, sigmaHigh float64, flag string) PriceBand {
	p1 := bs.Price(s, k, t, r, sigmaLow, flag)
	p2 := bs.Price(s, k, t, r, sigmaHigh, flag)
	if p1 < p2 {
		return PriceBand{Low: p1, High: p2}
	}
	return PriceBand{Low: p2, High: p1}
}

// ExpiryPriceBand is the intrinsic-only envelope.
func ExpiryPriceBand(s, k float64, flag string) PriceBand {
	if flag == "put" {
		value := k - s
		if value < 0 {
			value = 0
		}
		return PriceBand{Low: value, High: value}
	}
	value := s - k
	if value < 0 {
		value = 0
	}
	return PriceBand{Low: value, High: value}
}

// Width returns the band width.
func (b PriceBand) Width() float64 {
	return b.High - b.Low
}

// Midpoint returns the band midpoint.
func (b PriceBand) Midpoint() float64 {
	return (b.Low + b.High) / 2
}

// IsEmpty reports a zero-width band.
func (b PriceBand) IsEmpty() bool {
	return b.Low == b.High
}
