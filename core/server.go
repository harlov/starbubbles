package core

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"net/url"
	"github.com/harlov/starbubbles/game"
)

type serverHandler struct {
	app *App
	Handler func(*App, http.ResponseWriter, *http.Request)
}


func (sh serverHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sh.Handler(sh.app, w, r)
}

func gameWebsocketSocHandler(app *App, w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(w, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		return
	}

	params, _ := url.ParseQuery(r.URL.RawQuery)
	player := game.NewPlayer(params["name"][0])
	
	select_game := (*app).Games[0]
	session := game.NewSession(ws, player, select_game)
	select_game.JoinPlayer(session)

	log.Println("player joined to game")

	var game_name = (*select_game).Name

	

	ws.WriteMessage(websocket.TextMessage, []byte("hello, gamer! wellcome to game"))
	ws.WriteMessage(websocket.TextMessage, []byte(game_name))
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/dist/"+r.URL.Path[1:])
}

func (app App) runServer() {
	http.HandleFunc("/", staticHandler)
	http.Handle("/game_ws", serverHandler{&app, gameWebsocketSocHandler})
	log.Println("run server on port ", (*app.Configuration).Server.Port)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", (*app.Configuration).Server.Host, (*app.Configuration).Server.Port), nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
