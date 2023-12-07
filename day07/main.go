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

const (
	joker = '2' - 1 - '1'
	ace   = '9' + 5 - '1'
	king  = '9' + 4 - '1'
	queen = '9' + 3 - '1'
	jack  = '9' + 2 - '1'
	ten   = '9' + 1 - '1'
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

func parseHand(line string) (hand, hand, error) {
	var h, h2 hand
	parts := strings.Fields(line)
	handString := parts[0]
	bidString := parts[1]

	cmp, cmp2 := parseHandValues(handString)
	ht := parseHandType(cmp, false)
	ht2 := parseHandType(cmp2, true)
	bid, err := strconv.Atoi(bidString)
	if err != nil {
		return h, h2, err
	}

	h.val = handString
	h.cmp = cmp
	h.bid = bid
	h.handType = ht

	h2.val = handString
	h2.cmp = cmp2
	h2.bid = bid
	h2.handType = ht2

	return h, h2, nil
}

func parseHandType(hand string, withJoker bool) (ht handType) {
	frequency := make(map[byte]int, len(hand))
	for i := range hand {
		frequency[hand[i]]++
	}

	if withJoker {
		var maxFreqCard byte = joker
		for cardKey := range frequency {
			if cardKey != joker && (maxFreqCard == joker || frequency[cardKey] > frequency[maxFreqCard]) {
				maxFreqCard = cardKey
			}
		}
		if jokerFreq, exists := frequency[joker]; exists && maxFreqCard != joker {
			frequency[maxFreqCard] += jokerFreq
			delete(frequency, joker)
		}
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

func parseHandValues(hand string) (string, string) {
	var parsed []byte
	var parsed2 []byte
	for i := range hand {
		switch hand[i] {
		case 'T':
			parsed = append(parsed, ten)
			parsed2 = append(parsed2, ten)
		case 'J':
			parsed = append(parsed, jack)
			parsed2 = append(parsed2, joker)
		case 'Q':
			parsed = append(parsed, queen)
			parsed2 = append(parsed2, queen)
		case 'K':
			parsed = append(parsed, king)
			parsed2 = append(parsed2, king)
		case 'A':
			parsed = append(parsed, ace)
			parsed2 = append(parsed2, ace)
		default:
			parsed = append(parsed, hand[i]-'1')
			parsed2 = append(parsed2, hand[i]-'1')
		}
	}
	return string(parsed), string(parsed2)
}

func sortHands(hands []hand) []hand {
	slices.SortFunc(hands, func(a, b hand) int {
		if n := cmp.Compare(a.handType, b.handType); n != 0 {
			return n
		}
		return cmp.Compare(string(a.cmp), string(b.cmp))
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
	var hands2 []hand
	for scanner.Scan() {
		hand, hand2, err := parseHand(scanner.Text())
		if err != nil {
			panic(err)
		}
		hands = append(hands, hand)
		hands2 = append(hands2, hand2)
	}

	winnings := calculateWinnings(hands)
	fmt.Println("Part 1 solution:", winnings, winnings == 247823654) // 247823654

	winnings2 := calculateWinnings(hands2)
	fmt.Println("Part 2 solution:", winnings2, winnings2 == 245461700) // 245461700
}

func main() {
	inputFlag := flag.String("input", "day07/input.txt", "path to input file")
	flag.Parse()

	start := time.Now()
	solve(*inputFlag)
	duration := time.Since(start)
	fmt.Printf("Execution time: %v\n", duration)
}
