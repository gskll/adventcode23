package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var cubes = map[string]int{"red": 12, "green": 13, "blue": 14}

func parseLine(line string) (id int, maxCubes map[string]int) {
	gameData := strings.Split(line, ": ")
	idStr := ""
	for _, r := range gameData[0] {
		if unicode.IsDigit(r) {
			idStr += string(r)
		}
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}
	setsStr := strings.Split(gameData[1], "; ")
	maxCubes = map[string]int{"red": 0, "green": 0, "blue": 0}

	for _, setStr := range setsStr {
		sets := strings.Split(setStr, ", ")
		for _, set := range sets {
			tmp := strings.Split(set, " ")
			color := tmp[1]
			numCubes, err := strconv.Atoi(tmp[0])
			if err != nil {
				panic(err)
			}
			if maxCubes[color] < numCubes {
				maxCubes[color] = numCubes
			}
		}
	}
	return
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
		id, maxCubes := parseLine(line)
		gamePossible := true

		power := 1
		for color, value := range maxCubes {
			fmt.Println(color, value)
			if value > cubes[color] {
				gamePossible = false
			}
			power *= value
		}
		fmt.Println(power)
		if gamePossible {
			sum += id
		}
		powerSum += power
	}

	fmt.Println("The sum for part 1: ", sum) // 2101
	fmt.Println("The sum of powers for part 2: ", powerSum)
}
