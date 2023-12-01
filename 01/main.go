package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getConfigurationValue(input string) int {
	lines := strings.Split(input, "\n")
	total := 0
	for _, line := range lines {
		var firstNum *string = nil
		var lastNum string
		for _, rn := range line {
			char := string(rn)
			if _, err := strconv.Atoi(char); err == nil {
				if firstNum == nil {
					firstNum = &char
				}
				lastNum = char
			}
		}
		num, _ := strconv.Atoi(fmt.Sprintf("%s%s", *firstNum, lastNum))
		total = total + num
	}
	return total
}

func main() {
	buf, _ := os.ReadFile("data.txt")
	fmt.Printf("Result: %v", getConfigurationValue(string(buf)))
}
