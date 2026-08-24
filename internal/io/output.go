package io

import (
	"encoding/json"
	"io"

	"bs-euro/internal/model"
	"bs-euro/internal/report"
)

// WriteJSON encodes a value with indentation.
func WriteJSON(w io.Writer, v any) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(v)
}

// WritePriceOutput writes a price summary.
func WritePriceOutput(w io.Writer, s report.PriceSummary) error {
	return WriteJSON(w, s)
}

// WriteGreeksOutput writes a greeks summary.
func WriteGreeksOutput(w io.Writer, s report.GreeksSummary) error {
	return WriteJSON(w, s)
}

// WriteModelError writes a structured error body.
func WriteModelError(w io.Writer, err error) error {
	body := map[string]any{
		"error": map[string]any{
			"code":    model.Code(err),
			"message": err.Error(),
		},
	}
	return WriteJSON(w, body)
}

// WritePriceTable writes a human-readable price table.
func WritePriceTable(w io.Writer, s report.PriceSummary) error {
	t := &report.Table{}
	t.AddFloat("price", s.Price)
	t.AddFloat("d1", s.D1)
	t.AddFloat("d2", s.D2)
	t.AddFloat("intrinsic", s.Intrinsic)
	t.Add("flag", s.Flag)
	return t.Render(w)
}

// WriteGreeksTable writes a human-readable greeks table.
func WriteGreeksTable(w io.Writer, s report.GreeksSummary) error {
	t := &report.Table{}
	t.AddFloat("delta", s.Delta)
	t.AddFloat("gamma", s.Gamma)
	t.AddFloat("vega", s.Vega)
	t.AddFloat("theta", s.Theta)
	t.AddFloat("rho", s.Rho)
	t.Add("flag", s.Flag)
	return t.Render(w)
}
