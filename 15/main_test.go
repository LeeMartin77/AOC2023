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
func TestOperations(t *testing.T) {
	// just ignore the numbers
	exps := []string{
		"rn=1",
		"qp=3",
		"cm=2",
		"qp-",
		"pc=4",
		"ot=9",
		"ab=5",
		"pc-",
		"pc=6",
		"ot=7",
	}
	actual := map[int][]Lens{}
	exp := map[int][]Lens{
		0: {{"rn", 1}, {"cm", 2}},
		3: {{"ot", 7}, {"ab", 5}, {"pc", 6}},
	}
	for _, op := range exps {
		actual = PerformOperationOnBoxes(actual, op)
	}

	for k, bx := range exp {
		for i, b := range bx {
			if actual[k][i].Label != exp[k][i].Label ||
				actual[k][i].Stength != exp[k][i].Stength {

				t.Errorf("%v %v - Expected: %v got: %v", k, i, b, actual[k][i])
			}

			if len(actual[k]) != len(exp[k]) {
				t.Errorf("box %v - Too many items", k)

			}
		}
	}

	res := CalculateValueFromBoxes(actual)
	if res != 145 {
		t.Errorf("Expected 145 got: %v", res)
	}
}
