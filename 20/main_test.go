package main

import (
	"testing"
)

// % -- flip flop
// starts off
// high pulse -> do nothing
// low pulse -> switch on/off
// -- becomes on -> send high
// -- becomes off -> send low

// & -- Conjunction
// starts remembering low for all inputs
// When a pulse is received,
// the conjunction module first updates its memory for that input to received
// Then, if it remembers high pulses for all inputs, it sends a low pulse;
// otherwise, it sends a high pulse.

// broadcaster -- what it is
// When it receives a pulse, it sends the same pulse to all of its destination modules

var exampleOne = `broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a`

var exampleTwo = `broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output`

func TestTheWholeDamnThing(t *testing.T) {
	// Sometimes, you just do the do
	ntwrk := ParseCommsModules(exampleOne)
	low, high := HitButton(ntwrk)

	if low != 8 {
		t.Errorf("Expected 8 got %v", low)
	}
	if high != 4 {
		t.Errorf("Expected 4 got %v", high)
	}
}

func TestTheWholeDamnThingTwo(t *testing.T) {
	// Sometimes, you just do the do
	ntwrk := ParseCommsModules(exampleTwo)
	low, high := 0, 0
	// 1000 times, 4250 low pulses and 2750 high pulses are sent. Multipl
	for i := 0; i < 1000; i = i + 1 {
		lowAdd, highAdd := HitButton(ntwrk)
		low = lowAdd + low
		high = highAdd + high
	}
	if low != 4250 {
		t.Errorf("Expected 4250 got %v", low)
	}
	if high != 2750 {
		t.Errorf("Expected 2750 got %v", high)
	}
}
