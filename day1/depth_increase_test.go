package depthincrease_test

import (
	"testing"

	depth "github.com/merkata/go-aoc-2021/day1/depthincrease/v0.1.2"
)

func TestDepthIncrease(t *testing.T) {
	input := []string{
		"199",
		"200",
		"208",
		"210",
		"200",
		"207",
		"240",
		"269",
		"260",
		"263",
	}
	want := 7
	got := depth.MeasureIncrease(input)
	if want != got {
		t.Errorf("failed test, wanted %d got %d\n", want, got)
	}
}
