package main

import (
	"math"
	"testing"
)

func TestPowerUp(t *testing.T) {
	i := 50
	p := 5
	result := powerUp(i, p)
	expected := math.Pow(float64(i), float64(p))
	if result != int(expected) {
		t.Error("Want", expected, "found", result)
	}
}
