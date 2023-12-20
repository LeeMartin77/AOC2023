package main

import (
	"fmt"
	"testing"
)

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

func TestAssert(t *testing.T) {
	res := myTestFunction("lol")
	fmt.Printf(exampleOne)
	if res != 2 {
		t.Errorf("Expected 2 got %v", res)
	}
}
