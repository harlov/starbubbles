package core

import (
	"fmt"
	"log"
)

type App struct {
	Configuration *Configuration
}

func (app App) Run() {
	fmt.Println("core run...")
	app.Configuration = &Configuration{}
	err := app.Configuration.loadFromFile()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}
	app.runServer()
}
