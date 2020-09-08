package middleware

import (
	"math"
	"testing"
)

func TestPass(t *testing.T) {
	got := math.Abs(-1)
	expected := float64(1)

	if got != expected {
		t.Errorf("Abs(-1) = %v; want %v", got, expected)
	}
}

func TestFail(t *testing.T) {
	got := math.Abs(-1)
	expected := float64(2)

	if got != expected {
		t.Errorf("Abs(-1) = %v; want %v", got, expected)
	}
}
