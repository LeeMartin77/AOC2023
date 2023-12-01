package main

import "testing"

func TestAssert(t *testing.T) {
	res := myTestFunction()
	if res != 2 {
		t.Errorf("Expected 2 got %v", res)
	}
}
