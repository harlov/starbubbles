package game

import (
	"log"
)

const frameRate = 3.0

type Game struct {
	Name       string
	Sessions   []*PlayerSession
	JoinChanel chan *PlayerSession
	Field  *Field
}

func NewGame(name string) *Game {
	game := Game{
		Name:       name,
		JoinChanel: make(chan *PlayerSession),
		Field : NewField(1024, 1024),
	}
	go game.runListen()
	go game.loop()
	return &game
}

func (g *Game) runListen() {
	for {
		p_sess := <-g.JoinChanel
		g.Sessions = append(g.Sessions, p_sess)
		log.Printf("player %s has joined\n", p_sess.Player.Name)
	}
}

func (g *Game) loop() {
	//TODO: collisions detect
}

func (g Game) JoinPlayer(session *PlayerSession) {
	g.JoinChanel <- session
}
