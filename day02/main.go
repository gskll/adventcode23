package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cubes = map[string]int{"red": 12, "green": 13, "blue": 14}

func parseGameId(gameId string) (int, error) {
	parts := strings.Fields(gameId)
	id, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, err
	}
	return id, nil
}

func calculateMaxCubes(setsStr string) (map[string]int, error) {
	maxCubes := map[string]int{"red": 0, "green": 0, "blue": 0}

	for _, set := range strings.Split(setsStr, "; ") {
		for _, colorStr := range strings.Split(set, ", ") {
			parts := strings.Fields(colorStr)
			numCubes, err := strconv.Atoi(parts[0])
			if err != nil {
				return nil, err
			}
			color := parts[1]

			if numCubes > maxCubes[color] {
				maxCubes[color] = numCubes
			}
		}
	}
	return maxCubes, nil
}

func main() {
	// file, err := os.Open("test_input_p1.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sum := 0
	powerSum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		gameData := strings.SplitN(line, ": ", 2)
		id, err := parseGameId(gameData[0])
		if err != nil {
			panic(err)
		}
		maxCubes, err := calculateMaxCubes(gameData[1])
		if err != nil {
			panic(err)
		}

		gamePossible := true
		power := 1
		for color, value := range maxCubes {
			if value > cubes[color] {
				gamePossible = false
			}
			power *= value
		}
		if gamePossible {
			sum += id
		}
		powerSum += power
	}

	fmt.Println("The sum for part 1: ", sum)                // 2101
	fmt.Println("The sum of powers for part 2: ", powerSum) // 58269
}
