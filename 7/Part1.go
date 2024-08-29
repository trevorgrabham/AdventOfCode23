package main

// type Type int
// type Rank int

// type Hand struct {
// 	cards []Rank
// 	original []Rank
// 	kind  Type
// 	bet		int
// }

// const (
// 	HighCard Type = iota
// 	Pair
// 	TwoPair
// 	ThreeOfAKind
// 	FullHouse
// 	FourOfAKind
// 	FiveOfAKind
// )

// var Ranks map[rune]Rank = map[rune]Rank{
// 	'2': 0,
// 	'3': 1,
// 	'4': 2,
// 	'5': 3,
// 	'6': 4,
// 	'7': 5,
// 	'8': 6,
// 	'9': 7,
// 	'T': 8,
// 	'J': 9,
// 	'Q': 10,
// 	'K': 11,
// 	'A': 12,
// }

// func NewHand(cards string, bet int) Hand {
// 	hand := Hand{}
// 	hand.bet = bet
// 	hand.kind = -1
// 	temp := make([]int, 0, 5)
// 	for _, card := range cards {
// 		temp = append(temp, int(Ranks[card]))
// 		hand.original = append(hand.original, Ranks[card])
// 	}
// 	sort.Ints(temp)
// 	for _, r := range temp {
// 		hand.cards = append(hand.cards, Rank(r))
// 	}
// 	hand.kind = hand.GetKind()
// 	return hand
// }

// func (h *Hand) GetKind() Type {
// 	if h.kind >= 0 {
// 		return h.kind
// 	}
// 	if h.checkFiveOfAKind() {
// 		return FiveOfAKind
// 	}
// 	if h.checkFourOfAKind() {
// 		return FourOfAKind
// 	}
// 	if h.checkFullHouse() {
// 		return FullHouse
// 	}
// 	if h.checkThreeOfAKind() {
// 		return ThreeOfAKind
// 	}
// 	if h.checkTwoPair() {
// 		return TwoPair
// 	}
// 	if h.checkPair() {
// 		return Pair
// 	}
// 	return HighCard
// }

// func (h *Hand) checkFiveOfAKind() bool {
// 	return h.cards[0] == h.cards[4]
// }

// func (h *Hand) checkFourOfAKind() bool {
// 	otherCardIsLower := h.cards[1] == h.cards[4]
// 	otherCardIsHigher := h.cards[0] == h.cards[3]
// 	return otherCardIsHigher || otherCardIsLower
// }

// func (h *Hand) checkFullHouse() bool {
// 	containsTwoPairs := h.cards[0] == h.cards[1] && h.cards[3] == h.cards[4]
// 	oneIsThreeOfAKind := h.cards[2] == h.cards[1] || h.cards[2] == h.cards[3]
// 	return containsTwoPairs && oneIsThreeOfAKind
// }

// func (h *Hand) checkThreeOfAKind() bool {
// 	firstThree := h.cards[0] == h.cards[1] && h.cards[1] == h.cards[2]
// 	middleThree := h.cards[1] == h.cards[2] && h.cards[2] == h.cards[3]
// 	lastThree := h.cards[2] == h.cards[3] && h.cards[3] == h.cards[4]
// 	return firstThree || middleThree || lastThree
// }

// /*
// 	Since the cards are sorted, going to track where the non-paired card is. The pairs must be continuous, so we only have three situations:
// 		11X22
// 		X1122
// 		1122X
// */
// func (h *Hand) checkTwoPair() bool {
// 	misfitAtStart := h.cards[1] == h.cards[2] && h.cards[3] == h.cards[4]
// 	mistfitInMiddle := h.cards[0] == h.cards[1] && h.cards[3] == h.cards[4]
// 	mistfitAtEnd := h.cards[0] == h.cards[1] && h.cards[2] == h.cards[3]
// 	return misfitAtStart || mistfitInMiddle || mistfitAtEnd
// }

// func (h *Hand) checkPair() bool {
// 	firstTwo := h.cards[0] == h.cards[1]
// 	secondTwo := h.cards[1] == h.cards[2]
// 	thirdTwo := h.cards[2] == h.cards[3]
// 	lastTwo := h.cards[3] == h.cards[4]
// 	return firstTwo || secondTwo || thirdTwo || lastTwo
// }

// func (h *Hand) Len() int {
// 	return len(h.cards)
// }

// func (h *Hand) Less(i, j int) bool {
// 	return h.cards[i] < h.cards[j]
// }

// func (h *Hand) Swap(i, j int) {
// 	h.cards[i], h.cards[j] = h.cards[j], h.cards[i]
// }

// func main() {
// 	file, err := os.Open("input.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	scanner := bufio.NewScanner(file)
// 	var hands []Hand
// 	for scanner.Scan() {
// 		line := strings.Split(scanner.Text(), " ")
// 		bet, err := strconv.Atoi(line[1])
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		hands = append(hands, NewHand(line[0], bet))
// 	}
// 	if scanner.Err() != nil {
// 		log.Fatal(scanner.Err())
// 	}
// 	sort.Slice(hands, func(i, j int) bool {
// 		a, b := hands[i], hands[j]
// 		if a.kind < b.kind {
// 			return true
// 		}
// 		if a.kind != b.kind {
// 			return false
// 		}
// 		for index := range hands {
// 			if a.original[index] < b.original[index] {
// 				return true
// 			}
// 			if a.original[index] > b.original[index] {
// 				return false
// 			}
// 		}
// 		return false
// 	})
// 	var sum int
// 	for i, hand := range hands {
// 		sum += hand.bet * (i+1)
// 	}
// 	fmt.Println(sum)
// }
