package units

import (
	"math"
	"testing"
)

func TestPercentToDecimal(t *testing.T) {
	if math.Abs(PercentToDecimal(3)-0.03) > 1e-12 {
		t.Errorf("percent conversion wrong")
	}
}

func TestDaysToYears(t *testing.T) {
	if math.Abs(DaysToYears(365)-1) > 1e-12 {
		t.Errorf("days conversion wrong")
	}
}

func TestAnnualizeVol(t *testing.T) {
	if math.Abs(AnnualizeVol(0.01)-0.01*math.Sqrt(365)) > 1e-12 {
		t.Errorf("annualization wrong")
	}
}

func TestNormalizeFlag(t *testing.T) {
	if NormalizeFlag("c") != "call" || NormalizeFlag("P") != "put" {
		t.Errorf("flag normalization wrong")
	}
}

func TestSafeDivide(t *testing.T) {
	if SafeDivide(1, 0) != 0 {
		t.Errorf("safe divide wrong")
	}
}

func TestOptionInputFromPercent(t *testing.T) {
	in := OptionInputFromPercent(100, 100, 1, 3, 20, "call")
	if in.R != 0.03 || in.Sigma != 0.2 {
		t.Errorf("percent conversion wrong: %+v", in)
	}
}
