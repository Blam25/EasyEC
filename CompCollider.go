package EasyEC

/*
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"*/

type CompCollider struct {
	entity     *Entity
	Collisions []*Entity
	Collided   bool
}

/*
func NewCompCollider(enemy *Enemy, events []event) CompCollider{
	new := CompCollider{}
	new.enemy = enemy
	new.events = events
	return new
}*/

func NewCompCollider(entity *Entity) {
	new := CompCollider{}
	new.entity = entity
	Components.Collider = append(Components.Collider, &new)
}
