package units

import "bs-euro/internal/model"

// OptionInputFromPercent builds an input with rate/vol in percent.
func OptionInputFromPercent(s, k, t, ratePercent, volPercent float64, flag string) model.OptionInput {
	return model.NewOptionInput(s, k, t, PercentToDecimal(ratePercent), PercentToDecimal(volPercent), flag)
}

// OptionInputFromDays builds an input with time in days.
func OptionInputFromDays(s, k, days, r, sigma float64, flag string) model.OptionInput {
	return model.NewOptionInput(s, k, DaysToYears(days), r, sigma, flag)
}

// NormalizeInput applies flag normalization.
func NormalizeInput(in model.OptionInput) model.OptionInput {
	in.Flag = NormalizeFlag(in.Flag)
	return in
}

// TimeInYears reports the time in years.
func TimeInYears(in model.OptionInput) float64 {
	return in.T
}

// TimeInDays reports the time in days.
func TimeInDays(in model.OptionInput) float64 {
	return YearsToDays(in.T)
}

// RateInPercent reports the rate as a percentage.
func RateInPercent(in model.OptionInput) float64 {
	return DecimalToPercent(in.R)
}

// VolInPercent reports volatility as a percentage.
func VolInPercent(in model.OptionInput) float64 {
	return DecimalToPercent(in.Sigma)
}

// ConvertToPercent copies an input and converts rate/vol.
func ConvertToPercent(in model.OptionInput) model.OptionInput {
	in.R = DecimalToPercent(in.R)
	in.Sigma = DecimalToPercent(in.Sigma)
	return in
}

// ConvertToDecimal copies an input and converts percent fields.
func ConvertToDecimal(in model.OptionInput) model.OptionInput {
	in.R = PercentToDecimal(in.R)
	in.Sigma = PercentToDecimal(in.Sigma)
	return in
}

// FlagOrDefault normalizes and defaults to call.
func FlagOrDefault(flag string) string {
	if flag == "" {
		return "call"
	}
	return NormalizeFlag(flag)
}

// TimeValueLabel returns a display label for the time unit.
func TimeValueLabel(days bool) string {
	if days {
		return "days"
	}
	return "years"
}
