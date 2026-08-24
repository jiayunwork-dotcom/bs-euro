package compare

// ScaleCase pairs a volatility multiplier with an expected price direction.
type ScaleCase struct {
	InputMultiplier float64
	PriceRises      bool
}

// VolCases returns the volatility scaling cases.
func VolCases() []ScaleCase {
	return []ScaleCase{
		{InputMultiplier: 1, PriceRises: true},
		{InputMultiplier: 2, PriceRises: true},
		{InputMultiplier: 0.5, PriceRises: false},
	}
}

// DirectionMatches verifies the monotonic direction.
func DirectionMatches(priceBefore, priceAfter float64, rises bool) bool {
	if rises {
		return priceAfter > priceBefore
	}
	return priceAfter < priceBefore
}

// RelativeError returns the unsigned relative difference.
func RelativeError(actual, expected float64) float64 {
	if expected == 0 {
		if actual == 0 {
			return 0
		}
		return 1
	}
	diff := actual - expected
	if diff < 0 {
		diff = -diff
	}
	if expected < 0 {
		expected = -expected
	}
	return diff / expected
}

// WithinRelative reports whether a relative error is acceptable.
func WithinRelative(actual, expected, maxErr float64) bool {
	return RelativeError(actual, expected) <= maxErr
}
