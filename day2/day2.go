package day2

import (
	"strconv"
	"strings"
)

// Mover let's you move forward or up/down.
type Mover interface {
	Forward(length int)
	Down(length int)
	Up(length int)
	Coordinates() int
}

// Submarine tracks the position of your submarine.
type Submarine struct {
	HorizontalPos int
	Depth         int
}

// Forward increases the length travelled horizontally.
func (s *Submarine) Forward(length int) {
	s.HorizontalPos += length
}

// Down increases your submarine in depth.
func (s *Submarine) Down(length int) {
	s.Depth += length
}

// Up decreases your submraine in depth.
func (s *Submarine) Up(length int) {
	s.Depth -= length
}

// Coordinates returns your location based on
// horizontal postition multiplied by your depth.
func (s *Submarine) Coordinates() int {
	return s.HorizontalPos * s.Depth
}

// SophisticatedSubmarine tracks the position of your submarine
// by accounting for the aim as described in the manual.
type SophisticatedSubmarine struct {
	HorizontalPos int
	Depth         int
	Aim           int
}

// Forward increases the length travelled horizontally.
func (s *SophisticatedSubmarine) Forward(length int) {
	s.HorizontalPos += length
	s.Depth += s.Aim * length
}

// Down increases your submarine aim.
func (s *SophisticatedSubmarine) Down(length int) {
	s.Aim += length
}

// Up decreases your submraine aim.
func (s *SophisticatedSubmarine) Up(length int) {
	s.Aim -= length
}

// Coordinates returns your location based on
// horizontal postition multiplied by your depth.
func (s *SophisticatedSubmarine) Coordinates() int {
	return s.HorizontalPos * s.Depth
}

// CalculateCoordinates calculates your position
// after some mover commands (forward,up,down).
func CalculateCoordinates(m Mover, input []string) int {
	for _, command := range input {
		tokens := strings.Split(command, " ")
		verb, distance := tokens[0], tokens[1]
		length, err := strconv.Atoi(distance)
		if err != nil {
			length = 0
		}
		switch verb {
		case "forward":
			m.Forward(length)
		case "down":
			m.Down(length)
		case "up":
			m.Up(length)
		}
	}
	return m.Coordinates()
}
