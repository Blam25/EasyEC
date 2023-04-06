package EasyEC

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type CompPlayer struct {
	entity    *Entity
	Img       *ebiten.Image
	Op        ebiten.DrawImageOptions
	Xpos      float64
	Ypos      float64
	MoveSpeed float64
}

func NewCompPlayer(entity *Entity) CompPlayer {
	s := CompPlayer{}
	s.entity = entity
	var err error
	s.Img, _, err = ebitenutil.NewImageFromFile("assets/gopher.png")
	if err != nil {
		log.Fatal(err)
	}
	s.Xpos = 150
	s.Ypos = 150
	s.MoveSpeed = 7
	Components.Player = &s
	return s
}

func (s *CompPlayer) draw(screen *ebiten.Image) {
	s.Op.GeoM.Reset()
	s.Op.GeoM.Translate(s.Xpos, s.Ypos)
	screen.DrawImage(s.Img, &s.Op)
}

func (s *CompPlayer) move() (float64, float64) {
	var x float64
	var y float64
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		//s.ypos = s.ypos - s.moveSpeed
		y = s.MoveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		//s.ypos = s.ypos + s.moveSpeed
		y = -s.MoveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		//s.xpos = s.xpos + s.moveSpeed
		x = -s.MoveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		//s.xpos = s.xpos - s.moveSpeed
		x = s.MoveSpeed
	}
	return x, y
}
