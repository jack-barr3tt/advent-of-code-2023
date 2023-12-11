package main

import (
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func findExpandRows(grid []string) []int {
	expandRows := []int{}

	for i, line := range grid {
		sub := regexp.MustCompile(`\.+`).FindString(line)

		if len(sub) == len(line) {
			expandRows = append(expandRows, i)
		}
	}

	return expandRows
}

func findExpansions(grid []string) [2][]int {
	return [2][]int{findExpandRows(grid), findExpandRows(flipGrid(grid))}
}

func flipGrid(grid []string) []string {
	newGrid := []string{}

	for i := range grid[0] {
		temp := ""
		for j := range grid {
			temp += string(grid[j][i])
		}
		newGrid = append(newGrid, temp)
	}

	return newGrid
}

func findGalaxies(grid []string) [][2]int {
	galaxyregex := regexp.MustCompile(`#`)
	coords := [][2]int{}

	for i, line := range grid {
		indicies := galaxyregex.FindAllStringIndex(line, -1)

		for _, index := range indicies {
			coords = append(coords, [2]int{index[0], i})
		}
	}

	return coords
}

func galaxyDistance(a, b [2]int) int {
	return int(math.Abs(float64(a[0]-b[0])) + math.Abs(float64(a[1]-b[1])))
}

func totalDistances(galaxies [][2]int) int {
	total := 0

	dists := make(map[string]int)

	for i, a := range galaxies {
		for j, b := range galaxies {
			if i == j {
				continue
			}
			temp := ""
			if i < j {
				temp = strconv.Itoa(i) + "," + strconv.Itoa(j)
			} else {
				temp = strconv.Itoa(j) + "," + strconv.Itoa(i)
			}

			if dists[temp] == 0 {
				total += galaxyDistance(a, b)
				dists[temp] = galaxyDistance(a, b)
				continue
			}

		}
	}

	return total
}

func expandCoords(grid []string, amount int) [][2]int {
	coords := findGalaxies(grid)
	expansions := findExpansions(grid)

	expandedCoords := [][2]int{}

	for _, coord := range coords {
		rows := 0
		cols := 0
		for _, expansion := range expansions[0] {
			if coord[1] > expansion {
				rows++
			}
		}
		for _, expansion := range expansions[1] {
			if coord[0] > expansion {
				cols++
			}
		}

		expandedCoords = append(expandedCoords, [2]int{coord[0] + cols*(amount-1), coord[1] + rows*(amount-1)})
	}

	return expandedCoords
}

func main() {
	data, _ := os.ReadFile("input.txt")

	grid := strings.Split(string(data), "\n")

	println("part 1: ", totalDistances(expandCoords(grid, 2)))
	println("part 2: ", totalDistances(expandCoords(grid, 1000000)))
}
