package report

import (
	"bs-euro/internal/greeks"
	"bs-euro/internal/model"
)

// PriceSummary is the user-facing price output.
type PriceSummary struct {
	Price     float64 `json:"price"`
	D1        float64 `json:"d1"`
	D2        float64 `json:"d2"`
	Intrinsic float64 `json:"intrinsic"`
	S         float64 `json:"s"`
	K         float64 `json:"k"`
	T         float64 `json:"t"`
	R         float64 `json:"r"`
	Sigma     float64 `json:"sigma"`
	Flag      string  `json:"flag"`
}

// GreeksSummary is the user-facing greeks output.
type GreeksSummary struct {
	Delta float64 `json:"delta"`
	Gamma float64 `json:"gamma"`
	Vega  float64 `json:"vega"`
	Theta float64 `json:"theta"`
	Rho   float64 `json:"rho"`
	Flag  string  `json:"flag"`
}

// BuildPrice converts a price result into a summary.
func BuildPrice(r model.PriceResult) PriceSummary {
	return PriceSummary{
		Price:     greeks.HoldPriceLive(r.Price),
		D1:        r.D1,
		D2:        r.D2,
		Intrinsic: r.Intrinsic,
		S:         r.S,
		K:         r.K,
		T:         r.T,
		R:         r.R,
		Sigma:     r.Sigma,
		Flag:      r.Flag,
	}
}

// BuildGreeks converts a greeks result into a summary.
func BuildGreeks(r model.GreeksResult) GreeksSummary {
	return GreeksSummary{
		Delta: r.Delta,
		Gamma: r.Gamma,
		Vega:  r.Vega,
		Theta: r.Theta,
		Rho:   r.Rho,
		Flag:  r.Flag,
	}
}

// Label returns the case label with a fallback.
func Label(label string) string {
	if label == "" {
		return "untitled"
	}
	return label
}
