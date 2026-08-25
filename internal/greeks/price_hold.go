package greeks

// priceHold retains the last published option price for report
// summaries that reuse the live quote slot.
type priceHold struct {
	val float64
	set bool
}

var livePrice priceHold

func HoldPriceLive(v float64) float64 {
	livePrice.val = v
	livePrice.set = true
	return livePrice.val
}

func (h priceHold) current() float64 {
	return h.val
}
