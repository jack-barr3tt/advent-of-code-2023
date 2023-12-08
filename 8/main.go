package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	mapstuff "github.com/jack-barr3tt/gostuff/maps"
	numstuff "github.com/jack-barr3tt/gostuff/nums"
	slicestuff "github.com/jack-barr3tt/gostuff/slices"
)

func genNodes(lines []string) map[string][]string {
	nodesMap := make(map[string][]string)

	noderegex := regexp.MustCompile(`[A-Z0-9]{3}`)
	rootregex := regexp.MustCompile(`[A-Z0-9]{3} =`)
	nodesregex := regexp.MustCompile(`\(([A-Z0-9]{3}, )*([A-Z0-9]{3})\)`)

	for _, line := range lines {
		rootstr := rootregex.FindString(line)
		root := noderegex.FindString(rootstr)
		nodesstr := nodesregex.FindString(line)
		nodes := noderegex.FindAllString(nodesstr, -1)

		nodesMap[root] = nodes
	}
	return nodesMap
}

func countSteps(start, end, feed string, nodes map[string][]string) int {
	node := start
	exp := regexp.MustCompile(end)
	for x := 0; true; x++ {
		i := x % len(feed)

		if feed[i] == 'R' {
			node = nodes[node][1]
		} else {
			node = nodes[node][0]
		}
		if exp.Match([]byte(node)) {
			return x + 1
		}
	}

	return -1
}

func main() {
	data, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(data), "\n")

	feed := lines[0]

	nodes := genNodes(lines[2:])

	println("part 1:", countSteps("AAA", "ZZZ", feed, nodes))

	startNodes := slicestuff.Filter(func(k string) bool { return k[2] == 'A' }, mapstuff.Keys(nodes))

	dists := slicestuff.Map(func(node string) int {
		return countSteps(node, "[A-Z0-9]{2}Z", feed, nodes)
	}, startNodes)

	fmt.Println("part 2:", numstuff.FindLCM(dists))
}
