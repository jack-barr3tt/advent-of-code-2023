package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	sum := 0

	for i := range lines {
		str := lines[i]
		first := -1
		last := -1

		for j := 0; j < len(str); j++ {
			if first == -1 && rune(str[j]) >= 48 && rune(str[j]) <= 57 {
				first = j
			}

			if last == -1 && rune(str[len(str)-j-1]) >= 48 && rune(str[len(str)-j-1]) <= 57 {
				last = len(str) - j - 1
			}

			if first != -1 && last != -1 {
				break
			}
		}

		num, err := strconv.Atoi(string(str[first]) + string(str[last]))
		if err != nil {
			panic(err)
		}

		sum += num
	}

	println(sum)
}
