package core

import (
	"fmt"
)

type App struct {
	Configuration Configuration
	Server        Server
}

func (app App) Run() {
	fmt.Println("core run...")
	app.Configuration.loadFromFile()
	app.Server.run()
}
