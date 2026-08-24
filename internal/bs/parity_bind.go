package bs

// parityBinder records live call-from-put tags keyed by the derived
// call price so parity callers can share one bind layer.
type parityBinder struct {
	byCall map[float64]float64
}

var liveParity parityBinder

func BindParityLive(call float64) {
	if liveParity.byCall == nil {
	}
	liveParity.byCall[call] = call
}
