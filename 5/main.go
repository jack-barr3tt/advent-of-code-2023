package main

import (
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var numregex = regexp.MustCompile(`\d+`)

type PlantMap struct {
	destrange int
	srcrange  int
	width     int
}

type Range struct {
	start int
	width int
}

func getSeeds(data []byte) []int {
	seedsregex := regexp.MustCompile(`seeds: (\d+ )*\d+`)

	seedln := seedsregex.FindString(string(data))
	seedstr := numregex.FindAllString(seedln, -1)

	result := []int{}

	for _, seed := range seedstr {
		conv, _ := strconv.Atoi(seed)
		result = append(result, conv)
	}

	return result
}

func getPlantMap(data []byte, identifier string) []PlantMap {
	mapregex := regexp.MustCompile(identifier + ` map:(\n(\d+ )*\d+)+`)

	mapdata := mapregex.FindString(string(data))
	lines := strings.Split(mapdata, "\n")

	result := []PlantMap{}

	for i, line := range lines {
		if i == 0 {
			continue
		}

		numstrs := numregex.FindAllString(line, -1)
		start, _ := strconv.Atoi(numstrs[0])
		end, _ := strconv.Atoi(numstrs[1])
		target, _ := strconv.Atoi(numstrs[2])

		result = append(result, PlantMap{start, end, target})
	}

	return result
}

func useMap(val int, pmap []PlantMap) int {
	for _, m := range pmap {
		if val >= m.srcrange && val < m.srcrange+m.width {
			return m.destrange + (val - m.srcrange)
		}
	}
	return val
}

func useMapRange(r Range, pmap []PlantMap) []Range {
	result := []Range{}
	i := r
	for i.start != -1 {
		found := false
		for _, m := range pmap {
			if i.start >= m.srcrange && i.start < m.srcrange+m.width {
				if i.start+i.width < m.srcrange+m.width {
					// Case where range is completely enclapsed in map
					result = append(result, Range{useMap(i.start, pmap), i.width})
					i = Range{-1, -1}
					found = true
					break
				} else {
					// Case where a range needs splitting to be fully mapped
					result = append(result, Range{useMap(i.start, pmap), m.srcrange + m.width - i.start})
					i = Range{m.srcrange + m.width, i.width - (m.srcrange + m.width - i.start)}
					found = true
					break
				}
			}
		}
		// Case where a range is not mapped at all
		if !found {
			result = append(result, i)
			i = Range{-1, -1}
		}
	}

	return result
}

func main() {
	data, _ := os.ReadFile("input.txt")

	seeds := getSeeds(data)
	sts := getPlantMap(data, "seed-to-soil")
	stf := getPlantMap(data, "soil-to-fertilizer")
	ftw := getPlantMap(data, "fertilizer-to-water")
	wtl := getPlantMap(data, "water-to-light")
	ltt := getPlantMap(data, "light-to-temperature")
	tth := getPlantMap(data, "temperature-to-humidity")
	htl := getPlantMap(data, "humidity-to-location")

	pmaps := [][]PlantMap{sts, stf, ftw, wtl, ltt, tth, htl}

	min := int(^uint(0) >> 1)

	for _, seed := range seeds {
		for _, pmap := range pmaps {
			seed = useMap(seed, pmap)
		}
		min = int(math.Min(float64(min), float64(seed)))
	}

	println("part 1:", min)

	minr := int(^uint(0) >> 1)

	for i := 0; i < len(seeds); i += 2 {
		// For each seed range, we start with one range (obviously)
		ranges := []Range{Range{seeds[i], seeds[i+1]}}

		for _, pmap := range pmaps {
			// Expand out the ranges for each map in turn
			newranges := []Range{}
			for _, r := range ranges {
				newranges = append(newranges, useMapRange(r, pmap)...)
			}
			// And overwrite
			ranges = newranges
		}

		// The smallest start of a range is the minimum
		for _, r := range ranges {
			minr = int(math.Min(float64(minr), float64(r.start)))
		}
	}

	println("part 2:", minr)
}
