package game

import (
	"fmt"
	//"sync"
)

type DB interface {
	FindAllPlayers() []*Player
	FindPlayer(id string) *Player
	FindGame(id string) *Game
	SaveGame(game *Game) error
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
