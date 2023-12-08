package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"time"
)

type (
	graph map[string]node
	node  struct {
		left  string
		right string
	}
)

func isStartNode(label string) bool {
	return label[2] == 'A'
}

func isEndNode(label string) bool {
	return label[2] == 'Z'
}

func walk(directions string, g graph) (stepCount int) {
	curr, end := "AAA", "ZZZ"
	for curr != end {
		node := g[curr]
		directionIndex := stepCount % len(directions)
		direction := string(directions[directionIndex])
		if direction == "L" {
			curr = node.left
		} else {
			curr = node.right
		}
		stepCount++
	}
	return
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func walk2(directions string, g graph, nodes []string) (stepCount int) {
	stepsPerPath := make([]int, len(nodes))
	for i := range nodes {
		s := 0
		label := nodes[i]
		for stepsPerPath[i] == 0 {
			if isEndNode(label) {
				stepsPerPath[i] = s
				break
			}

			if directions[s%len(directions)] == 'L' {
				label = g[label].left
			} else {
				label = g[label].right
			}
			s++
		}
	}

	stepCount = 1
	for _, s := range stepsPerPath {
		stepCount = lcm(stepCount, s)
	}
	return
}

func insertNode(line string, g graph) string {
	re := regexp.MustCompile(`\w+`)
	matches := re.FindAllString(line, 3)
	g[matches[0]] = node{left: matches[1], right: matches[2]}
	return matches[0]
}

func parseLines(scanner *bufio.Scanner) (directions string, g graph, startNodes []string) {
	g = make(graph)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if directions == "" {
			directions = scanner.Text()
			continue
		}
		label := insertNode(line, g)
		if isStartNode(label) {
			startNodes = append(startNodes, label)
		}
	}
	return
}

func solve(input string) {
	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	directions, graph, startNodes := parseLines(scanner)

	// pathLength2 := walk(directions, graph)
	// fmt.Println("Part 1 count:", pathLength2, pathLength2 == 21409)
	pathLength2 := walk2(directions, graph, startNodes)
	fmt.Println("Part 2 count:", pathLength2, pathLength2 == 21409)
}

func main() {
	inputFlag := flag.String("input", "input.txt", "path to input file")
	flag.Parse()

	start := time.Now()
	solve(*inputFlag)
	duration := time.Since(start)
	fmt.Printf("Execution time: %v\n", duration)
}
