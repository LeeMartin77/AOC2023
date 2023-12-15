package main

import "testing"

func TestAssert(t *testing.T) {
	exps := map[string]int{
		"rn=1": 30,
		"qp=3": 97,
		"cm=2": 47,
		"qp-":  14,
		"pc=4": 180,
		"ot=9": 9,
		"ab=5": 197,
		"pc-":  48,
		"pc=6": 214,
		"ot=7": 231,
	}

	for inp, exp := range exps {
		res := CalculateHash(inp)
		if res != exp {
			t.Errorf("%v - Expected: %v got: %v", inp, exp, res)
		}
	}
}
