package core

import (
	"fmt"
	"log"
	"github.com/harlov/starbubbles/game"
)

type App struct {
	Configuration *Configuration
	Games []*game.Game
}

func (app App) Run() {
	fmt.Println("core run...")
	app.Configuration = &Configuration{}
	err := app.Configuration.loadFromFile()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}


	app.Games = append(app.Games, game.NewGame("main_game") )

	app.runServer()
}
