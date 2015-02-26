package main

import (
	"battleship/game"
	"code.google.com/p/go-uuid/uuid"
	"encoding/json"
	"fmt"
	"github.com/go-martini/martini"
	"io/ioutil"
	"net/http"
)

type PlayerTurn struct {
	gameId string
	X      int
	Y      int
	Player int
}

type Game struct {
	Id      string
	Turn    int
	Size    int
	Players []*game.Player
	State   string
}

func newGame(player1, player2 *game.Player) Game {
	fmt.Println("Starting a New Game!")

	//p1s1 := game.MakeShip(game.Point{1, 2}, "right", 4)
	//p1s2 := game.MakeShip(game.Point{6, 1}, "down", 3)
	//p1ships := [2]*game.Ship{&p1s1, &p1s2}

	//p2s1 := game.MakeShip(game.Point{3, 2}, "right", 3)
	//p2s2 := game.MakeShip(game.Point{6, 8}, "up", 4)
	//p2ships := [2]*game.Ship{&p2s1, &p2s2}

	player1.Misses = make([]game.Point, 0, 0)
	player2.Misses = make([]game.Point, 0, 0)
	gamePlayers := [2]*game.Player{player1, player2}

	return Game{uuid.New(), 0, 10, gamePlayers[0:], "new"}
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

func main() {
	var games = make(map[string]Game)
	var players = make(map[string]*game.Player)

	m := martini.Classic()
	static := martini.Static("ui/app", martini.StaticOptions{Fallback: "/index.html", Exclude: "/api/v"})
	m.NotFound(static, http.NotFound)

	m.Post("/api/ship", func(params martini.Params, req *http.Request) (int, string) {
		defer req.Body.Close()

		fmt.Println("post /ship params", params)

		requestBody, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return http.StatusInternalServerError, "internal error"
		}

		fmt.Println("post /ship requestBody", string(requestBody))

		type postShip struct {
			Player    string
			LocationX int
			LocationY int
			Direction string
			Size      int
		}
		shipRequest := new(postShip)
		err = json.Unmarshal(requestBody, shipRequest)
		if err != nil {
			return http.StatusBadRequest, "invalid JSON data"
		}
		fmt.Println("new shipRequest", shipRequest)

		player := players[shipRequest.Player]
		point := game.Point{shipRequest.LocationX, shipRequest.LocationY}
		ship := game.MakeShip(point, shipRequest.Direction, shipRequest.Size)
		player.Ships = append(player.Ships, &ship)

		shipJson, err := json.Marshal(ship)
		if err != nil {
			return http.StatusInternalServerError, "failed to marshal result to JSON data"
		}

		return 200, string(shipJson)
	})
	m.Post("/api/player", func(params martini.Params, req *http.Request) (int, string) {
		defer req.Body.Close()

		fmt.Println("post /player params", params)

		// Read request body.
		requestBody, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return http.StatusInternalServerError, "internal error"
		}

		fmt.Println("post /player requestBody", requestBody)

		// Unmarshal entry sent by the user.
		var player = new(game.Player)
		err = json.Unmarshal(requestBody, player)
		if err != nil {
			return http.StatusBadRequest, "invalid JSON data"
		}
		fmt.Println("new Player", player)

		uuid := uuid.New()
		player.Id = uuid
		players[uuid] = player

		playerJson, err := json.Marshal(player)
		if err != nil {
			return http.StatusInternalServerError, "failed to marshal players to JSON data"
		}

		return 200, string(playerJson)
	})
	m.Get("/api/player", func(params martini.Params, req *http.Request) (int, string) {
		defer req.Body.Close()

		list := make([]*game.Player, 0, 0)
		for key, value := range players {
			fmt.Println("Key:", key, "Value:", value)
			list = append(list, value)
		}
		playersJson, err := json.Marshal(list)
		if err != nil {
			return http.StatusInternalServerError, "failed to marshal players to JSON data"
		}

		return 200, string(playersJson)
	})
	//m.Get("/api/game", func() (int, string) {
	//fmt.Println("get game:", thisGame)
	//gameJson, err := json.Marshal(&thisGame)
	//if err != nil {
	//return http.StatusInternalServerError, "failed to marshal game to JSON data"
	//}

	//return 200, string(gameJson)
	//})
	m.Post("/api/game", func(params martini.Params, req *http.Request) (int, string) {
		defer req.Body.Close()

		fmt.Println("post /game params", params)

		// Read request body.
		requestBody, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return http.StatusInternalServerError, "internal error"
		}

		fmt.Println("post /game requestBody", string(requestBody))

		// Unmarshal players
		type newGamePlayers struct {
			Player1, Player2 string
		}

		gamePlayers := new(newGamePlayers)
		err = json.Unmarshal(requestBody, gamePlayers)
		if err != nil {
			return http.StatusBadRequest, "invalid JSON data"
		}
		fmt.Println("new game with ", gamePlayers)

		thisGame := newGame(players[gamePlayers.Player1], players[gamePlayers.Player2])
		fmt.Println("new game:", thisGame)
		gameJson, err := json.Marshal(&thisGame)
		if err != nil {
			return http.StatusInternalServerError, "failed to marshal game to JSON data"
		}

		return 200, string(gameJson)
	})
	m.Post("/api/turn", func(params martini.Params, req *http.Request) (int, string) {
		defer req.Body.Close()

		fmt.Println("post /turn", params)

		requestBody, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return http.StatusInternalServerError, "internal error"
		}

		fmt.Println("requestBody", string(requestBody))
		// Unmarshal entry sent by the user.
		var playerTurn PlayerTurn
		err = json.Unmarshal(requestBody, &playerTurn)
		if err != nil {
			// Could not unmarshal entry.
			return http.StatusBadRequest, "invalid JSON data"
		}

		thisGame := games[playerTurn.gameId]
		if thisGame.GameOver() {
			fmt.Println("Game Over!")
			return 403, "Game Over"
		} else {
			result := "Not Your Turn"
			if playerTurn.Player == thisGame.Turn {
				fmt.Println("fire! ", playerTurn)

				fmt.Println("runGameLoop on", thisGame)
				result = thisGame.RunGameLoop(playerTurn.X, playerTurn.Y)
			}

			//response, err := json.Marshal(&result)
			//if err != nil {
			//Could not marshal ship.
			//return http.StatusInternalServerError, "failed to marshal ship to JSON data"
			//}

			return 200, result
		}
	})
	m.Run()
}
