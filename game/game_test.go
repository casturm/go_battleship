package game

import (
	"testing"
)

func TestMakeShipUp(t *testing.T) {
	ship := MakeShip(Point{1, 2}, "up", 3)
	if len(ship.Location) != 3 {
		t.Error("Expected len 3")
	}
	if (ship.Location[0] != Point{1, 2}) {
		t.Error("wrong point at index 0", ship.Location)
	}
	if (ship.Location[1] != Point{1, 1}) {
		t.Error("wrong point at index 1", ship.Location)
	}
	if (ship.Location[2] != Point{1, 0}) {
		t.Error("wrong point at index 2", ship.Location)
	}
	if len(ship.Hits) != 0 {
		t.Error("len Hits is not 0", ship.Hits)
	}
}

func TestNewPlayer(t *testing.T) {
	player := NewPlayer("testplayer")
	if player.Name != "testplayer" {
		t.Error("Expected testplayer for Name")
	}
}

func TestShipHit(t *testing.T) {
	for hy := 0; hy < 10; hy++ {
		for hx := 0; hx < 10; hx++ {
			hit := Point{hx, hy}
			ship := Ship{[]Point{hit}, []Point{}}
			for y := 0; y < 10; y++ {
				for x := 0; x < 10; x++ {
					p := Point{x, y}
					if hit == p && !ship.Hit(p) {
						t.Error("Expected hit at", x, y)
					}
					if hit != p && ship.Hit(p) {
						t.Error("Expected no hit at", x, y)
					}
				}
			}
		}
	}
}

func TestShipSunk(t *testing.T) {
	p := Point{1, 2}
	ship := Ship{[]Point{p}, []Point{}}
	if ship.Sunk() {
		t.Error("Ship should not be sunk")
	}

	ship = Ship{[]Point{p}, []Point{p}}
	if !ship.Sunk() {
		t.Error("Ship should be sunk")
	}
}

func TestPlayerFire(t *testing.T) {
	player := Player{"", "", []*Ship{MakeShip(Point{2, 4}, "right", 5)}, []Point{}}

	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			p := Point{x, y}
			isShip := pointOnShip(p, player)
			if isShip && (player.Fire(p) == nil) {
				t.Error("Expected hit at", x, y)
			}
			if !isShip && (player.Fire(p) != nil) {
				t.Error("Expected no hit at", x, y)
			}
		}
	}
}

func pointOnShip(a Point, player Player) bool {
	for _, ship := range player.Ships {
		for _, b := range ship.Location {
			if b == a {
				return true
			}
		}
	}
	return false
}
