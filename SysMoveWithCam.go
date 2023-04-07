package EasyEC

import "github.com/hajimehoshi/ebiten/v2"

type SysMoveWithCam struct {
}

func NewSysMoveWithCam() {
	new := &SysMoveWithCam{}
	Systems.First = append(Systems.First, new)
}

func (s *SysMoveWithCam) Execute() {
	deltax, deltay := s.move()
	for _, z := range Components.MovedWithCamera {
		Components.PositionMap[z.Entity.GetId()].Ypos = Components.PositionMap[z.Entity.GetId()].Ypos + deltay
		Components.PositionMap[z.Entity.GetId()].Xpos = Components.PositionMap[z.Entity.GetId()].Xpos + deltax
	}
}

func (s *SysMoveWithCam) move() (float64, float64) {
	var x float64
	var y float64
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		//print("hej")
		//s.ypos = s.ypos - s.moveSpeed
		y = Components.Player.MoveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		//s.ypos = s.ypos + s.moveSpeed
		y = -Components.Player.MoveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		//s.xpos = s.xpos + s.moveSpeed
		x = -Components.Player.MoveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		//s.xpos = s.xpos - s.moveSpeed
		x = Components.Player.MoveSpeed
	}
	return x, y
}
