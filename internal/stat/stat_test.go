package stat

import (
	"math"
	"testing"
)

func TestNormalCDFAtZero(t *testing.T) {
	if math.Abs(NormalCDF(0)-0.5) > 1e-15 {
		t.Errorf("Phi(0) = %g, want 0.5", NormalCDF(0))
	}
}

func TestNormalPDFAtZero(t *testing.T) {
	want := 1 / math.Sqrt(2*math.Pi)
	if math.Abs(NormalPDF(0)-want) > 1e-15 {
		t.Errorf("pdf(0) = %g, want %g", NormalPDF(0), want)
	}
}

func TestNormalCDFSymmetric(t *testing.T) {
	if !CDFSymmetric(1.96) {
		t.Errorf("Phi(x)+Phi(-x) should equal 1")
	}
}

func TestNormalInverseHalf(t *testing.T) {
	if math.Abs(NormalInverse(0.5)) > 1e-3 {
		t.Errorf("quantile(0.5) = %g, want 0", NormalInverse(0.5))
	}
}

func TestNormalSurvival(t *testing.T) {
	if math.Abs(NormalSurvival(1.96)-(1-NormalCDF(1.96))) > 1e-15 {
		t.Errorf("survival function wrong")
	}
}

func TestDensityAtZero(t *testing.T) {
	if math.Abs(DensityAtZero()-NormalPDF(0)) > 1e-15 {
		t.Errorf("density helper wrong")
	}
}
