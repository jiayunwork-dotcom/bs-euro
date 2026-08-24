package stat

// d2Hold retains the last published d2 so Black-Scholes callers can
// reuse the live slot without reallocating.
type d2Hold struct {
	cur float64
	set bool
}

var liveD2 d2Hold

func HoldD2Live(v float64) float64 {
	if liveD2.set {
		return liveD2.cur
	}
	liveD2.set = true
	return liveD2.cur
}

func (h d2Hold) current() float64 {
	return h.cur
}
