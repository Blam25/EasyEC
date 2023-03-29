package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type CompRenderedNonPlayer struct {
	img       *ebiten.Image
	op        ebiten.DrawImageOptions
	xpos      float64
	ypos      float64
	entity		*Entity
}

//type NewRenderedNonPlayer func(*Entity)  

func NewCompRenderedNonPlayer(entity *Entity, xpos float64, ypos float64) {
	new := CompRenderedNonPlayer{}
	var err error
	new.img, _, err = ebitenutil.NewImageFromFile("assets/gopher.png")
	if err != nil {
		log.Fatal(err)
	}
	new.xpos = xpos
	new.ypos = ypos
	new.entity = entity
	g.compRenderedNonPlayers = append(g.compRenderedNonPlayers, &new)
}

func NewCompRenderedNonPlayer2(entity *Entity, data *dataArgs) {
	new := CompRenderedNonPlayer{}
	var err error
	new.img, _, err = ebitenutil.NewImageFromFile("assets/gopher.png")
	if err != nil {
		log.Fatal(err)
	}
	new.xpos = data.floats[0]
	new.ypos = data.floats[1]
	new.entity = entity
	g.compRenderedNonPlayers = append(g.compRenderedNonPlayers, &new)
}

func (s *CompRenderedNonPlayer) moveCamera(deltaX float64, deltaY float64) {
	s.ypos = s.ypos + deltaY
	s.xpos = s.xpos + deltaX
}

func (s *CompRenderedNonPlayer) getPosition() (float64, float64){
	return s.xpos, s.ypos
}

func (s *CompRenderedNonPlayer) draw(screen *ebiten.Image) {
	s.op.GeoM.Reset()
	s.op.GeoM.Translate(s.xpos, s.ypos)
	screen.DrawImage(s.img, &s.op)
}