package game

import "log"

type Field struct {
    Width float32
    Height float32
}

func NewField(width, height float32) *Field {
    field := Field{Width : width, Height: height}
    return &field
}


func (f *Field) moveBall(b *Ball, direction *Cordinate) {
    xNew := b.Position.X + direction.X
    yNew := b.Position.Y + direction.Y

    wL := b.Mass /2
    wR := f.Width + b.Mass/2

    hT := b.Mass / 2
    hB := f.Height + b.Mass /2

    if ( xNew <  wR && xNew > wL  ) {
        b.Position.X = xNew
    }

    if ( yNew < hB && yNew > hT ) {
        b.Position.Y = yNew
    }

    
    
    
    log.Printf(" ball moved %f, %f", b.Position.X, b.Position.Y)
}

func func_name() {
    
}