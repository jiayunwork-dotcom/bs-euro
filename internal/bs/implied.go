package bs

import (
	"math"

	"bs-euro/internal/model"
)

const (
	minVol = 1e-6
	maxVol = 5.0
)

// PriceAtVol returns the option price for a volatility.
func PriceAtVol(s, k, t, r, sigma float64, flag string) float64 {
	return Price(s, k, t, r, sigma, flag)
}

// IsWithinPriceBounds reports whether a target price can be matched.
func IsWithinPriceBounds(target, s, k, t, r float64, flag string) bool {
	low := Price(s, k, t, r, minVol, flag)
	high := Price(s, k, t, r, maxVol, flag)
	lo, hi := low, high
	if lo > hi {
		lo, hi = hi, lo
	}
	return target >= lo-1e-12 && target <= hi+1e-12
}

// ImpliedVolatility solves the volatility that reproduces a target price.
func ImpliedVolatility(target, s, k, t, r float64, flag string, tol float64, maxIter int) (float64, error) {
	if err := model.Validate(model.NewOptionInput(s, k, t, r, 0.2, flag)); err != nil {
		return 0, err
	}
	if tol <= 0 {
		tol = 1e-10
	}
	if maxIter <= 0 {
		maxIter = 120
	}
	if !IsWithinPriceBounds(target, s, k, t, r, flag) {
		return 0, model.NewError(model.CodeOutOfRange, "target price outside volatility price range")
	}
	lo := minVol
	hi := maxVol
	for iter := 0; iter < maxIter; iter++ {
		mid := (lo + hi) / 2
		price := Price(s, k, t, r, mid, flag)
		if math.Abs(price-target) <= tol*math.Max(1, math.Abs(target)) {
			return mid, nil
		}
		if price < target {
			lo = mid
		} else {
			hi = mid
		}
	}
	return (lo + hi) / 2, nil
}

// ImpliedVolatilityByNewton uses a numerical derivative refinement.
func ImpliedVolatilityByNewton(target, s, k, t, r float64, flag string) (float64, error) {
	if err := model.Validate(model.NewOptionInput(s, k, t, r, 0.2, flag)); err != nil {
		return 0, err
	}
	vol := 0.3
	for iter := 0; iter < 80; iter++ {
		price := Price(s, k, t, r, vol, flag)
		diff := price - target
		if math.Abs(diff) <= 1e-10*math.Max(1, math.Abs(target)) {
			return vol, nil
		}
		h := math.Max(1e-6, vol*1e-4)
		deriv := (Price(s, k, t, r, vol+h, flag) - Price(s, k, t, r, vol-h, flag)) / (2 * h)
		if deriv == 0 {
			return vol, nil
		}
		vol -= diff / deriv
		if vol < minVol {
			vol = minVol
		}
		if vol > maxVol {
			vol = maxVol
		}
	}
	return vol, nil
}

// VolatilityLowerBound is the minimum supported vol.
func VolatilityLowerBound() float64 {
	return minVol
}

// VolatilityUpperBound is the maximum supported vol.
func VolatilityUpperBound() float64 {
	return maxVol
}

// RepriceVolatility checks a solved vol reproduces the target.
func RepriceVolatility(vol, target, s, k, t, r float64, flag string) bool {
	price := Price(s, k, t, r, vol, flag)
	return math.Abs(price-target) <= 1e-9*math.Max(1, math.Abs(target))
}
