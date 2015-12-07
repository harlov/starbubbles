package game

import (
    "github.com/gorilla/websocket"
    "encoding/json"
    "log"
)

type PlayerSession struct {
    Socket *websocket.Conn
    Player *Player
    Game *Game
    Direction Cordinate
    Balls []*Ball
    Score float32
    
}

type PlayerSessionCommand struct {
    Method string
    Params map[string]interface{}
}


func NewSession(ws *websocket.Conn, player *Player, game *Game) *PlayerSession {
    ps := PlayerSession{Socket:ws, Player:player, Game:game}
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

        ps.Command(player_command)
    }
    ps.Socket.Close()
}


func (ps *PlayerSession) Command(command PlayerSessionCommand) {
    log.Printf("method : %s \n", command.Method)
}