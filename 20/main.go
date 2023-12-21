package main

import (
	"fmt"
	"os"
	"strings"
)

type ModuleType int64

type Broadcaster struct {
	Name      string
	SendsTo   []string
	HighQueue []bool
}

type FlipFlop struct {
	Name      string
	SendsTo   []string
	State     bool
	HighQueue []bool
}

type Conjunction struct {
	Name    string
	Memory  map[string]bool
	SendsTo []string
}

type CommModule interface {
	ReceivePulse(pulse bool, sender string)
	SendPulseTo() (bool, []string)
}

func (cms *Broadcaster) ReceivePulse(pulse bool, sender string) {
	cms.HighQueue = append(cms.HighQueue, pulse)
}

func (cms *Broadcaster) SendPulseTo() (bool, []string) {
	pls := false
	if len(cms.HighQueue) == 0 {
		return false, []string{}
	}
	pls, cms.HighQueue = cms.HighQueue[0], cms.HighQueue[1:]
	return pls, cms.SendsTo
}

// dule receives a low pulse, it flips between on and off.
// If it was off, it turns on and sends a high pulse.
// If it was on, it turns off and sends a low pulse.
func (cms *FlipFlop) ReceivePulse(pulse bool, sender string) {
	if pulse {
		//noop
		return
	}
	cms.State = !cms.State
	cms.HighQueue = append(cms.HighQueue, cms.State)
}

func (cms *FlipFlop) SendPulseTo() (bool, []string) {
	pls := false
	if len(cms.HighQueue) == 0 {
		return false, []string{}
	}
	pls, cms.HighQueue = cms.HighQueue[0], cms.HighQueue[1:]
	return pls, cms.SendsTo
}

func (cms *Conjunction) ReceivePulse(pulse bool, sender string) {
	cms.Memory[sender] = pulse
}

func (cms *Conjunction) SendPulseTo() (bool, []string) {
	for _, high := range cms.Memory {
		if !high {
			return true, cms.SendsTo
		}
	}
	return false, cms.SendsTo
}

func ParseCommsModules(input string) map[string]CommModule {
	// we need to make a "inverse" of it because of the memory thing
	res := map[string]CommModule{}
	receivesFrom := map[string][]string{}
	for _, str := range strings.Split(input, "\n") {
		firstSpaceIndex := strings.Index(str, " ")
		name := str[:firstSpaceIndex]
		if name[0] != 'b' {
			name = name[1:]
		}
		targets := strings.Split(str[firstSpaceIndex+1+3:], ", ")
		for _, tgt := range targets {
			_, ok := receivesFrom[tgt]
			if !ok {
				receivesFrom[tgt] = []string{name}
			} else {
				receivesFrom[tgt] = append(receivesFrom[tgt], name)
			}
		}
	}
	for _, str := range strings.Split(input, "\n") {
		firstSpaceIndex := strings.Index(str, " ")
		name := str[:firstSpaceIndex]
		typ := name[0]
		if name[0] != 'b' {
			name = name[1:]
		}
		targets := strings.Split(str[firstSpaceIndex+1+3:], ", ")
		if typ == 'b' {
			// broadcaster
			res[name] = &Broadcaster{
				Name:    name,
				SendsTo: targets,
			}
		}
		if typ == '&' {
			// conjunction
			mem := map[string]bool{}
			for _, k := range receivesFrom[name] {
				mem[k] = false
			}
			res[name] = &Conjunction{
				Name:    name,
				SendsTo: targets,
				Memory:  mem,
			}
		}
		if typ == '%' {
			// flipflop
			res[name] = &FlipFlop{
				Name:    name,
				SendsTo: targets,
				State:   false,
			}
		}
	}
	return res
}

type Pulse struct {
	Target string
	Sender string
	High   bool
}

// lowpulses, highpulses
func HitButton(network map[string]CommModule) (int, int) {
	queue := []string{"broadcaster"}
	lowCount := 1 //start at 1 from button
	network["broadcaster"].ReceivePulse(false, "button")
	highCount := 0
	query := ""
	for len(queue) > 0 {
		query, queue = queue[0], queue[1:]
		if network[query] == nil {
			continue
		}
		high, tgts := network[query].SendPulseTo()
		for _, tgt := range tgts {
			if high {
				highCount = highCount + 1
			} else {
				lowCount = lowCount + 1
			}
			queue = append(queue, tgt)
			if network[tgt] != nil {
				network[tgt].ReceivePulse(high, query)
			}
		}
	}
	return lowCount, highCount
}

func main() {
	buf, _ := os.ReadFile("data.txt")
	input := string(buf)

	ntwrk := ParseCommsModules(input)
	low, high := 0, 0
	for i := 0; i < 1000; i = i + 1 {
		lowAdd, highAdd := HitButton(ntwrk)
		low = lowAdd + low
		high = highAdd + high
	}
	fmt.Printf("Part One: %v\n", low*high)
}
