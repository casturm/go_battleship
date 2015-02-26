package game

type Point struct {
	X int
	Y int
}

type Ship struct {
	Size     int
	Location []Point
	Hits     []Point
}

type Player struct {
	Id     string
	Name   string
	Ships  []*Ship
	Misses []Point
}

func location(size int, direction string, x int, y int) []Point {
	loc := make([]Point, size)
	switch direction {
	case "right":
		for i := 0; i < size; i++ {
			loc[i] = Point{x + i, y}
		}
	case "left":
		for i := 0; i < size; i++ {
			loc[i] = Point{x - i, y}
		}
	case "up":
		for i := 0; i < size; i++ {
			loc[i] = Point{x, y - i}
		}
	case "down":
		for i := 0; i < size; i++ {
			loc[i] = Point{x, y + i}
		}
	}
	return loc
}

func (s *Ship) Hit(f Point) bool {
	for _, p := range s.Location {
		if f == p {
			s.Hits = append(s.Hits, f)
			return true
		}
	}
	return false
}

func (s *Ship) Sunk() bool {
	return len(s.Hits) == s.Size
}

func (p *Player) Fire(f Point) *Ship {
	for _, s := range p.Ships {
		if s.Hit(f) {
			return s
		}
	}
	p.Misses = append(p.Misses, f)
	return nil
}

func (p *Player) GameOver() bool {
	for _, s := range p.Ships {
		if s.Sunk() == false {
			return false
		}
	}
	return true
}

func MakeShip(nose Point, direction string, size int) Ship {
	location := location(size, direction, nose.X, nose.Y)
	return Ship{size, location, make([]Point, 0)}
}
