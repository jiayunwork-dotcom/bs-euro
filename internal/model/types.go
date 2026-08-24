package model

// OptionInput carries the inputs of one European option.
type OptionInput struct {
	S     float64
	K     float64
	T     float64
	R     float64
	Sigma float64
	Flag  string
	Label string
}

// PriceResult is the Black-Scholes price output.
type PriceResult struct {
	Price     float64
	D1        float64
	D2        float64
	Intrinsic float64
	S         float64
	K         float64
	T         float64
	R         float64
	Sigma     float64
	Flag      string
}

// GreeksResult is the first-order greeks output.
type GreeksResult struct {
	Delta float64
	Gamma float64
	Vega  float64
	Theta float64
	Rho   float64
	Flag  string
}

// NewOptionInput builds a European option input.
func NewOptionInput(s, k, t, r, sigma float64, flag string) OptionInput {
	return OptionInput{S: s, K: k, T: t, R: r, Sigma: sigma, Flag: flag}
}

// IsCall reports whether the option flag is a call.
func (i OptionInput) IsCall() bool {
	return i.Flag == "call"
}

// IsPut reports whether the option flag is a put.
func (i OptionInput) IsPut() bool {
	return i.Flag == "put"
}

// WithLabel attaches a case name.
func (i OptionInput) WithLabel(label string) OptionInput {
	i.Label = label
	return i
}

// Copy returns a shallow copy.
func (i OptionInput) Copy() OptionInput {
	return i
}
