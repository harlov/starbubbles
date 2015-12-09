package game

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"math"
	"time"
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
	ps.Balls = append(ps.Balls, NewBall(50, 100.0, 100.0))
	ps.Direction = Cordinate{X: 5, Y: 5}
	
	go ps.receiver()
	go ps.loop()
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
	switch command.Method {
	case "updateDirection" :
		ps.updateDirection(command)
	}
}

func (ps *PlayerSession) sendClientCommand(command PlayerSessionCommand) {
	message_b, _ := json.Marshal(command)
	err := ps.Socket.WriteMessage(websocket.TextMessage, message_b)
	if err != nil {
		ps.Socket.Close()
	}
}


func (ps *PlayerSession) loop() {
	var stepDelay int32 = 20
	for {
		time.Sleep(time.Duration(stepDelay) * time.Millisecond)
		ps.sendPlayerState()
	}
}


func (ps *PlayerSession) sendPlayerState() {
	balls_pos := ps.moveBalls()
	ps.sendClientCommand(PlayerSessionCommand{Method: "player_state", Params: CommandParams{"Balls": balls_pos}})
} 


func (ps *PlayerSession) updateDirection(command PlayerSessionCommand) {	
	xDif := math.Cos(command.Params["directionAngelRad"].(float64) ) * ps.getSpeed()
	yDif := math.Sin(command.Params["directionAngelRad"].(float64) )  * ps.getSpeed()
	log.Printf("update Direction : x : %f, y: %f  \n", xDif, yDif )
	ps.Direction.X = float32(xDif)
	ps.Direction.Y = float32(yDif)

}

func (ps *PlayerSession) getSpeed() float64 {
	return 5.0
}

func (ps *PlayerSession) moveBalls() []*Ball {
	for _, ball := range ps.Balls {
		ps.Game.Field.moveBall(ball, &ps.Direction)
	}
	return ps.Balls
}
