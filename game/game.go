package game

import (
	"code.google.com/p/go-uuid/uuid"
	"fmt"
)

type Point struct {
	X int
	Y int
}

type Ship struct {
	Location []Point
	Hits     []Point
}

type Player struct {
	Id     string
	Name   string
	Ships  []*Ship
	Misses []Point
}

type Game struct {
	Id      string
	Turn    int
	Size    int
	Players []*Player
	State   string
}

func MakeShip(nose Point, direction string, size int) *Ship {
	location := location(size, direction, nose.X, nose.Y)
	var ship = new(Ship)
	ship.Location = location
	ship.Hits = make([]Point, 0)
	return ship
}

func NewPlayer(name string) *Player {
	player := new(Player)
	player.Id = uuid.New()
	player.Name = name
	return player
}

func NewGame(player1, player2 *Player) Game {
	fmt.Println("New Game!")

	player1.Misses = make([]Point, 0, 0)
	player2.Misses = make([]Point, 0, 0)
	gamePlayers := [2]*Player{player1, player2}

	game := Game{uuid.New(), 0, 10, gamePlayers[0:], "new"}

	return game
}

func (s *Ship) Hit(f Point) bool {
	for _, p := range s.Location {
		if f == p {
			return true
		}
	}
	return false
}

func (s *Ship) Sunk() bool {
	return len(s.Hits) == len(s.Location)
}

func (p *Player) Fire(f Point) *Ship {
	for _, s := range p.Ships {
		if s.Hit(f) {
			s.Hits = append(s.Hits, f)
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

func (p *Player) AddShip(point Point, direction string, size int) *Ship {
	ship := MakeShip(point, direction, size)
	p.Ships = append(p.Ships, ship)
	return ship
}

func (g *Game) TakeTurn(x, y int) string {
	fmt.Println("take turn")
	for p, player := range g.Players {
		fmt.Println("player", p, player)
	}

	ship := g.Players[g.Turn].Fire(Point{x, y})
	g.Turn = (g.Turn + 1) % 2
	if ship != nil {
		return "hit"
	} else {
		return "miss"
	}
}

func (g *Game) GameOver() bool {
	for i, player := range g.Players {
		if player.GameOver() {
			fmt.Printf("\nPlayer %v Won!\n\n", (i+1)%2)
			fmt.Println("GAME OVER!")
			g.State = "game over"
			return true
		}
	}
	return false
}

func location(size int, direction string, x int, y int) []Point {
	loc := make([]Point, size)
	switch direction {
	case "right":
		for i := 0; i < size; i++ {
			loc[i] = Point{x + i, y}
		}
	case "up":
		for i := 0; i < size; i++ {
			loc[i] = Point{x, y - i}
		}
	}
	return loc
}
