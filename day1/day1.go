package day1

import "strconv"

// MeasureIncrease takes a list of strings and coverts them to integers
// in order to measure the increase between the current and the previous
// integer and returns the overall count of increases in the sequence.
func MeasureIncrease(input []string) int {
	var prev int
	increases := -1
	for _, depth := range input {
		depthInt, err := strconv.Atoi(depth)
		if err != nil {
			return 0
		}
		if depthInt > prev {
			increases++
		}
		prev = depthInt
	}
	return increases
}

// ThreeWindowMeasureIncrease takes a sequence of strings and calculates whether
// the measure from one window to the other increases, retuning the overal
// increase once there are no more three window sections in the sequence.
func ThreeWindowMeasureIncrease(input []string) int {
	var prevWindow, depthWindow int
	increases := -1
	for idx := range input {
		// determine depthWindow by looking ahead
		if idx+2 < len(input) {
			depthWindow = sum(input[idx], input[idx+1], input[idx+2])
			if depthWindow > prevWindow {
				increases++
			}
			prevWindow = depthWindow
		} else {
			break
		}
	}
	return increases
}

func sum(input ...string) int {
	var result int
	for _, val := range input {
		partial, err := strconv.Atoi(val)
		if err != nil {
			partial = 0
		}
		result += partial
	}
	return result
}
