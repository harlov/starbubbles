package game

type Player struct {
    Name string
}


func NewPlayer(name string) *Player {
    var player = Player{Name:name}
    return &player
}