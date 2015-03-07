package game

import (
	"fmt"
	//"sync"
)

type DB interface {
	FindAllPlayers() []*Player
	FindPlayer(id string) *Player
	FindPlayerByName(name string) (*Player, bool)
	FindGame(id string) *Game
	SaveGame(game *Game) error
	FindAllGames() []*Game
	SavePlayer(player *Player) error
}

type gameDB struct {
	//sync.RWMutex
	g map[string]*Game
	p map[string]*Player
}

var TheDB DB

func init() {
	TheDB = &gameDB{
		g: make(map[string]*Game),
		p: make(map[string]*Player),
	}
}

func (TheDB *gameDB) FindAllPlayers() []*Player {
	fmt.Println("find all players")
	//TheDB.RLock()
	//defer TheDB.Unlock()
	var all = make([]*Player, 0, 0)
	for _, value := range TheDB.p {
		all = append(all, value)
	}
	return all
}

func (TheDB *gameDB) FindPlayerByName(name string) (*Player, bool) {
	fmt.Println("find player by name", name)
	//TheDB.RLock()
	//defer TheDB.Unlock()
	var player = new(Player)
	for _, value := range TheDB.p {
		if value.Name == name {
			return value, true
		}
	}
	return player, false
}

func (TheDB *gameDB) FindPlayer(id string) *Player {
	//TheDB.RLock()
	//defer TheDB.Unlock()
	return TheDB.p[id]
}

func (TheDB *gameDB) FindGame(id string) *Game {
	//TheDB.RLock()
	//defer TheDB.Unlock()
	return TheDB.g[id]
}

func (TheDB *gameDB) FindAllGames() []*Game {
	fmt.Println("find all games")
	//TheDB.RLock()
	//defer TheDB.Unlock()
	var all = make([]*Game, 0, 0)
	for _, value := range TheDB.g {
		all = append(all, value)
	}
	return all
}
func (TheDB *gameDB) SaveGame(g *Game) error {
	//TheDB.RLock()
	//defer TheDB.Unlock()
	TheDB.g[g.Id] = g
	return nil
}

func (TheDB *gameDB) SavePlayer(p *Player) error {
	fmt.Println("save new player in db", p)
	//TheDB.RLock()
	//defer TheDB.Unlock()
	TheDB.p[p.Id] = p
	return nil
}
