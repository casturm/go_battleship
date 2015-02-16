package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

type Ship struct {
	start     Point
	direction string
	size      int
}

type Board struct {
	ships []Ship
}

func (s *Ship) Location() []Point {
	loc := make([]Point, s.size)
	loc[0] = s.start
	for i := 1; i < s.size; i++ {
		loc[i] = Point{s.start.x + i, s.start.y}
	}
	return loc
}

func (b *Board) Fire(f Point) bool {
	for _, s := range b.ships {
		for _, p := range s.Location() {
			if f == p {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println("starting")

	s1 := Ship{Point{1, 2}, "right", 4}
	ships := [1]Ship{s1}
	b := Board{ships[0:]}

	fmt.Println("s1: ", s1)
	fmt.Println("location(s1)", s1.Location())
	fmt.Println("Fire 4,3 Hit?", b.Fire(Point{4, 3}))
	fmt.Println("Fire 4,2 Hit?", b.Fire(Point{4, 2}))
}
