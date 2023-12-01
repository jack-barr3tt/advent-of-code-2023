package main

import (
	"os"
	"strconv"
	"strings"
)

func genWordMap() map[string]int {
	m := make(map[string]int)

	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
	m["four"] = 4
	m["five"] = 5
	m["six"] = 6
	m["seven"] = 7
	m["eight"] = 8
	m["nine"] = 9

	return m
}

func getInt(str string, i int) int {
	if rune(str[i]) >= 48 && rune(str[i]) <= 57 {
		num, err := strconv.Atoi(string(str[i]))
		if err != nil {
			panic(err)
		}
		return num
	}

	return -1
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	wordMap := genWordMap()

	lines := strings.Split(string(data), "\n")
	lens := []int{3, 4, 5}

	sum := 0

	for i := range lines {
		str := lines[i]
		first := -1
		last := -1

		for j := 0; j < len(str); j++ {

			if first == -1 {
				first = getInt(str, j)

				if first == -1 {
					for _, l := range lens {
						if j+l >= len(str) {
							continue
						}
						if wordMap[str[j:j+l]] != 0 {
							first = wordMap[str[j:j+l]]
							break
						}
					}
				}
			}

			if last == -1 {
				last = getInt(str, len(str)-j-1)

				if last == -1 {
					for _, l := range lens {
						if len(str)-j-1-l < 0 {
							continue
						}
						if wordMap[str[len(str)-j-l:len(str)-j]] != 0 {
							last = wordMap[str[len(str)-j-l:len(str)-j]]
							break
						}
					}
				}
			}

			if first != -1 && last != -1 {
				break
			}
		}

		num, err := strconv.Atoi(strconv.Itoa(first) + strconv.Itoa(last))
		if err != nil {
			panic(err)
		}
		sum += num
	}

	println(sum)
}
