package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIPriceSuccess(t *testing.T) {
	mux := NewMux()
	body := `{"s":100,"k":100,"t":1,"r":0.03,"sigma":0.2,"flag":"call"}`
	req := httptest.NewRequest(http.MethodPost, "/api/price", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200; body=%s", rec.Code, rec.Body.String())
	}
	var out map[string]any
	if err := json.Unmarshal(rec.Body.Bytes(), &out); err != nil {
		t.Fatalf("unexpected JSON error: %v", err)
	}
	if out["price"] == nil || out["d1"] == nil {
		t.Errorf("response missing price or d1: %v", out)
	}
}

func TestAPIGreeksSuccess(t *testing.T) {
	mux := NewMux()
	body := `{"s":100,"k":100,"t":1,"r":0.03,"sigma":0.2,"flag":"put"}`
	req := httptest.NewRequest(http.MethodPost, "/api/greeks", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200; body=%s", rec.Code, rec.Body.String())
	}
	var out map[string]any
	if err := json.Unmarshal(rec.Body.Bytes(), &out); err != nil {
		t.Fatalf("unexpected JSON error: %v", err)
	}
	if out["delta"] == nil || out["gamma"] == nil {
		t.Errorf("response missing delta or gamma: %v", out)
	}
}

func TestAPIInvalidInputErrorBody(t *testing.T) {
	mux := NewMux()
	body := `{"s":-1,"k":100,"t":1,"r":0.03,"sigma":0.2,"flag":"call"}`
	req := httptest.NewRequest(http.MethodPost, "/api/price", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400; body=%s", rec.Code, rec.Body.String())
	}
	var out ErrorBody
	if err := json.Unmarshal(rec.Body.Bytes(), &out); err != nil {
		t.Fatalf("unexpected JSON error: %v", err)
	}
	if out.Error.Code != "invalid_spot" {
		t.Errorf("error code = %q, want invalid_spot", out.Error.Code)
	}
}

func TestAPIInvalidFlagErrorBody(t *testing.T) {
	mux := NewMux()
	body := `{"s":100,"k":100,"t":1,"r":0.03,"sigma":0.2,"flag":"straddle"}`
	req := httptest.NewRequest(http.MethodPost, "/api/greeks", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400; body=%s", rec.Code, rec.Body.String())
	}
	var out ErrorBody
	if err := json.Unmarshal(rec.Body.Bytes(), &out); err != nil {
		t.Fatalf("unexpected JSON error: %v", err)
	}
	if out.Error.Code != "invalid_flag" {
		t.Errorf("error code = %q, want invalid_flag", out.Error.Code)
	}
}

func TestAPIHealth(t *testing.T) {
	mux := NewMux()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
}
