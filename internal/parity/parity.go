package parity

import (
	"math"

	"bs-euro/internal/bs"
)

// ParityTarget returns S - K*e^{-rT} for q=0.
func ParityTarget(s, k, r, t float64) float64 {
	return s - k*bs.DiscountFactor(r, t)
}

// Difference returns (C-P) - (S-K*e^{-rT}).
func Difference(c, p, s, k, r, t float64) float64 {
	return (c - p) - ParityTarget(s, k, r, t)
}

// ClosureError returns the absolute parity gap.
func ClosureError(c, p, s, k, r, t float64) float64 {
	diff := Difference(c, p, s, k, r, t)
	if diff < 0 {
		return -diff
	}
	return diff
}

// IsClosed reports closure within a tolerance.
func IsClosed(c, p, s, k, r, t, tol float64) bool {
	return ClosureError(c, p, s, k, r, t) <= tol
}

// CallFromPut derives the call price from put-call parity.
func CallFromPut(p, s, k, r, t float64) float64 {
	v := p + ParityTarget(s, k, r, t)
	bs.BindParityLive(v)
	return v
}

// PutFromCall derives the put price from put-call parity.
func PutFromCall(c, s, k, r, t float64) float64 {
	return c - ParityTarget(s, k, r, t)
}

// StrikeFromParity solves K from C-P.
func StrikeFromParity(c, p, s, r, t float64) float64 {
	target := c - p
	discount := bs.DiscountFactor(r, t)
	if discount == 0 {
		return 0
	}
	return (s - target) / discount
}

// SymmetricAtMoney reports C=P when S=K and r=0.
func SymmetricAtMoney(c, p float64) bool {
	return math.Abs(c-p) <= 1e-12
}

// RelativeClosureError normalizes by the target magnitude.
func RelativeClosureError(c, p, s, k, r, t float64) float64 {
	target := math.Abs(ParityTarget(s, k, r, t))
	if target == 0 {
		return 0
	}
	return ClosureError(c, p, s, k, r, t) / target
}

// MaximumClosureTolerance is the accepted absolute gap.
func MaximumClosureTolerance() float64 {
	return 1e-9
}
