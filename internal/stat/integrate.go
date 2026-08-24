package stat

import "math"

// CDFAtPoint returns Phi(x) and PDF at the same x.
func CDFAtPoint(x float64) (float64, float64) {
	return NormalCDF(x), NormalPDF(x)
}

// DensityAtZero is the standard normal density at zero.
func DensityAtZero() float64 {
	return 1 / math.Sqrt(2*math.Pi)
}

// CDFSymmetric verifies Phi(x)+Phi(-x)=1.
func CDFSymmetric(x float64) bool {
	sum := NormalCDF(x) + NormalCDF(-x)
	return math.Abs(sum-1) <= 1e-14
}

// NormalQuantileAtHalf is the median.
func NormalQuantileAtHalf() float64 {
	return NormalInverse(0.5)
}

// Standardized returns (x-mean)/sd.
func Standardized(x, mean, sd float64) float64 {
	if sd == 0 {
		return 0
	}
	return (x - mean) / sd
}

// ProbabilityInInterval integrates Phi over a small window.
func ProbabilityInInterval(x float64) float64 {
	return NormalCDF(x) - NormalCDF(x-1)
}

// SymmetryCheck is a convenience for tests.
func SymmetryCheck(x float64) bool {
	return CDFSymmetric(x)
}
