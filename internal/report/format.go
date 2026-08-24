package report

import (
	"fmt"
	"math"
	"strings"
)

// FormatNumber prints a float with enough digits.
func FormatNumber(v float64) string {
	if math.Abs(v) >= 1e5 || (math.Abs(v) < 1e-4 && v != 0) {
		return fmt.Sprintf("%.6e", v)
	}
	return fmt.Sprintf("%.8g", v)
}

// FormatCompact trims trailing zeros.
func FormatCompact(v float64) string {
	s := fmt.Sprintf("%.6g", v)
	return strings.TrimRight(strings.TrimRight(s, "0"), ".")
}

// FormatPercent prints a percentage.
func FormatPercent(v float64) string {
	return fmt.Sprintf("%.2f%%", v)
}

// FormatMoney prints a price.
func FormatMoney(v float64) string {
	return FormatNumber(v)
}

// FormatGreek prints a greek with four decimals.
func FormatGreek(v float64) string {
	return fmt.Sprintf("%.6g", v)
}

// PadRight right-pads a label.
func PadRight(s string, width int) string {
	if len(s) >= width {
		return s
	}
	return s + strings.Repeat(" ", width-len(s))
}
