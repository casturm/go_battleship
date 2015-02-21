package main

import (
	"battleship/game"
	"encoding/json"
	"fmt"
	"github.com/go-martini/martini"
	"html/template"
	"io/ioutil"
	"net/http"
)

type PlayerTurn struct {
	X      int
	Y      int
	Player int
}

type Game struct {
	Turn    int
	Size    int
	Players []*game.Player
}

func newGame() Game {
	fmt.Println("Starting a New Game!")

	p1s1 := game.MakeShip(game.Point{1, 2}, "right", 4)
	p1s2 := game.MakeShip(game.Point{6, 1}, "down", 3)
	p1ships := [2]*game.Ship{&p1s1, &p1s2}

	p2s1 := game.MakeShip(game.Point{3, 2}, "right", 3)
	p2s2 := game.MakeShip(game.Point{6, 8}, "up", 4)
	p2ships := [2]*game.Ship{&p2s1, &p2s2}

	p1 := game.Player{p1ships[0:], make([]game.Point, 0)}
	p2 := game.Player{p2ships[0:], make([]game.Point, 0)}

	players := [2]*game.Player{&p1, &p2}
	//players[0] = &game.Player{&p1b}
	//players[1] = &game.Player{&p2b}

	return Game{0, 10, players[0:]}
}

func (g *Game) RunGameLoop(x int, y int) string {
	fmt.Println("game:", g)
	for p, player := range g.Players {
		fmt.Println("player", p, player)
	}

	ship := g.Players[g.Turn].Fire(game.Point{x, y})
	g.Turn = (g.Turn + 1) % 2
	if ship != nil {
		return "hit"
	} else {
		return "miss"
	}
}

func (g *Game) gameOver() bool {
	for i, player := range g.Players {
		if player.GameOver() {
			fmt.Printf("\nPlayer %v Won!\n\n", (i+1)%2)
			fmt.Println("GAME OVER!")
			return true
		}
	}
	return false
}

var index = template.Must(template.ParseFiles(
	"templates/_base.html",
	"templates/index.html",
))

func main() {
	var thisGame Game
	m := martini.Classic()
	static := martini.Static("client/app", martini.StaticOptions{Fallback: "/index.html", Exclude: "/api/v"})
	m.NotFound(static, http.NotFound)

	m.Get("/hello", func(w http.ResponseWriter, req *http.Request) {
		if err := index.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	m.Post("/api/game", func() (int, string) {
		thisGame = newGame()
		fmt.Println("new game:", thisGame)
		gameJson, err := json.Marshal(&thisGame)
		if err != nil {
			// Could not marshal ship.
			return http.StatusInternalServerError, "failed to marshal ship to JSON data"
		}

		return 200, string(gameJson)
	})
	m.Post("/api/turn", func(params martini.Params, req *http.Request) (int, string) {
		defer req.Body.Close()

		fmt.Println("post /turn", params)

		// Read request body.
		requestBody, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return http.StatusInternalServerError, "internal error"
		}

		fmt.Println("requestBody", string(requestBody))

		if len(params) != 0 {
			// No keys in params. This is not supported.
			return http.StatusMethodNotAllowed, "method not allowed"
		}

		// Unmarshal entry sent by the user.
		var playerTurn PlayerTurn
		err = json.Unmarshal(requestBody, &playerTurn)
		if err != nil {
			// Could not unmarshal entry.
			return http.StatusBadRequest, "invalid JSON data"
		}
		fmt.Println("fire! ", playerTurn)

		fmt.Println("runGameLoop on", thisGame)
		result := thisGame.RunGameLoop(playerTurn.X, playerTurn.Y)

		//response, err := json.Marshal(&result)
		//if err != nil {
		//Could not marshal ship.
		//return http.StatusInternalServerError, "failed to marshal ship to JSON data"
		//}

		return 200, result
	})
	m.Run()
}
