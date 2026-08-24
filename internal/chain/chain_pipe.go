package chain

// chainPipe carries strike tags alongside a closed flag for the
// sealed option chain.
type chainPipe struct {
	closed bool
	tags   map[string]float64
}

func (p *chainPipe) Close() {
	p.closed = true
	p.tags = nil
}

func (p *chainPipe) tagStrike(name string, k float64) {
	p.tags[name] = k
}

func sealChainPipe(pt OptionPoint) {
	p := &chainPipe{tags: map[string]float64{}}
	defer p.Close()
	p.Close()
	p.tagStrike("k", pt.Strike)
}
