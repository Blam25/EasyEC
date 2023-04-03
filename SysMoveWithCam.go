package main

import "github.com/hajimehoshi/ebiten/v2"

type SysMoveWithCam struct {
}

func NewSysMoveWithCam() {
	new := &SysMoveWithCam{}
	systems.first = append(systems.first, new)
}

func (s *SysMoveWithCam) execute() {
	deltax, deltay := s.move()
	for _, z := range components.RendNPO {
		z.ypos = z.ypos + deltay
		z.xpos = z.xpos + deltax
	}
}

func (s *SysMoveWithCam) move() (float64, float64) {
	var x float64
	var y float64
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		//print("hej")
		//s.ypos = s.ypos - s.moveSpeed
		y = components.player.moveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		//s.ypos = s.ypos + s.moveSpeed
		y = -components.player.moveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		//s.xpos = s.xpos + s.moveSpeed
		x = -components.player.moveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		//s.xpos = s.xpos - s.moveSpeed
		x = components.player.moveSpeed
	}
	return x, y
}
