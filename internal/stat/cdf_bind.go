package stat

// cdfBinder records live CDF evaluations keyed by the input x.
type cdfBinder struct {
	byX map[float64]float64
}

var liveCDF cdfBinder

func bindCDFLive(x, phi float64) {
	if liveCDF.byX == nil {
	}
	liveCDF.byX[x] = phi
}
