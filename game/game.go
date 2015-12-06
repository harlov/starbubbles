package game

import "log"

type Game struct {
    Name string
    Sessions []*PlayerSession
    JoinChanel chan *PlayerSession
}

func NewGame(name string) *Game {
    game := Game{
        Name:name,
        JoinChanel : make(chan *PlayerSession),
    }
    go game.runListen()
    return &game
}


func (g Game) runListen() {
    for {
        p_sess := <-g.JoinChanel
        g.Sessions = append(g.Sessions, p_sess)
        log.Printf("player %s has joined\n", p_sess.Player.Name)
    }
}

func (g Game) JoinPlayer(session *PlayerSession) {
    g.JoinChanel <- session
}

