package core

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func gameWebsocketSocHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(w, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		return
	}

	ws.WriteMessage(websocket.TextMessage, []byte("hello, gamer!"))
	ws.WriteMessage(websocket.TextMessage, []byte("by, gamer!"))
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path[1:])
	http.ServeFile(w, r, "static/dist/"+r.URL.Path[1:])
}

func (app App) runServer() {
	http.HandleFunc("/", staticHandler)
	http.HandleFunc("/game_ws", gameWebsocketSocHandler)
	log.Println("run server on port ", (*app.Configuration).Server.Port)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", (*app.Configuration).Server.Host, (*app.Configuration).Server.Port), nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
