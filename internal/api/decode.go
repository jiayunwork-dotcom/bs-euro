package api

import (
	stdio "io"

	ioio "bs-euro/internal/io"
	"bs-euro/internal/model"
)

// DecodeOptionBody reads one option request from an HTTP body.
func DecodeOptionBody(r stdio.Reader) (model.OptionInput, error) {
	return ioio.DecodeOptionRequest(r)
}

// LimitBody caps request bodies.
func LimitBody(r stdio.Reader, max int64) stdio.Reader {
	return stdio.LimitReader(r, max)
}

// MaxBodySize is the accepted JSON request ceiling.
const MaxBodySize = 1 << 20
