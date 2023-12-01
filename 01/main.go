package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getConfigurationValue(input string) int {
	tokens := map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"0":     "0",
		"1":     "1",
		"2":     "2",
		"3":     "3",
		"4":     "4",
		"5":     "5",
		"6":     "6",
		"7":     "7",
		"8":     "8",
		"9":     "9",
	}
	lines := strings.Split(input, "\n")
	total := 0
	for _, line := range lines {
		var firstNum string
		var lastNum string
	first:
		for {
			for k, v := range tokens {
				if strings.HasPrefix(line, k) {
					firstNum = v
					break first
				}

			}
			line = line[1:]
		}
	last:
		for {
			for k, v := range tokens {
				if strings.HasSuffix(line, k) {
					lastNum = v
					break last
				}
			}
			line = line[:len(line)-1]
		}

		num, _ := strconv.Atoi(fmt.Sprintf("%s%s", firstNum, lastNum))
		total = total + num
	}
	return total
}

func main() {
	buf, _ := os.ReadFile("data.txt")
	fmt.Printf("Result: %v", getConfigurationValue(string(buf)))
}
