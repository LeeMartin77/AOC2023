package main

import (
	"fmt"
	"os"
	"strings"
)

func CalculateHash(input string) int {
	cuml := 0
	for _, ch := range input {
		cuml = cuml + int(ch)
		cuml = cuml * 17
		cuml = cuml % 256
	}
	return cuml
}

func main() {
	buf, _ := os.ReadFile("data.txt")
	input := string(buf)
	cuml := 0
	for _, com := range strings.Split(input, ",") {
		cuml = cuml + CalculateHash(com)
	}
	fmt.Printf("Result: %v\n", cuml)
}
