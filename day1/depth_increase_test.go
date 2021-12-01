package depthincrease_test

import (
	"testing"

	"github.com/merkata/go-aoc-2021/day1/depthincrease"
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
	got := depthincrease.MeasureIncrease(input)
	if want != got {
		t.Errorf("failed test, wanted %d got %d\n", want, got)
	}
}
