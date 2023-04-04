package EasyEC

/*
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"*/
import	(
	"math/rand")

type CompMoveRand struct {
	entity *Entity
	Velocity float64
}
/*
func NewCompCollider(enemy *Enemy, events []event) CompCollider{
	new := CompCollider{}
	new.enemy = enemy
	new.events = events
	return new
}*/

func NewCompMoveRand(entity *Entity, data *data) {
	new := CompMoveRand{}
	new.entity = entity
	g.compMoveRand = append(g.compMoveRand, &new)
}

func (s *CompMoveRand) MoveRand() {
	g.compRendNPOMap[s.entity.id].xpos = g.compRendNPOMap[s.entity.id].xpos + (float64(rand.Intn(10))-5)
	g.compRendNPOMap[s.entity.id].ypos = g.compRendNPOMap[s.entity.id].ypos + (float64(rand.Intn(10))-5)
}