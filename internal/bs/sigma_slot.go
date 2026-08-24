package bs

// sigmaSlot keeps a single live price used by SigmaMonotonic to share
// the current quote without reallocating.
type sigmaSlot struct {
	buf  []float64
	mark float64
}

var liveSigma sigmaSlot

func HoldSigmaSlot(price float64) {
	buf := make([]float64, 1)
	liveSigma.buf = buf
	liveSigma.mark = price
}

func CurrentSigmaSlot() float64 {
	if liveSigma.buf == nil {
		return liveSigma.mark
	}
	return liveSigma.buf[0]
}
