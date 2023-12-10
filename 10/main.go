package main

import (
	"math"
	"os"
	"regexp"
	"strings"

	slicestuff "github.com/jack-barr3tt/gostuff/slices"
)

var moveInMap = makeMoveInMap()
var moveOutMap = makeMoveOutMap()
var pieceMap = makePieceMap()

func makeMoveInMap() map[string][]bool {
	m := make(map[string][]bool)

	// Up, Down, Left, Right
	m["|"] = []bool{true, true, false, false}
	m["-"] = []bool{false, false, true, true}
	m["L"] = []bool{false, true, true, false}
	m["J"] = []bool{false, true, false, true}
	m["7"] = []bool{true, false, false, true}
	m["F"] = []bool{true, false, true, false}
	m["S"] = []bool{true, true, true, true}
	m["."] = []bool{false, false, false, false}

	return m
}

func makeMoveOutMap() map[string][]bool {
	m := make(map[string][]bool)

	// Up, Down, Left, Right
	m["|"] = []bool{true, true, false, false}
	m["-"] = []bool{false, false, true, true}
	m["L"] = []bool{true, false, false, true}
	m["J"] = []bool{true, false, true, false}
	m["7"] = []bool{false, true, true, false}
	m["F"] = []bool{false, true, false, true}
	m["S"] = []bool{true, true, true, true}
	m["."] = []bool{false, false, false, false}

	return m
}

func makePieceMap() map[string][]string {
	m := make(map[string][]string)

	m["|"] = []string{" # ", " # ", " # "}
	m["-"] = []string{"   ", "###", "   "}
	m["L"] = []string{" # ", " ##", "   "}
	m["J"] = []string{" # ", "## ", "   "}
	m["7"] = []string{"   ", "## ", " # "}
	m["F"] = []string{"   ", " ##", " # "}
	m["S"] = []string{"###", "###", "###"}
	m["."] = []string{"   ", "   ", "   "}

	return m
}

func canMove(grid []string, curr [2]int, dir int) bool {
	pos := move(curr, dir)

	if pos[0] < 0 || pos[0] >= len(grid[0]) || pos[1] < 0 || pos[1] >= len(grid) {
		return false
	}

	return moveInMap[string(grid[pos[1]][pos[0]])][dir]
}

func move(curr [2]int, dir int) [2]int {
	pos := curr
	if dir == 0 {
		pos[1]--
	} else if dir == 1 {
		pos[1]++
	} else if dir == 2 {
		pos[0]--
	} else if dir == 3 {
		pos[0]++
	}

	return pos
}

func replace(gridVals [][]int, pos [2]int, val int) [][]int {
	if gridVals[pos[1]][pos[0]] == 0 {
		gridVals[pos[1]][pos[0]] = val
	} else {
		gridVals[pos[1]][pos[0]] = int(math.Min(float64(val), float64(gridVals[pos[1]][pos[0]])))
	}
	return gridVals
}

func floodFill(grid []string, pos [2]int, val byte) []string {
	moveq := [][2]int{pos}

	for len(moveq) > 0 {
		curr := moveq[0]
		moveq = moveq[1:]

		grid[curr[1]] = grid[curr[1]][:curr[0]] + string(val) + grid[curr[1]][curr[0]+1:]

		for i := 0; i < 4; i++ {
			newPos := move(curr, i)

			if newPos[0] < 0 || newPos[0] >= len(grid[0]) || newPos[1] < 0 || newPos[1] >= len(grid) {
				continue
			}

			if grid[newPos[1]][newPos[0]] == ' ' {
				if !slicestuff.Some(func(val [2]int) bool {
					return val[0] == newPos[0] && val[1] == newPos[1]
				}, moveq) {
					moveq = append(moveq, newPos)
				}
			}
		}
	}

	return grid
}

func expandGrid(grid []string) []string {
	newGrid := []string{}

	for _, line := range grid {
		temp := [][]string{}

		for _, char := range line {
			temp = append(temp, pieceMap[string(char)])
		}

		for i := 0; i < 3; i++ {
			newLine := ""
			for _, piece := range temp {
				newLine += piece[i]
			}
			newGrid = append(newGrid, newLine)
		}
	}

	return newGrid
}

func main() {
	data, _ := os.ReadFile("input.txt")

	grid := strings.Split(string(data), "\n")

	gridVals := make([][]int, len(grid))
	for i := range grid {
		gridVals[i] = make([]int, len(grid[i]))
	}

	startpos := slicestuff.Reduce(func(line string, acc [2]int) [2]int {
		loc := regexp.MustCompile("S").FindStringIndex(string(line))
		if acc[0] == -1 {
			acc[1]++
		}
		if loc != nil {
			return [2]int{int(math.Max(float64(acc[0]), float64(loc[0]))), acc[1]}
		}
		return acc
	}, grid, [2]int{-1, -1})

	gridVals[startpos[1]][startpos[0]] = 1

	moveq := [][2]int{startpos}

	last := gridVals[startpos[1]][startpos[0]]

	for len(moveq) > 0 {
		curr := moveq[0]
		moveq = moveq[1:]

		last = gridVals[curr[1]][curr[0]]

		for i := 0; i < 4; i++ {
			if moveOutMap[string(grid[curr[1]][curr[0]])][i] && canMove(grid, curr, i) {
				newPos := move(curr, i)

				if gridVals[newPos[1]][newPos[0]] == 0 || gridVals[curr[1]][curr[0]]+1 < gridVals[newPos[1]][newPos[0]] {
					moveq = append(moveq, newPos)
					gridVals = replace(gridVals, newPos, gridVals[curr[1]][curr[0]]+1)
				}
			}
		}
	}
	println("part 1:", last-1)

	// Take all the unconnected junk out of the grid
	filterGrid := []string{}
	for i := 0; i < len(grid); i++ {
		temp := ""
		for j := 0; j < len(grid[i]); j++ {
			if gridVals[i][j] != 0 {
				temp += string(grid[i][j])
			} else {
				temp += "."
			}
		}
		filterGrid = append(filterGrid, temp)
	}

	// Expand each piece into a 3x3 grid
	exGrid := expandGrid(filterGrid)

	// Flood fill to show the area enclosed by the pipe
	exGrid = floodFill(exGrid, [2]int{0, 0}, '#')

	// Count the number of 3x3 empty grids, which translate to a single empty pixel
	count := 0
	for i := 0; i < len(exGrid); i += 3 {
		for j := 0; j < len(exGrid[i]); j += 3 {
			found := false
			for x := 0; x < 3; x++ {
				for y := 0; y < 3; y++ {
					if exGrid[i+x][j+y] != ' ' {
						found = true
					}
				}
			}

			if !found {
				count++
			}
		}
	}

	println("part 2:", count)
}
