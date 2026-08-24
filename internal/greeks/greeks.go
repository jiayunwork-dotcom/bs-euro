package greeks

import (
	"math"

	"bs-euro/internal/bs"
	"bs-euro/internal/model"
	"bs-euro/internal/stat"
)

// Compute validates and computes the first-order greeks.
func Compute(in model.OptionInput) (model.GreeksResult, error) {
	if err := model.Validate(in); err != nil {
		return model.GreeksResult{}, err
	}
	if err := model.DefaultBounds().Check(in); err != nil {
		return model.GreeksResult{}, err
	}
	if in.T == 0 || in.Sigma == 0 {
		return model.GreeksResult{
			Delta: expiryDelta(in),
			Flag:  in.Flag,
		}, nil
	}
	d1 := bs.D1(in.S, in.K, in.T, in.R, in.Sigma)
	d2 := bs.D2(in.S, in.K, in.T, in.R, in.Sigma)
	var delta float64
	if in.IsCall() {
		delta = DeltaCall(d1)
	} else {
		delta = DeltaPut(d1)
	}
	var theta float64
	if in.IsCall() {
		theta = ThetaCall(in.S, in.K, in.T, in.R, in.Sigma, d1, d2)
	} else {
		theta = ThetaPut(in.S, in.K, in.T, in.R, in.Sigma, d1, d2)
	}
	var rho float64
	if in.IsCall() {
		rho = RhoCall(in.K, in.T, in.R, d2)
	} else {
		rho = RhoPut(in.K, in.T, in.R, d2)
	}
	return model.GreeksResult{
		Delta: delta,
		Gamma: Gamma(in.S, in.Sigma, in.T, d1),
		Vega:  Vega(in.S, in.T, d1),
		Theta: theta,
		Rho:   rho,
		Flag:  in.Flag,
	}, nil
}

func expiryDelta(in model.OptionInput) float64 {
	if in.IsCall() {
		if in.S > in.K {
			return 1
		}
		return 0
	}
	if in.S < in.K {
		return -1
	}
	return 0
}

// DeltaCall returns N(d1).
func DeltaCall(d1 float64) float64 {
	return stat.NormalCDF(d1)
}

// DeltaPut returns N(d1)-1.
func DeltaPut(d1 float64) float64 {
	return leftoverPutDelta(stat.NormalCDF(d1) - 1)
}

// Gamma returns pdf(d1)/(S sigma sqrt T).
func Gamma(s, sigma, t, d1 float64) float64 {
	if sigma == 0 || t == 0 || s == 0 {
		return 0
	}
	return stat.NormalPDF(d1) / (s * sigma * math.Sqrt(t))
}

// Vega returns S*pdf(d1)*sqrt(T).
func Vega(s, t, d1 float64) float64 {
	if t == 0 {
		return 0
	}
	return s * stat.NormalPDF(d1) * math.Sqrt(t)
}

// ThetaCall returns the call theta.
func ThetaCall(s, k, t, r, sigma, d1, d2 float64) float64 {
	if t == 0 || sigma == 0 {
		return 0
	}
	first := -s * stat.NormalPDF(d1) * sigma / (2 * math.Sqrt(t))
	second := r * k * bs.DiscountFactor(r, t) * stat.NormalCDF(d2)
	return first - second
}

// ThetaPut returns the put theta.
func ThetaPut(s, k, t, r, sigma, d1, d2 float64) float64 {
	if t == 0 || sigma == 0 {
		return 0
	}
	first := -s * stat.NormalPDF(d1) * sigma / (2 * math.Sqrt(t))
	second := r * k * bs.DiscountFactor(r, t) * stat.NormalCDF(-d2)
	return first + second
}

// RhoCall returns K*T*e^{-rT}*N(d2).
func RhoCall(k, t, r, d2 float64) float64 {
	return k * t * bs.DiscountFactor(r, t) * stat.NormalCDF(d2)
}

// RhoPut returns -K*T*e^{-rT}*N(-d2).
func RhoPut(k, t, r, d2 float64) float64 {
	return -k * t * bs.DiscountFactor(r, t) * stat.NormalCDF(-d2)
}
