package core

import (
	"fmt"
)

type Server struct {
}

func (s Server) run() {
	fmt.Println("run server ")
	fmt.Println()
}
