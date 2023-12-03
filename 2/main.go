package main

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Draw struct {
	red   int
	green int
	blue  int
}

var noregex = regexp.MustCompile("[0-9]{1,3}")

func getGameNo(input string) int {
	gnoregex, _ := regexp.Compile("Game [0-9]{1,3}")

	gnostr := gnoregex.FindString(input)
	nostr := noregex.FindString(gnostr)

	gameno, _ := strconv.Atoi(nostr)

	return gameno
}

func getDraws(input string) []Draw {
	redregex := regexp.MustCompile("[0-9]{1,3} red")
	greenregex := regexp.MustCompile("[0-9]{1,3} green")
	blueregex := regexp.MustCompile("[0-9]{1,3} blue")

	drawsstr := strings.Split(input, ": ")[1]
	drawssstr := strings.Split(drawsstr, "; ")

	draws := make([]Draw, len(drawssstr))

	for i := range drawssstr {
		drawstr := drawssstr[i]

		redstr := redregex.FindString(drawstr)
		greenstr := greenregex.FindString(drawstr)
		bluestr := blueregex.FindString(drawstr)

		red, green, blue := 0, 0, 0

		if redstr != "" {
			red, _ = strconv.Atoi(noregex.FindString(redstr))
		}
		if greenstr != "" {
			green, _ = strconv.Atoi(noregex.FindString(greenstr))
		}
		if bluestr != "" {
			blue, _ = strconv.Atoi(noregex.FindString(bluestr))
		}

		draws[i] = Draw{red, green, blue}
	}

	return draws
}

func printDraw(d Draw) {
	println("Red:", d.red, "Green:", d.green, "Blue:", d.blue)
}

func printDraws(draws []Draw) {
	for i := range draws {
		printDraw(draws[i])
	}
}

func maxForColor(draws []Draw) Draw {
	maxred, maxgreen, maxblue := 0, 0, 0

	for i := range draws {
		draw := draws[i]

		if draw.red > maxred {
			maxred = draw.red
		}
		if draw.green > maxgreen {
			maxgreen = draw.green
		}
		if draw.blue > maxblue {
			maxblue = draw.blue
		}
	}

	return Draw{maxred, maxgreen, maxblue}
}

func main() {

	data, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(data), "\n")

	idsum := 0
	powersum := 0

	for i := range lines {
		str := lines[i]

		draws := getDraws(str)
		max := maxForColor(draws)

		if max.red <= 12 && max.green <= 13 && max.blue <= 14 {
			idsum += getGameNo(str)
		}

		powersum += max.red * max.green * max.blue
	}

	println("Part 1: ", idsum)
	println("Part 2: ", powersum)
}
