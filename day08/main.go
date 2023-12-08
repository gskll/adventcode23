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

func insertNode(line string, g graph) {
	re := regexp.MustCompile(`\w+`)
	matches := re.FindAllString(line, 3)
	g[matches[0]] = node{left: matches[1], right: matches[2]}
}

func parseLines(scanner *bufio.Scanner) (directions string, g graph) {
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
		insertNode(line, g)
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
	directions, nodes := parseLines(scanner)

	pathLength := walk(directions, nodes)
	fmt.Println("Part 1 count:", pathLength)
}

func main() {
	inputFlag := flag.String("input", "input.txt", "path to input file")
	flag.Parse()

	start := time.Now()
	solve(*inputFlag)
	duration := time.Since(start)
	fmt.Printf("Execution time: %v\n", duration)
}
