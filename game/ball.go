package game

import "log"

type Ball struct {
	Position Cordinate
	Mass     float32
}

func NewBall(init_mass float32, ix float32, iy float32) *Ball {
	var ball = Ball{
		Mass: init_mass,
		Position: Cordinate{
			X: ix,
			Y: iy,
		},
	}
	return &ball
}

func (b *Ball) move(direction Cordinate) {
	b.Position.X += 1
	log.Println("   ball moved")
}
