package io

import (
	"encoding/json"
	"strings"
	"testing"

	"bs-euro/internal/model"
)

func TestLoadOptionFileExample(t *testing.T) {
	in, err := LoadOptionFile("../../example/atm-1y.json")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if in.S != 100 || in.K != 100 || in.T != 1 {
		t.Errorf("example input wrong: %+v", in)
	}
}

func TestDecodeOptionRequestMissingField(t *testing.T) {
	body := `{"s":100,"k":100,"t":1,"r":0.03}`
	_, err := DecodeOptionRequest(strings.NewReader(body))
	if err == nil {
		t.Fatalf("expected missing-field error, got nil")
	}
	if !model.IsCode(err, model.CodeMissingField) {
		t.Errorf("expected code %q, got %q", model.CodeMissingField, model.Code(err))
	}
}

func TestDecodeOptionRequest(t *testing.T) {
	body := `{"s":100,"k":100,"t":1,"r":0.03,"sigma":0.2,"flag":"put"}`
	in, err := DecodeOptionRequest(strings.NewReader(body))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !in.IsPut() {
		t.Errorf("flag should be put")
	}
}

func TestWriteJSONRoundTrip(t *testing.T) {
	var buf strings.Builder
	if err := WriteJSON(&buf, map[string]float64{"price": 8}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var out map[string]float64
	if err := json.Unmarshal([]byte(buf.String()), &out); err != nil {
		t.Fatalf("unexpected decode error: %v", err)
	}
	if out["price"] != 8 {
		t.Errorf("decoded price = %g, want 8", out["price"])
	}
}

func TestDescribeInput(t *testing.T) {
	in := model.NewOptionInput(100, 100, 1, 0.03, 0.2, "call")
	if DescribeInput(in) == "" {
		t.Errorf("description should not be empty")
	}
}
