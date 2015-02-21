package main

import (
	"battleship/game"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func result(s *game.Ship) string {
	if s != nil {
		return fmt.Sprintf("Hit! sunk? %v\n", s.Sunk())
	} else {
		return fmt.Sprintf("Miss!\n")
	}
}

func display(players []*game.Player) {
	for _, p := range players {
		for _, s := range p.Board.Ships {
			fmt.Printf("ship: %v hits: %v\n", s.Location(), s.Hits)
		}
		fmt.Printf("misses: %v\n", p.Board.Misses)
	}
}

func main() {
	fmt.Println("starting")

	p1s1 := game.MakeShip(game.Point{1, 2}, "right", 4)
	p1s2 := game.MakeShip(game.Point{6, 9}, "down", 5)
	p1ships := [2]*game.Ship{&p1s1, &p1s2}

	for i, s := range p1ships {
		fmt.Printf("p1s%v: location: %v\n", i+1, s.Location())
	}

	p2s1 := game.MakeShip(game.Point{9, 2}, "right", 4)
	p2s2 := game.MakeShip(game.Point{5, 4}, "down", 4)
	p2ships := [2]*game.Ship{&p2s1, &p2s2}

	for i, s := range p2ships {
		fmt.Printf("p2s%v: location: %v\n", i+1, s.Location())
	}

	p1b := game.Board{25, p1ships[0:], make([]game.Point, 0)}
	p2b := game.Board{25, p2ships[0:], make([]game.Point, 0)}

	var players []*game.Player
	players = make([]*game.Player, 2)
	players[0] = &game.Player{&p1b}
	players[1] = &game.Player{&p2b}

	var input string
	var x, y int
	var player *game.Player
	turn := 0

	rand := rand.New(rand.NewSource(25))

	for {
		player = players[turn]

		if turn == 0 {
			_, err := fmt.Scanln(&input)
			if err != nil {
				fmt.Println("Error: ", err)
			} else {
				loc := strings.Split(input, ",")
				x, _ = strconv.Atoi(loc[0])
				y, _ = strconv.Atoi(loc[1])
			}
		} else {
			x, y = rand.Intn(25), rand.Intn(25)
		}

		turn = (turn + 1) % 2
		player = players[turn]

		ship := player.Board.Fire(game.Point{x, y})
		fmt.Printf("Fired at (%v,%v) ... %v\n", x, y, result(ship))
		display(players)

		if player.Board.GameOver() {
			turn = (turn + 1) % 2
			fmt.Printf("\nPlayer %v Won!\n\n", turn)
			fmt.Println("GAME OVER!")
			break
		}
	}
}
