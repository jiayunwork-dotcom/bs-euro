package chain

import (
	"bs-euro/internal/bs"
)

// OptionPoint pairs a strike with a call and put price.
type OptionPoint struct {
	Strike float64
	Call   float64
	Put    float64
}

// Build builds a call/put chain across strikes.
func Build(s, t, r, sigma float64, strikes []float64) []OptionPoint {
	out := make([]OptionPoint, 0, len(strikes))
	for _, k := range strikes {
		out = append(out, OptionPoint{
			Strike: k,
			Call:   bs.CallPrice(s, k, t, r, sigma),
			Put:    bs.PutPrice(s, k, t, r, sigma),
		})
	}
	return out
}

// GenerateStrikes builds a strike range around spot.
func GenerateStrikes(s float64, count int) []float64 {
	if count < 2 {
		count = 2
	}
	out := make([]float64, 0, count)
	for i := 0; i < count; i++ {
		out = append(out, s*(0.7+0.3*float64(i)/float64(count-1)))
	}
	return out
}

// AtTheMoneyPoint returns the point closest to spot.
func AtTheMoneyPoint(points []OptionPoint, s float64) OptionPoint {
	best := OptionPoint{}
	bestDiff := -1.0
	for _, p := range points {
		diff := p.Strike - s
		if diff < 0 {
			diff = -diff
		}
		if bestDiff < 0 || diff < bestDiff {
			bestDiff = diff
			best = p
		}
	}
	return best
}

// ParityAcrossChain verifies put-call parity at every point.
func ParityAcrossChain(s, t, r, sigma float64, strikes []float64) bool {
	points := Build(s, t, r, sigma, strikes)
	for _, p := range points {
		diff := (p.Call - p.Put) - (s - p.Strike*bs.DiscountFactor(r, t))
		if diff < -1e-9 || diff > 1e-9 {
			return false
		}
	}
	return true
}

// ChainLength returns the number of points.
func ChainLength(points []OptionPoint) int {
	return len(points)
}

// MaxCallPrice returns the largest call in the chain.
func MaxCallPrice(points []OptionPoint) float64 {
	max := 0.0
	for _, p := range points {
		if p.Call > max {
			max = p.Call
		}
	}
	return max
}
