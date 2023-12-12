package main

import (
	"os"
	"regexp"
	"strconv"
	"strings"

	slicestuff "github.com/jack-barr3tt/gostuff/slices"
	strstuff "github.com/jack-barr3tt/gostuff/strings"
)

// If it matches the pattern and if we've got anything ? to deal with
func matchesPattern(pattern string, counts []int) (bool, bool) {
	unknownregex := regexp.MustCompile(`\?`)
	springregex := regexp.MustCompile(`#+`)

	qs := len(unknownregex.FindAllString(pattern, -1))
	springs := slicestuff.Reduce(func(spring string, acc int) int { return acc + len(spring) }, springregex.FindAllString(pattern, -1), 0)

	if springs > 0 && len(counts) == 0 {
		return false, true
	}
	pattern = unknownregex.ReplaceAllString(pattern, ".")

	newSprings := springregex.FindAllString(pattern, -1)
	if len(newSprings) != len(counts) {
		return false, qs == 0
	}
	for i, c := range counts {
		if c != len(newSprings[i]) {
			return false, qs == 0
		}
	}

	return true, true
}

// How much we can reduce the size of the string by
func reduce(springs string, count int) int {
	if len(springs) < count {
		return 0
	}
	for i := 0; i < count; i++ {
		if springs[i] == '.' {
			return 0
		}
	}
	if len(springs) == count {
		return count
	}

	if springs[count] == '.' || springs[count] == '?' {
		return count + 1
	}

	return 0
}

func countArrangements(s string, cond []int, visited map[string]int) int {
	key := s + strings.Join(slicestuff.Map(func(v int) string { return strconv.Itoa(v) }, cond), ",")
	if count, ok := visited[key]; ok {
		return count
	}

	visited[key] = 0
	if valid, end := matchesPattern(s, cond); end {
		if valid {
			visited[key] = 1
			return 1
		}
		return 0
	}

	if s[0] == '.' {
		n := countArrangements(s[1:], cond, visited)
		visited[key] = n
		return n
	}

	n := 0
	if s[0] == '?' {
		n += countArrangements(s[1:], cond, visited)
	}

	count := reduce(s, cond[0])
	if count == 0 {
		visited[key] = n
		return n
	}
	n += countArrangements(s[count:], cond[1:], visited)
	visited[key] = n
	return n
}

func main() {
	data, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(data), "\n")

	springsregex := regexp.MustCompile(`[#.\?]+`)

	p1 := 0
	p2 := 0

	for _, line := range lines {
		springs := springsregex.FindString(line)
		counts := strstuff.GetNums(line)

		visited1 := map[string]int{}

		newSprings := springs
		newCounts := counts

		for i := 0; i < 4; i++ {
			newSprings = newSprings + "?" + springs
			newCounts = append(newCounts, counts...)
		}

		visited2 := map[string]int{}

		p2 += countArrangements(newSprings, newCounts, visited2)

		p1 += countArrangements(springs, counts, visited1)
	}

	println("part 1:", p1)
	println("part 2:", p2)
}
