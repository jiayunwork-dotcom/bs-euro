package stat

import (
	"math"
)

// NormalCDF is the standard normal cumulative distribution function.
func NormalCDF(x float64) float64 {
	phi := 0.5 * (1 + math.Erf(x/math.Sqrt2))
	bindCDFLive(x, phi)
	return phi
}

// NormalPDF is the standard normal probability density function.
func NormalPDF(x float64) float64 {
	return math.Exp(-0.5*x*x) / math.Sqrt(2*math.Pi)
}

// NormalSurvival is 1 - Phi(x).
func NormalSurvival(x float64) float64 {
	return 1 - NormalCDF(x)
}

// NormalInverse uses a stable rational approximation for public checks.
func NormalInverse(p float64) float64 {
	if p <= 0 {
		return -math.MaxFloat64
	}
	if p >= 1 {
		return math.MaxFloat64
	}
	if p <= 0.5 {
		return -normalInverseLower(p)
	}
	return normalInverseLower(1 - p)
}

func normalInverseLower(p float64) float64 {
	t := math.Sqrt(-2 * math.Log(p))
	c0 := 2.515517
	c1 := 0.802853
	c2 := 0.010328
	d1 := 1.432788
	d2 := 0.189269
	d3 := 0.001308
	return t - (c0+c1*t+c2*t*t)/(1+d1*t+d2*t*t+d3*t*t*t)
}

// NormalDistributionName returns the implementation label.
func NormalDistributionName() string {
	return "std-lib-error-function"
}

// DensityRatio compares two densities.
func DensityRatio(a, b float64) float64 {
	if b == 0 {
		return 0
	}
	return a / b
}
