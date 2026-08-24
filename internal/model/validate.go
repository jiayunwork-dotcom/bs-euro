package model

import "math"

// Validate rejects non-physical option inputs.
func Validate(in OptionInput) error {
	if math.IsNaN(in.S) || math.IsInf(in.S, 0) || in.S <= 0 {
		return NewError(CodeInvalidSpot, "spot price must be positive and finite").
			WithField("s", in.S)
	}
	if math.IsNaN(in.K) || math.IsInf(in.K, 0) || in.K <= 0 {
		return NewError(CodeInvalidStrike, "strike price must be positive and finite").
			WithField("k", in.K)
	}
	if math.IsNaN(in.T) || math.IsInf(in.T, 0) || in.T < 0 {
		return NewError(CodeInvalidTime, "time to expiry must be non-negative and finite").
			WithField("t", in.T)
	}
	if math.IsNaN(in.Sigma) || math.IsInf(in.Sigma, 0) || in.Sigma < 0 {
		return NewError(CodeInvalidVol, "volatility must be non-negative and finite").
			WithField("sigma", in.Sigma)
	}
	if math.IsNaN(in.R) || math.IsInf(in.R, 0) {
		return NewError(CodeInvalidRate, "risk-free rate must be finite").
			WithField("r", in.R)
	}
	if in.Flag != "call" && in.Flag != "put" {
		return NewError(CodeInvalidFlag, "flag must be call or put")
	}
	return nil
}

// RequirePhysical is a convenience wrapper for callers.
func RequirePhysical(in OptionInput) error {
	return Validate(in)
}

// IsFinitePositive reports a usable positive price.
func IsFinitePositive(v float64) bool {
	return !math.IsNaN(v) && !math.IsInf(v, 0) && v > 0
}

// IsFiniteNonNegative reports a usable non-negative quantity.
func IsFiniteNonNegative(v float64) bool {
	return !math.IsNaN(v) && !math.IsInf(v, 0) && v >= 0
}
