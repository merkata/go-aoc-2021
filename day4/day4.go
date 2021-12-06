package day4

import (
	"strconv"
	"strings"
)

// BingoPlay is a set of bingo cards.
type BingoPlay struct {
	Cards   []BingoCard
	Winners int
}

// Check looks whether you have a bingo in one of your cards
// with the latest announced number and returns the sum of
// all unannounced numbers of the winning card.
func (bp *BingoPlay) Check(candidate int) (int, bool) {

	for _, card := range bp.Cards {
		card.TryMark(candidate)
		if card.Wins() {
			return card.Sum(), true
		}
	}
	return 0, false
}

// CheckLast looks whether you have a bingo in one of your cards
// with the latest announced number and whether it is the last winning card.
func (bp *BingoPlay) CheckLast(candidate int) (int, bool) {

	for _, card := range bp.Cards {
		if !card.Wins() {
			card.TryMark(candidate)
			if card.Wins() {
				bp.Winners++
				if bp.Winners == len(bp.Cards) {
					return card.Sum(), true
				}
			}
		}
	}
	return 0, false
}

// BingoCard implements a card of bingo
// keeping a reference of columns and rows
// and which nnumbers have been announced.
type BingoCard struct {
	Cols   [5][5]int
	Rows   [5][5]int
	Marked map[int]bool
	HasWon bool
}

// TryMark checks whether a number is in a given card
// and marks it in the card.
func (bc *BingoCard) TryMark(number int) {
	for _, col := range bc.Cols {
		for _, val := range col {
			if val == number {
				bc.Marked[number] = true
			}
		}
	}
}

// Wins determines whether you have a bingo
// on your card horizontally or vertically.
func (bc *BingoCard) Wins() bool {
	for _, col := range bc.Cols {
		if bc.Marked[col[0]] == true &&
			bc.Marked[col[1]] == true &&
			bc.Marked[col[2]] == true &&
			bc.Marked[col[3]] == true &&
			bc.Marked[col[4]] == true {
			return true
		}
	}
	for _, row := range bc.Rows {
		if bc.Marked[row[0]] == true &&
			bc.Marked[row[1]] == true &&
			bc.Marked[row[2]] == true &&
			bc.Marked[row[3]] == true &&
			bc.Marked[row[4]] == true {
			return true
		}
	}
	return false
}

// Sum returns the result of all unannounced numbers
// from a (winning) bingo card.
func (bc *BingoCard) Sum() int {
	var sum int
	for _, col := range bc.Cols {
		for _, val := range col {
			if !bc.Marked[val] {
				sum += val
			}
		}
	}
	return sum
}

// BingoWinner implements the game of Bingo
// and determines the first winner
func BingoWinner(input []string) int {
	numbers := strings.Split(input[0], ",")
	bp := BingoPlay{}
	card := []string{}
	for _, cardline := range input[2:] {
		if cardline == "" {
			bp.Cards = append(bp.Cards, NewBingoCardFromString(card))
			card = []string{}
			continue
		}
		card = append(card, cardline)
	}
	bp.Cards = append(bp.Cards, NewBingoCardFromString(card))
	for _, number := range numbers {
		candidate, err := strconv.Atoi(number)
		if err != nil {
			candidate = 0
		}
		if result, ok := bp.Check(candidate); ok {
			return result * candidate
		}
	}
	return 0
}

// BingoLoser implements the game of Bingo
// and determines the last winner
func BingoLoser(input []string) int {
	numbers := strings.Split(input[0], ",")
	bp := BingoPlay{}
	card := []string{}
	for _, cardline := range input[2:] {
		if cardline == "" {
			bp.Cards = append(bp.Cards, NewBingoCardFromString(card))
			card = []string{}
			continue
		}
		card = append(card, cardline)
	}
	bp.Cards = append(bp.Cards, NewBingoCardFromString(card))
	for _, number := range numbers {
		candidate, err := strconv.Atoi(number)
		if err != nil {
			candidate = 0
		}
		if result, ok := bp.CheckLast(candidate); ok {
			return result * candidate
		}
	}
	return 0
}

// NewBingoCardFromString takes a string that is
// 5x5 in size and parses it into a bingo card struct.
func NewBingoCardFromString(a []string) BingoCard {
	var bcRows, bcCols [5][5]int
	marked := make(map[int]bool)

	for idx, col := range a {
		for idy, num := range strings.Fields(col) {
			val, err := strconv.Atoi(num)
			if err != nil {
				val = 0
			}
			bcCols[idx][idy] = val
			bcRows[idy][idx] = val
			marked[val] = false
		}
	}

	return BingoCard{Rows: bcRows, Cols: bcCols, Marked: marked}
}
