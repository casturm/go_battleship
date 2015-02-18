package main

import (
	"fmt"
	"strconv"
	"strings"
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
	size  int
	ships []*Ship
}

func (s *Ship) Location() []Point {
	loc := make([]Point, s.size)
	switch s.direction {
	case "right":
		for i := 0; i < s.size; i++ {
			loc[i] = Point{s.start.x + i, s.start.y}
		}
	case "left":
		for i := 0; i < s.size; i++ {
			loc[i] = Point{s.start.x - i, s.start.y}
		}
	case "up":
		for i := 0; i < s.size; i++ {
			loc[i] = Point{s.start.x, s.start.y + i}
		}
	case "down":
		for i := 0; i < s.size; i++ {
			loc[i] = Point{s.start.x, s.start.y - i}
		}
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

func (b *Board) Fire(f Point) *Ship {
	for _, s := range b.ships {
		if s.Hit(f) {
			return s
		}
	}
	return nil
}

func (b *Board) GameOver() bool {
	for _, s := range b.ships {
		if s.Sunk() == false {
			return false
		}
	}
	return true
}

func sunk(s *Ship) string {
	if s != nil {
		return fmt.Sprintf("Hit! total hits: %v, sunk? %v", s.hits, s.Sunk())
	} else {
		return fmt.Sprintf("Miss!")
	}
}

func main() {
	fmt.Println("starting")

	s1 := Ship{Point{1, 2}, "right", 4, make([]Point, 0, 4)}
	s2 := Ship{Point{6, 9}, "down", 5, make([]Point, 0, 5)}
	//s3 := Ship{Point{8, 12}, "left", 4, make([]Point, 0, 4)}
	//s4 := Ship{Point{5, 4}, "up", 3, make([]Point, 0, 3)}
	//s5 := Ship{Point{9, 2}, "right", 2, make([]Point, 0, 2)}

	ships := [2]*Ship{&s1, &s2} //, &s3, &s4, &s5}

	for i, s := range ships {
		fmt.Printf("s%v: location: %v\n", i+1, s.Location())
	}

	//for i := 0; i < 4; i++ {
	//f := Point{1 + i, 2}
	//fmt.Println("board.Fire", f, "hit?", b.Fire(f))
	//fmt.Println("s1:", s1, "sunk?", s1.Sunk())
	//}

	b := Board{25, ships[0:]}
	var input string

	for {
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Error: ", err)
		} else {
			//fmt.Println(n, " ", input)
			loc := strings.Split(input, ",")
			x, _ := strconv.Atoi(loc[0])
			y, _ := strconv.Atoi(loc[1])
			ship := b.Fire(Point{x, y})
			fmt.Printf("Fired at (%v,%v) ... %v\n", x, y, sunk(ship))
		}
		if b.GameOver() {
			fmt.Println("GAME OVER!")
			break
		}
	}
}
