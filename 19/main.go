package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Condition struct {
	Prop        string
	GreaterThan bool
	Number      int
	Destination string
}

type Workflow struct {
	Name     string
	Commands []Condition
	Default  string
}

type Part struct {
	X int `json:"x"`
	M int `json:"m"`
	A int `json:"a"`
	S int `json:"s"`
}

func (prt Part) GetTotal() int {
	return prt.X + prt.A + prt.M + prt.S
}

func (wflw Workflow) TestConditions(prt Part) string {
	for _, com := range wflw.Commands {
		compare := 0
		if com.Prop == "x" {
			compare = prt.X
		}
		if com.Prop == "m" {
			compare = prt.M
		}
		if com.Prop == "a" {
			compare = prt.A
		}
		if com.Prop == "s" {
			compare = prt.S
		}
		if com.GreaterThan && com.Number < compare {
			return com.Destination
		}
		if !com.GreaterThan && com.Number > compare {
			return com.Destination
		}
	}
	return wflw.Default
}

func ParseWorkflow(input string) Workflow {
	startIndex := strings.Index(input, "{")
	name := input[:startIndex]
	allCommands := strings.Split(input[startIndex+1:len(input)-1], ",")
	cons := []Condition{}
	for _, com := range allCommands[:len(allCommands)-1] {
		comparisonIndex := strings.Index(com, ">")
		greaterthan := true
		if comparisonIndex == -1 {
			comparisonIndex = strings.Index(com, "<")
			greaterthan = false
		}
		destIndex := strings.Index(com, ":")
		prop := string(com[0])
		number, _ := strconv.Atoi(com[comparisonIndex+1 : destIndex])
		destination := com[destIndex+1:]
		cons = append(cons, Condition{
			Prop:        prop,
			GreaterThan: greaterthan,
			Number:      number,
			Destination: destination,
		})
	}

	return Workflow{
		Name:     name,
		Commands: cons,
		Default:  allCommands[len(allCommands)-1],
	}
}

func ParsePart(input string) Part {
	res := Part{}
	input = strings.ReplaceAll(input, "=", ":")
	input = strings.ReplaceAll(input, "x", "\"x\"")
	input = strings.ReplaceAll(input, "m", "\"m\"")
	input = strings.ReplaceAll(input, "a", "\"a\"")
	input = strings.ReplaceAll(input, "s", "\"s\"")
	json.Unmarshal([]byte(input), &res)
	return res
}

func ParseCommandsAndParts(input string) (map[string]Workflow, []Part) {
	wflws := map[string]Workflow{}
	prts := []Part{}
	flows := true
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			flows = false
			continue
		}
		if flows {
			flow := ParseWorkflow(line)
			wflws[flow.Name] = flow
		} else {
			prts = append(prts, ParsePart(line))
		}
	}
	return wflws, prts
}

func AcceptPart(flows map[string]Workflow, part Part, next string) bool {
	res := flows[next].TestConditions(part)
	if res == "A" {
		return true
	}
	if res == "R" {
		return false
	}
	return AcceptPart(flows, part, res)
}

func AcceptAndRejectParts(flows map[string]Workflow, parts []Part) ([]Part, []Part) {
	accepted := []Part{}
	rejected := []Part{}
	for _, part := range parts {
		if AcceptPart(flows, part, "in") {
			accepted = append(accepted, part)
		} else {
			rejected = append(rejected, part)
		}
	}
	return accepted, rejected
}

func main() {
	buf, _ := os.ReadFile("data.txt")
	stringput := string(buf)
	wflws, prts := ParseCommandsAndParts(stringput)
	acc, _ := AcceptAndRejectParts(wflws, prts)
	cuml := 0
	for _, prt := range acc {
		cuml = cuml + prt.GetTotal()
	}
	fmt.Printf("Part One: %d\n", cuml)
}
