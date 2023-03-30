package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	collideables []collide
	player Sprite
	renderDistance float64
	colliders []CompCollider
	colliderMap map[*Enemy]CompCollider
	entities []*Entity
	compRendNPO []*CompRendNPO
	compRendNPOMap map[int]*CompRendNPO
	compCollider []*CompCollider
	compMoveRand []*CompMoveRand
	compMovePatrol []*CompMovePatrol
	entityId int
}

func NewGame() Game {
	new := Game{}
	new.renderDistance = 500
	return new
}

func (g *Game) Update() error {
	deltax, deltay := g.player.move()
	for _, s := range g.compRendNPO {
		s.moveCamera(deltax, deltay)
	}
	rectPlayer := image.Rect(int(g.player.xpos), int(g.player.ypos), int(g.player.xpos)+30, int(g.player.ypos)+30)
	/*for _, s := range g.collideables {
		s.collide(rectPlayer)
	}*/
	for _, s := range g.compMoveRand {
		s.MoveRand()
	}
	for _, s := range g.compCollider {
		s.collide(rectPlayer)
	}
	return nil
}

var i float64 = 0

func (g *Game) Draw(screen *ebiten.Image) {

	g.player.draw(screen)
	for _, s := range g.compRendNPO {
		x, y := s.getPosition()
		if (x-g.player.xpos) > -g.renderDistance && (x-g.player.xpos) < g.renderDistance && (y-g.player.ypos) > -g.renderDistance && (y-g.player.ypos) < g.renderDistance {
			s.draw(screen)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

