package main

import (
	"reflect"
	"testing"
)

var example string = `px{a<2006:qkq,m>2090:A,rfg}
pv{a>1716:R,A}
lnx{m>1548:A,A}
rfg{s<537:gd,x>2440:R,A}
qs{s>3448:A,lnx}
qkq{x<1416:A,crn}
crn{x>2662:A,R}
in{s<1351:px,qqz}
qqz{s>2770:qs,m<1801:hdj,R}
gd{a>3333:R,R}
hdj{m>838:A,pv}

{x=787,m=2655,a=1222,s=2876}
{x=1679,m=44,a=2067,s=496}
{x=2036,m=264,a=79,s=2244}
{x=2461,m=1339,a=466,s=291}
{x=2127,m=1623,a=2188,s=1013}`

func TestParsePart(t *testing.T) {
	res := ParsePart("{x=2036,m=264,a=79,s=2244}")
	expected := Part{X: 2036, M: 264, A: 79, S: 2244}
	if res != expected {
		t.Errorf("Expected %v got %v", expected, res)
	}
}

func TestParseWorkflow(t *testing.T) {
	res := ParseWorkflow("lnx{m>1548:A,A}")
	expected := Workflow{Name: "lnx", Commands: []Condition{{Prop: "m", GreaterThan: true, Number: 1548, Destination: "A"}}, Default: "A"}
	if !reflect.DeepEqual(res.Commands, expected.Commands) || res.Name != expected.Name || res.Default != expected.Default {
		t.Errorf("Expected %v got %v", expected, res)
	}
}

func TestTestConditions(t *testing.T) {
	wflw := ParseWorkflow("px{a<2006:qkq,m>2090:A,rfg}")
	res := wflw.TestConditions(Part{X: 2036, M: 264, A: 79, S: 2244})
	expected := "qkq"
	if res != expected {
		t.Errorf("Expected %v got %v", expected, res)
	}
}

func TestFullIntegrationWithExample(t *testing.T) {
	wflws, prts := ParseCommandsAndParts(example)
	acc, _ := AcceptAndRejectParts(wflws, prts)
	cuml := 0
	for _, prt := range acc {
		cuml = cuml + prt.GetTotal()
	}
	if cuml != 19114 {
		t.Errorf("Expected %v got %v", 19114, cuml)
	}
}
