package report

import (
	"testing"

	"bs-euro/internal/model"
)

func TestBuildPriceSummary(t *testing.T) {
	res := model.PriceResult{Price: 8, D1: 0.1, D2: -0.1, Intrinsic: 0, Flag: "call"}
	s := BuildPrice(res)
	if s.Price != 8 || s.D1 != 0.1 || s.D2 != -0.1 {
		t.Errorf("summary fields wrong: %+v", s)
	}
}

func TestBuildGreeksSummary(t *testing.T) {
	res := model.GreeksResult{Delta: 0.6, Gamma: 0.05, Vega: 30, Flag: "call"}
	s := BuildGreeks(res)
	if s.Delta != 0.6 || s.Gamma != 0.05 || s.Vega != 30 {
		t.Errorf("summary fields wrong: %+v", s)
	}
}

func TestFormatNumber(t *testing.T) {
	if got := FormatNumber(0.03); got != "0.03" {
		t.Errorf("format = %q, want 0.03", got)
	}
}

func TestLabelFallback(t *testing.T) {
	if Label("") != "untitled" {
		t.Errorf("label fallback wrong")
	}
}
