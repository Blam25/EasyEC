package EasyEC

/*
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"*/
import (
	"math/rand"
)

type CompMoveRand struct {
	entity   *Entity
	Velocity float64
}

/*
func NewCompCollider(enemy *Enemy, events []event) CompCollider{
	new := CompCollider{}
	new.enemy = enemy
	new.events = events
	return new
}*/

func NewCompMoveRand(entity *Entity, data *Data) {
	new := CompMoveRand{}
	new.entity = entity
	G.compMoveRand = append(G.compMoveRand, &new)
}

func (s *CompMoveRand) MoveRand() {
	G.compRendNPOMap[s.entity.id].Xpos = G.compRendNPOMap[s.entity.id].Xpos + (float64(rand.Intn(10)) - 5)
	G.compRendNPOMap[s.entity.id].Ypos = G.compRendNPOMap[s.entity.id].Ypos + (float64(rand.Intn(10)) - 5)
}
