package EasyEC

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type CompRend struct {
	Img    *ebiten.Image
	Op     ebiten.DrawImageOptions
	Xpos   float64
	Ypos   float64
	Entity *Entity
}

//type NewRenderedNonPlayer func(*Entity)

func NewCompRenderedNonPlayer(entity *Entity, xpos float64, ypos float64) {
	new := CompRend{}
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
func NewCompRend(entity *Entity, image string) {
	new := CompRend{}
	var err error
	new.Img, _, err = ebitenutil.NewImageFromFile(image)
	if err != nil {
		log.Fatal(err)
	}
	new.Entity = entity
	Components.RendNPO = append(Components.RendNPO, &new)
}

func (s *CompRend) getPosition() (float64, float64) {
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
func (s *CompRend) getEntity() *Entity {
	return s.Entity
}
