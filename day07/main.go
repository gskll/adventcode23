package main

import (
	"bufio"
	"cmp"
	"flag"
	"fmt"
	"os"
	"slices"
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
	val      string // debugging only
	cmp      string
	bid      int
	handType handType
}

func parseHand(line string) (hand, error) {
	var h hand
	parts := strings.Fields(line)
	handString := parts[0]
	bidString := parts[1]

	cmp := parseHandValues(handString)
	ht := parseHandType(handString)
	bid, err := strconv.Atoi(bidString)
	if err != nil {
		return h, err
	}

	h.val = handString
	h.cmp = cmp
	h.bid = bid
	h.handType = ht

	return h, nil
}

func parseHandType(hand string) (ht handType) {
	frequency := make(map[byte]int, len(hand))
	for i := range hand {
		frequency[hand[i]]++
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

func parseHandValues(hand string) string {
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
	return string(parsed)
}

func sortHands(hands []hand) []hand {
	slices.SortFunc(hands, func(a, b hand) int {
		if n := cmp.Compare(a.handType, b.handType); n != 0 {
			return n
		}
		return cmp.Compare(a.cmp, b.cmp)
	})
	return hands
}

func calculateWinnings(hands []hand) int {
	sortedHands := sortHands(hands)
	res := 0
	for i, hand := range sortedHands {
		res += (i + 1) * hand.bid
	}
	return res
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

	winnings := calculateWinnings(hands)
	fmt.Println("Part 1 solution:", winnings) // 247823654
}

func main() {
	inputFlag := flag.String("input", "day07/input.txt", "path to input file")
	flag.Parse()

	start := time.Now()
	solve(*inputFlag)
	duration := time.Since(start)
	fmt.Printf("Execution time: %v\n", duration)
}
