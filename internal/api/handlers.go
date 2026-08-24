package api

import (
	"net/http"

	"bs-euro/internal/bs"
	"bs-euro/internal/greeks"
	"bs-euro/internal/report"
)

// HandlePrice computes the Black-Scholes price for POST /api/price.
func HandlePrice(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		methodNotAllowed(w)
		return
	}
	in, err := DecodeOptionBody(LimitBody(r.Body, MaxBodySize))
	if err != nil {
		writeError(w, err)
		return
	}
	res, err := bs.Compute(in)
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, report.BuildPrice(res))
}

// HandleGreeks computes first-order greeks for POST /api/greeks.
func HandleGreeks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		methodNotAllowed(w)
		return
	}
	in, err := DecodeOptionBody(LimitBody(r.Body, MaxBodySize))
	if err != nil {
		writeError(w, err)
		return
	}
	res, err := greeks.Compute(in)
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, report.BuildGreeks(res))
}

// HandleHealth reports service health.
func HandleHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		methodNotAllowed(w)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{
		"status": "ok",
		"name":   "bs-euro",
	})
}

// HandleRoot describes the option-pricing endpoints.
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		methodNotAllowed(w)
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"service": "bs-euro",
		"endpoints": []string{
			"POST /api/price",
			"POST /api/greeks",
		},
	})
}

// NewRoutes registers every public route.
func NewRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", HandleRoot)
	mux.HandleFunc("/health", HandleHealth)
	mux.HandleFunc("/api/price", HandlePrice)
	mux.HandleFunc("/api/greeks", HandleGreeks)
}
