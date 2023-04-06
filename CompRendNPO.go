package EasyEC

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type CompRendNPO struct {
	Img    *ebiten.Image
	Op     ebiten.DrawImageOptions
	Xpos   float64
	Ypos   float64
	Entity *Entity
}

//type NewRenderedNonPlayer func(*Entity)

func NewCompRenderedNonPlayer(entity *Entity, xpos float64, ypos float64) {
	new := CompRendNPO{}
	var err error
	new.Img, _, err = ebitenutil.NewImageFromFile("assets/gopher.png")
	if err != nil {
		log.Fatal(err)
	}
	new.Xpos = xpos
	new.Ypos = ypos
	new.Entity = entity
	//systemsRenderMap["SystemRender"].getArray() = append(systemsRenderMap["SystemRender"].getArray(), &new)
	print("append2")
	//g.compRendNPO = append(g.compRendNPO, &new)
}

// takes x, y (Float64) as position
func NewCompRendNPO(entity *Entity, data *Data) {
	new := CompRendNPO{}
	new.Xpos = data.Floats[0]
	new.Ypos = data.Floats[1]
	var err error
	new.Img, _, err = ebitenutil.NewImageFromFile("assets/gopher.png")
	if err != nil {
		log.Fatal(err)
	}
	new.Entity = entity
	G.compRendNPO = append(G.compRendNPO, &new)
	Components.RendNPO = append(Components.RendNPO, &new)
	//systemsRenderMap["s1"].appendArray(&new)
	G.compRendNPOMap[entity.id] = &new
}
func (s *CompRendNPO) getPosition() (float64, float64) {
	return s.Xpos, s.Ypos
}

/*
	func (s *CompRendNPO) moveCamera(deltaX float64, deltaY float64) {
		s.ypos = s.ypos + deltaY
		s.xpos = s.xpos + deltaX
	}

	func (s *CompRendNPO) draw(screen *ebiten.Image) {
		s.op.GeoM.Reset()
		s.op.GeoM.Translate(s.xpos, s.ypos)
		screen.DrawImage(s.img, &s.op)
	}
*/
func (s *CompRendNPO) getEntity() *Entity {
	return s.Entity
}
