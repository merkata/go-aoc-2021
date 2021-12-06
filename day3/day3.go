package day3

import (
	"strconv"
	"strings"
)

// CalculatePowerConsumption takes an input of bits
// and derives gamma and epsilon rates, returning their product.
func CalculatePowerConsumption(input []string) int64 {
	gr, er := calculatePowerRates(input)
	return gr * er
}

func calculatePowerRates(input []string) (int64, int64) {
	var min, max int64
	lookup := make(map[byte]int)
	var minstr, maxstr strings.Builder
	for row := 0; row < len(input[0]); row++ {
		lookup[byte('0')] = 0
		lookup[byte('1')] = 0
		for col := 0; col < len(input); col++ {
			lookup[input[col][row]]++
		}
		if lookup['0'] > lookup['1'] {
			maxstr.WriteByte('0')
			minstr.WriteByte('1')
		} else {
			maxstr.WriteByte('1')
			minstr.WriteByte('0')
		}
	}
	min, _ = strconv.ParseInt(minstr.String(), 2, 64)
	max, _ = strconv.ParseInt(maxstr.String(), 2, 64)
	if min > max {
		return max, min
	}
	return min, max
}

// CalculateLifeSupportRating is the result of multiplying
// calculated oxygen generator and CO2scribber ratings
func CalculateLifeSupportRating(input []string) int64 {
	og := calculateOxyGenRate(input)
	co := calculateCO2Rate(input)
	return og * co
}

func calculateOxyGenRate(input []string) int64 {
	var oxy int64
	lookup := make(map[byte]int)
	marksweepoxy := make(map[int]bool, len(input))
	for row := 0; row < len(input[0]); row++ {
		lookup[byte('0')] = 0
		lookup[byte('1')] = 0
		for col := 0; col < len(input); col++ {
			if marksweepoxy[col] {
				continue
			}
			marksweepoxy[col] = false
			lookup[input[col][row]]++
		}
		if lookup['0'] > lookup['1'] {
			sweep(input, marksweepoxy, row, byte('1'))
		} else if lookup['1'] >= lookup['0'] {
			sweep(input, marksweepoxy, row, byte('0'))
		}
		if oneid(marksweepoxy) {
			break
		}
	}
	id := nonsweeped(marksweepoxy)
	oxy, _ = strconv.ParseInt(input[id], 2, 64)
	return oxy
}

func calculateCO2Rate(input []string) int64 {
	var co2 int64
	lookup := make(map[byte]int)
	marksweepco2 := make(map[int]bool, len(input))
	for row := 0; row < len(input[0]); row++ {
		lookup[byte('0')] = 0
		lookup[byte('1')] = 0
		for col := 0; col < len(input); col++ {
			if marksweepco2[col] {
				continue
			}
			marksweepco2[col] = false
			lookup[input[col][row]]++
		}
		if lookup['0'] <= lookup['1'] {
			sweep(input, marksweepco2, row, byte('1'))
		} else if lookup['1'] < lookup['0'] {
			sweep(input, marksweepco2, row, byte('0'))
		}
		if oneid(marksweepco2) {
			break
		}
	}
	id := nonsweeped(marksweepco2)
	co2, _ = strconv.ParseInt(input[id], 2, 64)
	return co2
}

func sweep(input []string, mark map[int]bool, row int, char byte) {
	for idx, line := range input {
		if line[row] == char {
			mark[idx] = true
		}
	}
}

func nonsweeped(mark map[int]bool) int {
	for k, v := range mark {
		if !v {
			return k
		}
	}
	return 666
}

func oneid(mark map[int]bool) bool {
	var seen bool
	for _, v := range mark {
		if !v && seen {
			return false
		} else if !v {
			seen = true
		}
	}
	return seen
}
