package eccomp

import (
	E "github.com/Blam25/Test/Pkg/ecentity"

	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Rend struct {
	Img    *ebiten.Image
	Op     ebiten.DrawImageOptions
	Xpos   float64
	Ypos   float64
	Entity *E.Entity
}

func NewRend(entity *E.Entity, image string) {
	new := Rend{}
	var err error
	new.Img, _, err = ebitenutil.NewImageFromFile(image)
	if err != nil {
		log.Fatal(err)
	}
	new.Entity = entity
	Components.RendNPO = append(Components.RendNPO, &new)
	print("Hje")
}

func (s *Rend) getEntity() *E.Entity {
	return s.Entity
}
