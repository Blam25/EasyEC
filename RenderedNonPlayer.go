package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type CompRendNPO struct {
	img       *ebiten.Image
	op        ebiten.DrawImageOptions
	xpos      float64
	ypos      float64
	entity		*Entity
}

//type NewRenderedNonPlayer func(*Entity)  

func NewCompRenderedNonPlayer(entity *Entity, xpos float64, ypos float64) {
	new := CompRendNPO{}
	var err error
	new.img, _, err = ebitenutil.NewImageFromFile("assets/gopher.png")
	if err != nil {
		log.Fatal(err)
	}
	new.xpos = xpos
	new.ypos = ypos
	new.entity = entity
	g.compRendNPO = append(g.compRendNPO, &new)
}

//takes x, y (Float64) as position
func NewCompRendNPO(entity *Entity, data *dataArgs) {
	new := CompRendNPO{}
	new.xpos = data.floats[0]
	new.ypos = data.floats[1]
	var err error
	new.img, _, err = ebitenutil.NewImageFromFile("assets/gopher.png")
	if err != nil {
		log.Fatal(err)
	}
	new.entity = entity
	g.compRendNPO = append(g.compRendNPO, &new)
	g.compRendNPOMap[entity.id] = &new
}

func (s *CompRendNPO) moveCamera(deltaX float64, deltaY float64) {
	s.ypos = s.ypos + deltaY
	s.xpos = s.xpos + deltaX
}

func (s *CompRendNPO) getPosition() (float64, float64){
	return s.xpos, s.ypos
}

func (s *CompRendNPO) draw(screen *ebiten.Image) {
	s.op.GeoM.Reset()
	s.op.GeoM.Translate(s.xpos, s.ypos)
	screen.DrawImage(s.img, &s.op)
}