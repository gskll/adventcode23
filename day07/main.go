package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type handType int

const (
	highCard     handType = 0
	onePair      handType = 1
	twoPair      handType = 2
	threeOfAKind handType = 3
	fullHouse    handType = 4
	fourOfAKind  handType = 5
	fiveOfAKind  handType = 6
)

type hand struct {
	val      []byte
	bid      int
	handType handType
}

func parseHand(line string) (hand, error) {
	var h hand
	parts := strings.Fields(line)

	val := parseHandValues(parts[0])
	bid, err := strconv.Atoi(parts[1])
	if err != nil {
		return h, err
	}
	ht := parseHandType(val)

	h.val = val
	h.bid = bid
	h.handType = ht

	return h, nil
}

func parseHandType(handVal []byte) (ht handType) {
	frequency := make(map[byte]int, len(handVal))
	for i := range handVal {
		frequency[handVal[i]]++
	}
	maxFreq := 0
	for _, v := range frequency {
		if v > maxFreq {
			maxFreq = v
		}
	}

	switch len(frequency) {
	case 5:
		ht = highCard
	case 4:
		ht = onePair
	case 3:
		if maxFreq == 2 {
			ht = twoPair
		} else {
			ht = threeOfAKind
		}
	case 2:
		if maxFreq == 3 {
			ht = fullHouse
		} else {
			ht = fourOfAKind
		}
	case 1:
		ht = fiveOfAKind
	}
	return
}

func parseHandValues(hand string) []byte {
	var parsed []byte
	for i := range hand {
		switch hand[i] {
		case 'T':
			parsed = append(parsed, '9'+1-'1')
		case 'J':
			parsed = append(parsed, '9'+2-'1')
		case 'Q':
			parsed = append(parsed, '9'+3-'1')
		case 'K':
			parsed = append(parsed, '9'+4-'1')
		case 'A':
			parsed = append(parsed, '9'+5-'1')
		default:
			parsed = append(parsed, hand[i]-'1')
		}
	}
	return parsed
}

func solve(input string) {
	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var hands []hand
	for scanner.Scan() {
		hand, err := parseHand(scanner.Text())
		if err != nil {
			panic(err)
		}
		hands = append(hands, hand)
	}

	fmt.Println(hands)
}

func main() {
	inputFlag := flag.String("input", "day07/input.txt", "path to input file")
	flag.Parse()

	start := time.Now()
	solve(*inputFlag)
	duration := time.Since(start)
	fmt.Printf("Execution time: %v\n", duration)
}
