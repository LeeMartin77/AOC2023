package main

import (
	"fmt"
	"os"
)

type ModuleType int64

const (
	Broadcaster ModuleType = 0
	FlipFlop    ModuleType = 1
	Conjunction ModuleType = 2
)

type CommsModule struct {
	Name    string
	SendsTo []string
	Type    ModuleType
}

func (cms CommsModule) SendPulse(pulse bool) {

}

func (cms CommsModule) ProcessPulses() {

}

func ParseCommsModule()

func myTestFunction(input string) int {
	return 1
}

func main() {
	buf, _ := os.ReadFile("data.txt")
	stringput := string(buf)
	fmt.Printf("Result: %v", myTestFunction(stringput))
}
