package Components

import E "github.com/Blam25/Test/Pkg/Entities"

/*
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"*/

type Collider struct {
	Entity     *E.Entity
	Collisions []*E.Entity
	Collided   bool
}

func NewCollider(entity *E.Entity) {
	new := Collider{}
	new.Entity = entity
	Components.Collider = append(Components.Collider, &new)
}
