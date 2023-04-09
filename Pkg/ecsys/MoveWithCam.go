package ecsys

import (
	"github.com/hajimehoshi/ebiten/v2"

	C "github.com/Blam25/Test/Pkg/eccomp"
)

type MoveWithCam struct {
}

func NewMoveWithCam() {
	new := &MoveWithCam{}
	Systems.First = append(Systems.First, new)
}

func (s *MoveWithCam) Execute() {
	deltax, deltay := s.move()
	for _, z := range C.Components.MovedWithCamera {
		C.Components.PositionMap[z.Entity.GetId()].Ypos = C.Components.PositionMap[z.Entity.GetId()].Ypos + deltay
		C.Components.PositionMap[z.Entity.GetId()].Xpos = C.Components.PositionMap[z.Entity.GetId()].Xpos + deltax
	}
}

func (s *MoveWithCam) move() (float64, float64) {
	var x float64
	var y float64
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		//print("hej")
		//s.ypos = s.ypos - s.moveSpeed
		y = C.Components.Player.MoveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		//s.ypos = s.ypos + s.moveSpeed
		y = -C.Components.Player.MoveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		//s.xpos = s.xpos + s.moveSpeed
		x = -C.Components.Player.MoveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		//s.xpos = s.xpos - s.moveSpeed
		x = C.Components.Player.MoveSpeed
	}
	return x, y
}
