package main

import (
	"fmt"
	"os"
)

func myTestFunction(input string) int {
	return 1
}

func main() {
	buf, _ := os.ReadFile("data.txt")
	stringput := string(buf)
	fmt.Printf("Result: %v", myTestFunction(stringput))
}
