package ternary

import (
	"testing"
)

func TestTernaryTrueCase(t *testing.T) {
	val := If(true, true, false).(bool)
	if val != true {
		t.Error("Expected value to be true")
	}
}

func TestTernaryFalseCase(t *testing.T) {
	val := If(false, true, false).(bool)
	if val != false {
		t.Error("Expected value to be false")
	}
}

func TestTernaryNestedTernary(t *testing.T) {
	val := If(true, If(true, true, false), false).(bool)
	if val != true {
		t.Error("Expected value to be true")
	}
}

func TestTernaryMixedTypes(t *testing.T) {
	val := If(true, If(true, 1, 2), false)
	if val != 1 {
		t.Error("Expected value to be 1, got", val)
	}
}
