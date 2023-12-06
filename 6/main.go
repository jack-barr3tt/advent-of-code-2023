package main

import (
	"os"
	"regexp"
	"strconv"
	"strings"

	sliceutils "github.com/jack-barr3tt/gostuff/slices"
	strutils "github.com/jack-barr3tt/gostuff/strings"
	"github.com/jack-barr3tt/gostuff/types"
)

func waysToBeat(time int, distance int) int {
	count := 0
	for i := 0; i < time; i++ {
		if (time-i)*i > distance {
			count++
		}
	}

	return count
}

func main() {
	data, _ := os.ReadFile("input.txt")

	timeregex := regexp.MustCompile(`Time: +(\d+( +)?)+`)
	distregex := regexp.MustCompile(`Distance: +(\d+( +)?)+`)
	numregex := regexp.MustCompile(`\d+`)

	timestr := timeregex.FindString(string(data))
	diststr := distregex.FindString(string(data))

	times := strutils.GetNums(timestr)
	dists := strutils.GetNums(diststr)

	pairs := sliceutils.Zip(times, dists)

	println("part 1:", sliceutils.Reduce(func(curr types.Pair[int, int], acc int) int {
		return acc * waysToBeat(curr.First, curr.Second)
	}, pairs, 1))

	combtime := strings.Join(numregex.FindAllString(timestr, -1), "")
	combdist := strings.Join(numregex.FindAllString(diststr, -1), "")

	time, _ := strconv.Atoi(combtime)
	dist, _ := strconv.Atoi(combdist)

	println("part 2:", waysToBeat(time, dist))
}
