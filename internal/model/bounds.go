package model

// Bounds keeps option quantities inside a defensive operating envelope.
type Bounds struct {
	MaxSpot   float64
	MaxStrike float64
	MaxTime   float64
	MaxRate   float64
	MaxVol    float64
}

// DefaultBounds returns the standard defensive limits.
func DefaultBounds() Bounds {
	return Bounds{
		MaxSpot:   1e9,
		MaxStrike: 1e9,
		MaxTime:   200,
		MaxRate:   10,
		MaxVol:    10,
	}
}

// Check returns the first bound violation, if any.
func (b Bounds) Check(in OptionInput) error {
	if in.S > b.MaxSpot {
		return NewError(CodeOutOfRange, "spot price exceeds the supported bound").
			WithField("s", in.S)
	}
	if in.K > b.MaxStrike {
		return NewError(CodeOutOfRange, "strike price exceeds the supported bound").
			WithField("k", in.K)
	}
	if in.T > b.MaxTime {
		return NewError(CodeOutOfRange, "time to expiry exceeds the supported bound").
			WithField("t", in.T)
	}
	if in.R > b.MaxRate || in.R < -b.MaxRate {
		return NewError(CodeOutOfRange, "rate exceeds the supported bound").
			WithField("r", in.R)
	}
	if in.Sigma > b.MaxVol {
		return NewError(CodeOutOfRange, "volatility exceeds the supported bound").
			WithField("sigma", in.Sigma)
	}
	return nil
}

// Summary returns a human-readable envelope.
func (b Bounds) Summary() string {
	return "S<=1e9 K<=1e9 T<=200 |r|<=10 sigma<=10"
}
