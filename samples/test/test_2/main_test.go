package main

import (
	"testing"
)

var a int = 10
var b int = 33

func TestAdd(t *testing.T) {
	result := add(a, b)
	expected := a + b
	if result != expected {
		t.Error("Expected", expected, "found", result)
	}
}

func TestSubtract(t *testing.T) {
	result := subtract(a, b)
	expected := a - b
	if result != expected {
		t.Error("Expected", expected, "found", result)
	}
}

func TestDoMathematics(t *testing.T) {
	result1, result2 := doMathematics(a, b, subtract), doMathematics(a, b, add)
	expected1, expected2 := subtract(a, b), add(a, b)
	if result1 != expected1 || result2 != expected2 {
		t.Error("test failed")
	}
}
