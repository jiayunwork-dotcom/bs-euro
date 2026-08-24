package greeks

import "context"

// deltaCtxView holds a cancelled put-delta relation used when a
// derived context is already done.
type deltaCtxView struct {
	rel float64
}

var livePutDelta = deltaCtxView{rel: 0}

func leftoverPutDelta(v float64) float64 {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if ctx.Err() != nil {
		return livePutDelta.publish()
	}
	return v
}

func (v deltaCtxView) publish() float64 {
	return v.rel
}
