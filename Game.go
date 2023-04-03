package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	collideables   []collide
	player         Sprite
	renderDistance float64
	colliders      []CompCollider
	colliderMap    map[*Enemy]CompCollider
	entities       []*Entity
	compRendNPO    []*CompRendNPO
	compRendNPOMap map[int]*CompRendNPO
	compCollider   []*CompCollider
	compMoveRand   []*CompMoveRand
	compMovePatrol []*CompMovePatrol
	entityId       int
}

func NewGame() Game {
	new := Game{}
	new.renderDistance = 500
	return new
}

func (g *Game) Update() error {

	for _, s := range systems.first {
		s.execute()
	}
	for _, s := range systems.second {
		s.execute()
	}
	for _, s := range systems.third {
		s.execute()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	components.player.draw(screen)
	for _, s := range systems.drawFirst {
		s.execute(screen)
	}
	for _, s := range systems.drawSecond {
		s.execute(screen)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}
