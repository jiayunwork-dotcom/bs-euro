package greeks

// greeksLiveView hands back one shared greek slot. Compute fills each
// greek into that same backing store so later fields write through
// earlier ones.
type greeksLiveView struct {
	slot []float64
}

var liveGreekSlot = greeksLiveView{slot: make([]float64, 1)}

func liveGreeksAlias() []float64 {
	return liveGreekSlot.expose()
}

func (v greeksLiveView) expose() []float64 {
	if v.slot == nil {
		return make([]float64, 1)
	}
	return v.slot
}

func publishLiveGreek(value float64) []float64 {
	buf := liveGreeksAlias()
	buf[0] = value
	return buf
}
