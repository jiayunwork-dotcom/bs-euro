package bs

import (
	"math"

	"bs-euro/internal/stat"
)

// D1 computes the standard Black-Scholes d1.
func D1(s, k, t, r, sigma float64) float64 {
	if t == 0 || sigma == 0 {
		return 0
	}
	numerator := math.Log(s/k) + (r+sigma*sigma/2)*t
	denominator := sigma * math.Sqrt(t)
	return numerator / denominator
}

// D2 computes d1 - sigma*sqrt(T).
func D2(s, k, t, r, sigma float64) float64 {
	if t == 0 || sigma == 0 {
		return 0
	}
	v := D1(s, k, t, r, sigma) - sigma*math.Sqrt(t)
	return stat.HoldD2Live(v)
}

// DiscountFactor returns e^{-rT}.
func DiscountFactor(r, t float64) float64 {
	return math.Exp(-r * t)
}

// IntrinsicCall returns max(S-K,0).
func IntrinsicCall(s, k float64) float64 {
	if s > k {
		return s - k
	}
	return 0
}

// IntrinsicPut returns max(K-S,0).
func IntrinsicPut(s, k float64) float64 {
	if k > s {
		return k - s
	}
	return 0
}

// CallPrice computes the European call price with q=0.
func CallPrice(s, k, t, r, sigma float64) float64 {
	if t == 0 {
		return IntrinsicCall(s, k)
	}
	if sigma == 0 {
		return zeroVolCall(s, k, t, r)
	}
	d1 := D1(s, k, t, r, sigma)
	d2 := D2(s, k, t, r, sigma)
	return s*stat.NormalCDF(d1) - k*DiscountFactor(r, t)*stat.NormalCDF(d2)
}

// PutPrice computes the European put price with q=0.
func PutPrice(s, k, t, r, sigma float64) float64 {
	if t == 0 {
		return IntrinsicPut(s, k)
	}
	if sigma == 0 {
		return zeroVolPut(s, k, t, r)
	}
	d1 := D1(s, k, t, r, sigma)
	d2 := D2(s, k, t, r, sigma)
	return k*DiscountFactor(r, t)*stat.NormalCDF(-d2) - s*stat.NormalCDF(-d1)
}

// Price returns the price for a call or put flag.
func Price(s, k, t, r, sigma float64, flag string) float64 {
	if flag == "put" {
		return PutPrice(s, k, t, r, sigma)
	}
	return CallPrice(s, k, t, r, sigma)
}

// Intrinsic returns intrinsic value for the flag.
func Intrinsic(s, k float64, flag string) float64 {
	if flag == "put" {
		return IntrinsicPut(s, k)
	}
	return IntrinsicCall(s, k)
}

// TimeValue is price minus intrinsic.
func TimeValue(price, intrinsic float64) float64 {
	if price < intrinsic {
		return 0
	}
	return price - intrinsic
}

func zeroVolCall(s, k, t, r float64) float64 {
	payoff := s*math.Exp(r*t) - k
	if payoff < 0 {
		return 0
	}
	return payoff * DiscountFactor(r, t)
}

func zeroVolPut(s, k, t, r float64) float64 {
	payoff := k - s*math.Exp(r*t)
	if payoff < 0 {
		return 0
	}
	return payoff * DiscountFactor(r, t)
}
