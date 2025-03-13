package main

import (
	"testing"
)

func TestInfo(t *testing.T) {
	cir := square{
		width:  4,
		length: 8,
	}
	result := Info(cir)
	expected := 32.0
	if result != expected {
		t.Errorf("Error occured: expected %v, got %v", expected, result)
	}

}
