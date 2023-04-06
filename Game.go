package EasyEC

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	collideables   []collide
	player         CompPlayer
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

	for _, s := range Systems.First {
		s.Execute()
	}
	for _, s := range Systems.Second {
		s.Execute()
	}
	for _, s := range Systems.Third {
		s.Execute()
	}
	for _, s := range Systems.Fourth {
		s.Execute()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	Components.Player.draw(screen)
	for _, s := range Systems.DrawFirst {
		s.execute(screen)
	}
	for _, s := range Systems.DrawSecond {
		s.execute(screen)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}
