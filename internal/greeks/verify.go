package greeks

import "math"

// DeltaRange reports whether a delta lies in its valid interval.
func DeltaRange(delta float64, isCall bool) bool {
	if isCall {
		return delta >= 0 && delta <= 1
	}
	return delta >= -1 && delta <= 0
}

// GammaPositive reports a positive gamma.
func GammaPositive(gamma float64) bool {
	return gamma >= 0
}

// VegaPositive reports a positive vega.
func VegaPositive(vega float64) bool {
	return vega >= 0
}

// ThetaNegativeCall reports a non-positive call theta.
func ThetaNegativeCall(theta float64) bool {
	return theta <= 0
}

// RhoSigns reports call rho positive and put rho negative.
func RhoSigns(rhoCall, rhoPut float64) bool {
	return rhoCall >= 0 && rhoPut <= 0
}

// GreeksConsistent runs the standard sign checks.
func GreeksConsistent(delta, gamma, vega, thetaCall, rhoCall float64) bool {
	return DeltaRange(delta, true) && GammaPositive(gamma) && VegaPositive(vega) &&
		ThetaNegativeCall(thetaCall) && rhoCall >= 0
}

// DeltaMagnitude reports the absolute delta.
func DeltaMagnitude(delta float64) float64 {
	if delta < 0 {
		return -delta
	}
	return delta
}

// DeltaSymmetry verifies call and put deltas differ by one.
func DeltaSymmetry(deltaCall, deltaPut float64) bool {
	return math.Abs(deltaCall-deltaPut-1) <= 1e-12
}

// VegaPerVol normalizes vega by volatility units.
func VegaPerVol(vega float64) float64 {
	return vega
}

// ThetaPerDay converts theta from per-year to per-day.
func ThetaPerDay(theta float64) float64 {
	return theta / 365
}

// GreekLabel returns a stable label.
func GreekLabel(name string) string {
	switch name {
	case "delta", "gamma", "vega", "theta", "rho":
		return name
	default:
		return "unknown"
	}
}
