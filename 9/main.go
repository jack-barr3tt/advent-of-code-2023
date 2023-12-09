package main

import (
	"os"
	"strings"

	strstuff "github.com/jack-barr3tt/gostuff/strings"
)

func elementChange(nums []int) ([]int, bool) {
	foundNonZero := false

	newNums := make([]int, len(nums)-1)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != 0 {
			foundNonZero = true
		}

		newNums[i] = nums[i+1] - nums[i]
	}

	return newNums, !foundNonZero
}

func main() {
	data, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(data), "\n")

	futureSum := 0
	historySum := 0

	for _, line := range lines {
		nums := strstuff.GetNums(line)

		allDiffs := [][]int{nums}

		for diffs, done := elementChange(nums); !done; diffs, done = elementChange(diffs) {
			allDiffs = append(allDiffs, diffs)
		}
		for i := len(allDiffs) - 2; i >= 0; i-- {
			allDiffs[i] = append(allDiffs[i], allDiffs[i][len(allDiffs[i])-1]+allDiffs[i+1][len(allDiffs[i+1])-1])
			allDiffs[i] = append([]int{allDiffs[i][0] - allDiffs[i+1][0]}, allDiffs[i]...)
		}

		futureSum += allDiffs[0][len(allDiffs[0])-1]
		historySum += allDiffs[0][0]
	}

	println("part 1:", futureSum)
	println("part 2:", historySum)
}
