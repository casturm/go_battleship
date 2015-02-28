package api

import (
	"battleship/game"
	"encoding/json"
	"fmt"
	"github.com/go-martini/martini"
	"io/ioutil"
	"net/http"
)

func GetGame(params martini.Params, req *http.Request, db game.DB) (int, string) {
	game := db.FindGame(params["id"])
	gameJson, err := json.Marshal(game)
	if err != nil {
		return http.StatusInternalServerError, "failed to marshal result to JSON data"
	}

	return 200, string(gameJson)
}

func AddShip(params martini.Params, req *http.Request, db game.DB) (int, string) {
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

	player := db.FindPlayer(shipRequest.Player)
	point := game.Point{shipRequest.LocationX, shipRequest.LocationY}
	ship := game.MakeShip(point, shipRequest.Direction, shipRequest.Size)
	player.Ships = append(player.Ships, ship)

	shipJson, err := json.Marshal(ship)
	if err != nil {
		return http.StatusInternalServerError, "failed to marshal result to JSON data"
	}

	return 200, string(shipJson)
}

func AddPlayer(params martini.Params, req *http.Request, db game.DB) (int, string) {
	defer req.Body.Close()

	fmt.Println("post /player params", params)

	// Read request body.
	requestBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return http.StatusInternalServerError, "internal error"
	}

	fmt.Println("post /player requestBody", string(requestBody))

	type NewPlayerRequest struct {
		Name string
	}
	playerRequest := new(NewPlayerRequest)
	err = json.Unmarshal(requestBody, &playerRequest)
	if err != nil {
		return http.StatusBadRequest, "invalid JSON data"
	}
	fmt.Println("NewPlayerRequest", playerRequest)

	player := game.NewPlayer(playerRequest.Name)
	fmt.Println("new player", player)
	db.SavePlayer(player)

	playerJson, err := json.Marshal(player)
	if err != nil {
		return http.StatusInternalServerError, "failed to marshal players to JSON data"
	}

	return 200, string(playerJson)
}

func GetPlayers(params martini.Params, req *http.Request, db game.DB) (int, string) {
	defer req.Body.Close()

	list := db.FindAllPlayers()
	playersJson, err := json.Marshal(list)
	if err != nil {
		return http.StatusInternalServerError, "failed to marshal players to JSON data"
	}

	return 200, string(playersJson)
}

func AddGame(params martini.Params, req *http.Request, db game.DB) (int, string) {
	defer req.Body.Close()

	fmt.Println("post /game params", params)

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

	thisGame := game.NewGame(db.FindPlayer(gamePlayers.Player1), db.FindPlayer(gamePlayers.Player2))
	fmt.Println("new game:", thisGame)
	db.SaveGame(&thisGame)
	gameJson, err := json.Marshal(&thisGame)
	if err != nil {
		return http.StatusInternalServerError, "failed to marshal game to JSON data"
	}

	return 200, string(gameJson)
}

func TakeTurn(params martini.Params, req *http.Request, db game.DB) (int, string) {
	defer req.Body.Close()

	fmt.Println("post /turn", params)

	requestBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return http.StatusInternalServerError, "internal error"
	}

	fmt.Println("requestBody", string(requestBody))

	var playerTurn struct {
		gameId string
		X      int
		Y      int
		Player int
	}
	err = json.Unmarshal(requestBody, &playerTurn)
	if err != nil {
		return http.StatusBadRequest, "invalid JSON data"
	}

	thisGame := db.FindGame(playerTurn.gameId)
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
		return 200, result
	}
}
