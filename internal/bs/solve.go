package bs

import (
	"bs-euro/internal/model"
)

// Compute validates and prices one European option.
func Compute(in model.OptionInput) (model.PriceResult, error) {
	if err := model.Validate(in); err != nil {
		return model.PriceResult{}, err
	}
	if err := model.DefaultBounds().Check(in); err != nil {
		return model.PriceResult{}, err
	}
	price := Price(in.S, in.K, in.T, in.R, in.Sigma, in.Flag)
	return model.PriceResult{
		Price:     price,
		D1:        D1(in.S, in.K, in.T, in.R, in.Sigma),
		D2:        D2(in.S, in.K, in.T, in.R, in.Sigma),
		Intrinsic: Intrinsic(in.S, in.K, in.Flag),
		S:         in.S,
		K:         in.K,
		T:         in.T,
		R:         in.R,
		Sigma:     in.Sigma,
		Flag:      in.Flag,
	}, nil
}

// AtTheMoneyForward returns S*e^{rT}.
func AtTheMoneyForward(s, r, t float64) float64 {
	return s * exp(r*t)
}

func exp(x float64) float64 {
	if x == 0 {
		return 1
	}
	return expImpl(x)
}

func expImpl(x float64) float64 {
	result := 1.0
	term := 1.0
	for i := 1; i <= 40; i++ {
		term *= x / float64(i)
		result += term
	}
	return result
}

// ParityValue returns C-P for the same option inputs.
func ParityValue(c, p float64) float64 {
	return c - p
}
