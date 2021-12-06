package day5

import (
	"strconv"
	"strings"
)

// HydroBoard stores the x,y values of a board
// where lines cross.
type HydroBoard struct {
	Visited map[[2]int]int
}

// Visit updates a HydroBoard with new lines.
func (h *HydroBoard) Visit(from, to [2]int) {
	x1, y1 := from[0], from[1]
	x2, y2 := to[0], to[1]
	h.Visited[[2]int{x1, y1}]++
	for x1 != x2 || y1 != y2 {
		if x1 > x2 {
			x1--
		} else if x1 < x2 {
			x1++
		}
		if y1 > y2 {
			y1--
		} else if y1 < y2 {
			y1++
		}
		h.Visited[[2]int{x1, y1}]++
	}
}

// HydrothermalOverlap measures where hydrothermal lines
// overlap in a dangerous fashion (2 or more lines).
func HydrothermalOverlap(input []string, withDiagonal bool) int {
	visited := make(map[[2]int]int)
	hb := HydroBoard{Visited: visited}
	for _, coordinates := range input {
		if from, to, ok := calculate(coordinates, withDiagonal); ok {
			hb.Visit(from, to)
		}
	}
	var overlap int
	for _, v := range hb.Visited {
		if v >= 2 {
			overlap++
		}
	}
	return overlap
}

func calculate(coordinates string, withDiagonal bool) ([2]int, [2]int, bool) {
	xy := strings.Split(coordinates, " -> ")
	left := strings.Split(xy[0], ",")
	x1, _ := strconv.Atoi(left[0])
	y1, _ := strconv.Atoi(left[1])
	right := strings.Split(xy[1], ",")
	x2, _ := strconv.Atoi(right[0])
	y2, _ := strconv.Atoi(right[1])
	if !withDiagonal {
		if x1 == x2 || y1 == y2 {
			return [2]int{x1, y1}, [2]int{x2, y2}, true
		}
		return [2]int{x1, y1}, [2]int{x2, y2}, false
	}
	return [2]int{x1, y1}, [2]int{x2, y2}, true
}
