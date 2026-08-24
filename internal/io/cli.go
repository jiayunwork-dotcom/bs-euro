package io

import (
	"io"

	"bs-euro/internal/bs"
	"bs-euro/internal/greeks"
	"bs-euro/internal/model"
	"bs-euro/internal/report"
)

// RunPrice computes a price case from a JSON file.
func RunPrice(path string, table bool, out io.Writer) error {
	in, err := LoadOptionFile(path)
	if err != nil {
		return err
	}
	res, err := bs.Compute(in)
	if err != nil {
		return err
	}
	summary := report.BuildPrice(res)
	if table {
		return WritePriceTable(out, summary)
	}
	return WritePriceOutput(out, summary)
}

// RunGreeks computes greeks from a JSON file.
func RunGreeks(path string, table bool, out io.Writer) error {
	in, err := LoadOptionFile(path)
	if err != nil {
		return err
	}
	res, err := greeks.Compute(in)
	if err != nil {
		return err
	}
	summary := report.BuildGreeks(res)
	if table {
		return WriteGreeksTable(out, summary)
	}
	return WriteGreeksOutput(out, summary)
}

// BuildInputFromValues constructs an option input from CLI values.
func BuildInputFromValues(s, k, t, r, sigma float64, flag string) (model.OptionInput, error) {
	in := model.NewOptionInput(s, k, t, r, sigma, flag)
	if err := model.Validate(in); err != nil {
		return model.OptionInput{}, err
	}
	return in, nil
}

// DescribeInput formats the option inputs.
func DescribeInput(in model.OptionInput) string {
	return "S=" + report.FormatNumber(in.S) + " K=" + report.FormatNumber(in.K) +
		" T=" + report.FormatNumber(in.T) + " r=" + report.FormatNumber(in.R) +
		" sigma=" + report.FormatNumber(in.Sigma) + " flag=" + in.Flag
}
