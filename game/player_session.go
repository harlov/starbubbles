package game

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
)

type PlayerSession struct {
	Socket    *websocket.Conn
	Player    *Player
	Game      *Game
	Direction Cordinate
	Balls     []*Ball
	Score     float32
}

type CommandParams map[string]interface{}

type PlayerSessionCommand struct {
	Method string
	Params CommandParams
}

func NewSession(ws *websocket.Conn, player *Player, game *Game) *PlayerSession {
	ps := PlayerSession{Socket: ws, Player: player, Game: game}
	go ps.receiver()
	return &ps
}

func (ps *PlayerSession) receiver() {
	for {
		_, command, err := ps.Socket.ReadMessage()
		if err != nil {
			break
		}
		var player_command PlayerSessionCommand = PlayerSessionCommand{}
		err = json.Unmarshal(command, &player_command)

		ps.serverCommand(player_command)
	}
	ps.Socket.Close()
}

func (ps *PlayerSession) serverCommand(command PlayerSessionCommand) {
	log.Printf("method : %s \n", command.Method)
}

func (ps *PlayerSession) sendClientCommand(command PlayerSessionCommand) {
	message_b, _ := json.Marshal(command)
	err := ps.Socket.WriteMessage(websocket.TextMessage, message_b)
	if err != nil {
		ps.Socket.Close()
	}
}

func (ps *PlayerSession) moveBalls() []*Ball {
	for _, ball := range ps.Balls {
		ball.move(ps.Direction)
	}
	return ps.Balls
}
