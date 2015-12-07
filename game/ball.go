package game

import "log"

type Ball struct {
    Position Cordinate
    Mass float32
}


func NewBall(init_mass float32, ix float32, iy float32) *Ball {
    var ball = Ball{
            Mass: init_mass,
            Position: Cordinate{
                x: ix,
                y: iy,
            },
        }
    return &ball
}


func (b *Ball) move(direction Cordinate) {
    log.Println("   ball moved")
}