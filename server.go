package main

import (
	"battleship/api"
	"battleship/game"
	"github.com/go-martini/martini"
	"net/http"
)

func main() {
	m := martini.Classic()
	static := martini.Static("ui/app", martini.StaticOptions{Fallback: "/index.html", Exclude: "/api/v"})
	m.NotFound(static, http.NotFound)

	m.Post("/api/player", api.AddPlayer)
	m.Get("/api/players", api.GetPlayers)
	m.Post("/api/game", api.MakeGame)
	m.Get("/api/games", api.GetGames)
	m.Get("/api/game/:id", api.GetGame)
	m.Post("/api/ship", api.AddShip)
	m.Post("/api/turn", api.TakeTurn)

	// Inject database
	m.MapTo(game.TheDB, (*game.DB)(nil))
	m.Run()
}
