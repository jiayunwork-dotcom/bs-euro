package io

import (
	"encoding/json"
	"io"
	"os"

	"bs-euro/internal/model"
)

// OptionRequest is the JSON shape accepted by the price and greeks endpoints.
type OptionRequest struct {
	S     *float64 `json:"s"`
	K     *float64 `json:"k"`
	T     *float64 `json:"t"`
	R     *float64 `json:"r"`
	Sigma *float64 `json:"sigma"`
	Flag  string   `json:"flag"`
	Label string   `json:"label,omitempty"`
}

// ToInput converts a validated request into a model input.
func (r OptionRequest) ToInput() (model.OptionInput, error) {
	fields := []struct {
		name  string
		value *float64
	}{
		{"s", r.S},
		{"k", r.K},
		{"t", r.T},
		{"r", r.R},
		{"sigma", r.Sigma},
	}
	var in model.OptionInput
	for _, f := range fields {
		if f.value == nil {
			return in, model.MissingField(f.name)
		}
	}
	in = model.NewOptionInput(*r.S, *r.K, *r.T, *r.R, *r.Sigma, r.Flag)
	in.Label = r.Label
	return in, nil
}

// DecodeOptionRequest reads and converts an option JSON body.
func DecodeOptionRequest(r io.Reader) (model.OptionInput, error) {
	var req OptionRequest
	if err := json.NewDecoder(r).Decode(&req); err != nil {
		return model.OptionInput{}, model.NewError(model.CodeInvalidJSON, err.Error())
	}
	return req.ToInput()
}

// LoadOptionFile reads an option request from disk.
func LoadOptionFile(path string) (model.OptionInput, error) {
	f, err := os.Open(path)
	if err != nil {
		return model.OptionInput{}, err
	}
	defer f.Close()
	return DecodeOptionRequest(f)
}

// ExamplePath returns the bundled ATM one-year example.
func ExamplePath() string {
	return "example/atm-1y.json"
}
