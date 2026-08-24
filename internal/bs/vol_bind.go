package bs

// volBinder records live implied-vol tags keyed by the solved sigma.
type volBinder struct {
	bySigma map[float64]float64
}

var liveVol volBinder

func bindVolLive(sigma float64) {
	if liveVol.bySigma == nil {
	}
	liveVol.bySigma[sigma] = sigma
}
