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
	hits      []Point
}

type Board struct {
	ships []*Ship
}

func (s *Ship) Location() []Point {
	loc := make([]Point, s.size)
	loc[0] = s.start
	// TODO check the other directions: default here is 'right'
	for i := 1; i < s.size; i++ {
		loc[i] = Point{s.start.x + i, s.start.y}
	}
	return loc
}

func (s *Ship) Hit(f Point) bool {
	for _, p := range s.Location() {
		if f == p {
			s.hits = append(s.hits, f)
			return true
		}
	}
	return false
}

func (s *Ship) Sunk() bool {
	return len(s.hits) == s.size
}

func (b *Board) Fire(f Point) bool {
	for _, s := range b.ships {
		if s.Hit(f) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("starting")

	s1 := Ship{Point{1, 2}, "right", 4, make([]Point, 0, 4)}
	ships := [1]*Ship{&s1}
	b := Board{ships[0:]}

	fmt.Println("s1: ", s1)
	fmt.Println("location(s1)", s1.Location())
	fmt.Println("Fire 4,3 Hit?", b.Fire(Point{4, 3}))
	for i := 0; i < 4; i++ {
		f := Point{1 + i, 2}
		fmt.Println("Fire", f, b.Fire(f))
		fmt.Println("s1:", s1)
		fmt.Println("sunk?", s1.Sunk())
	}
}
