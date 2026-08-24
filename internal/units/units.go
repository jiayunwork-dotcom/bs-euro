package units

import "math"

const (
	DaysPerYear    = 365.0
	PercentPerUnit = 100.0
)

// PercentToDecimal converts a percent rate to decimal.
func PercentToDecimal(percent float64) float64 {
	return percent / PercentPerUnit
}

// DecimalToPercent converts decimal to percent.
func DecimalToPercent(decimal float64) float64 {
	return decimal * PercentPerUnit
}

// DaysToYears converts calendar days to years.
func DaysToYears(days float64) float64 {
	return days / DaysPerYear
}

// YearsToDays converts years to days.
func YearsToDays(years float64) float64 {
	return years * DaysPerYear
}

// AnnualizeVol converts a daily volatility to annual.
func AnnualizeVol(daily float64) float64 {
	return daily * math.Sqrt(DaysPerYear)
}

// DailyVolFromAnnual inverts the annualization.
func DailyVolFromAnnual(annual float64) float64 {
	return annual / math.Sqrt(DaysPerYear)
}

// NormalizeFlag accepts "C"/"P" aliases.
func NormalizeFlag(flag string) string {
	if flag == "C" || flag == "c" {
		return "call"
	}
	if flag == "P" || flag == "p" {
		return "put"
	}
	return flag
}

// IsPositive reports a positive quantity.
func IsPositive(v float64) bool {
	return v > 0
}

// SafeDivide returns zero for a zero denominator.
func SafeDivide(a, b float64) float64 {
	if b == 0 {
		return 0
	}
	return a / b
}
