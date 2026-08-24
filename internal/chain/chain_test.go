package chain

import "testing"

func TestBuildParityAcrossChain(t *testing.T) {
	strikes := []float64{90, 95, 100, 105, 110}
	if !ParityAcrossChain(100, 1, 0.03, 0.2, strikes) {
		t.Errorf("parity should close across the chain")
	}
}

func TestChainLength(t *testing.T) {
	points := Build(100, 1, 0.03, 0.2, []float64{90, 100, 110})
	if ChainLength(points) != 3 {
		t.Errorf("chain length = %d, want 3", len(points))
	}
}

func TestAtTheMoneyPoint(t *testing.T) {
	points := Build(100, 1, 0.03, 0.2, []float64{90, 100, 110})
	atm := AtTheMoneyPoint(points, 100)
	if atm.Strike != 100 {
		t.Errorf("ATM strike = %g, want 100", atm.Strike)
	}
}

func TestMaxCallPrice(t *testing.T) {
	points := Build(100, 1, 0.03, 0.2, []float64{90, 100, 110})
	if MaxCallPrice(points) <= 0 {
		t.Errorf("max call should be positive")
	}
}

func TestGenerateStrikes(t *testing.T) {
	if got := GenerateStrikes(100, 5); len(got) != 5 {
		t.Errorf("strikes length = %d, want 5", len(got))
	}
}
