package main

import (
	"reflect"
	"strings"
	"testing"
)

var example string = `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`

func TestParseLine(t *testing.T) {
	actual := ParseLine(strings.Split(example, "\n")[0])
	expected := Record{
		Entries:  []rune{'?', '?', '?', '.', '#', '#', '#'},
		Checksum: []int{1, 1, 3},
	}
	for i, rn := range expected.Entries {
		if actual.Entries[i] != rn {
			t.Errorf("Expected %v got %v", actual.Entries[i], rn)
		}
	}
	for i, rn := range expected.Checksum {
		if actual.Checksum[i] != rn {
			t.Errorf("Expected %v got %v", actual.Checksum[i], rn)
		}
	}
}

func TestFulfillsChecksum(t *testing.T) {
	if !FulfilsChecksum([]rune("#.#.###"), []int{1, 1, 3}) {

		t.Error("Expected true got false")
	}
	if FulfilsChecksum([]rune("##..###"), []int{1, 1, 3}) {
		t.Error("Expected false got true")
	}
}

func TestUnfold(t *testing.T) {
	orig := ParseLine("???.### 1,1,3")
	expect := ParseLine("???.###????.###????.###????.###????.### 1,1,3,1,1,3,1,1,3,1,1,3,1,1,3")
	mod := Unfold(orig)
	if !reflect.DeepEqual(expect.Checksum, mod.Checksum) {
		t.Errorf("Exp: %v \nGot: %v", expect.Checksum, mod.Checksum)
	}
	if !reflect.DeepEqual(expect.Entries, mod.Entries) {
		t.Errorf("Exp: %v \nGot: %v", expect.Entries, mod.Entries)
	}
}

func TestPossibleConfigurations(t *testing.T) {
	res := PossibleConfigurations(ParseLine(strings.Split(example, "\n")[0]))
	if res != 1 {
		t.Errorf("Expected 1 got %v", res)
	}
	res = PossibleConfigurations(ParseLine(strings.Split(example, "\n")[1]))
	if res != 4 {
		t.Errorf("Expected 4 got %v", res)
	}
	res = PossibleConfigurations(ParseLine(strings.Split(example, "\n")[5]))
	if res != 10 {
		t.Errorf("Expected 10 got %v", res)
	}
}
func TestPossibleUnfoldedConfigurations(t *testing.T) {
	res := PossibleConfigurations(Unfold(ParseLine(strings.Split(example, "\n")[0])))
	if res != 1 {
		t.Errorf("Expected 1 got %v", res)
	}
	res = PossibleConfigurations(Unfold(ParseLine(strings.Split(example, "\n")[1])))
	if res != 16384 {
		t.Errorf("Expected 16384 got %v", res)
	}
	res = PossibleConfigurations(Unfold(ParseLine(strings.Split(example, "\n")[5])))
	if res != 506250 {
		t.Errorf("Expected 506250 got %v", res)
	}
}
