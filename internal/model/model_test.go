package model

import "testing"

func TestValidateRejectsNonPositiveSpot(t *testing.T) {
	in := NewOptionInput(0, 100, 1, 0.03, 0.2, "call")
	err := Validate(in)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if !IsCode(err, CodeInvalidSpot) {
		t.Errorf("expected code %q, got %q", CodeInvalidSpot, Code(err))
	}
}

func TestValidateRejectsNegativeTime(t *testing.T) {
	in := NewOptionInput(100, 100, -1, 0.03, 0.2, "call")
	err := Validate(in)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if !IsCode(err, CodeInvalidTime) {
		t.Errorf("expected code %q, got %q", CodeInvalidTime, Code(err))
	}
}

func TestValidateRejectsBadFlag(t *testing.T) {
	in := NewOptionInput(100, 100, 1, 0.03, 0.2, "straddle")
	err := Validate(in)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if !IsCode(err, CodeInvalidFlag) {
		t.Errorf("expected code %q, got %q", CodeInvalidFlag, Code(err))
	}
}

func TestValidateAcceptsZeroVolatility(t *testing.T) {
	in := NewOptionInput(100, 100, 1, 0.03, 0, "put")
	if err := Validate(in); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestMissingField(t *testing.T) {
	if !IsCode(MissingField("s"), CodeMissingField) {
		t.Errorf("missing field code wrong")
	}
}

func TestBoundsRejectHugeSpot(t *testing.T) {
	in := NewOptionInput(2e9, 100, 1, 0.03, 0.2, "call")
	if err := DefaultBounds().Check(in); err == nil {
		t.Fatalf("expected bounds error, got nil")
	}
}
