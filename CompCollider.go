package EasyEC

/*
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"*/
import "image"

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

func NewCompCollider(entity *Entity, data *Data) {
	new := CompCollider{}
	new.entity = entity
	Components.Collider = append(Components.Collider, &new)
}

func (s *CompCollider) collide(rectPlayer image.Rectangle) {

	rectEnemy := image.Rect(int(G.compRendNPOMap[s.entity.id].Xpos), int(G.compRendNPOMap[s.entity.id].Ypos), int(G.compRendNPOMap[s.entity.id].Xpos)+30, int(G.compRendNPOMap[s.entity.id].Ypos)+30)
	//rectPlayer := image.Rect(int(player.xpos), int(player.ypos), int(player.xpos) + 30, int(player.ypos) + 30)
	if rectEnemy.Overlaps(rectPlayer) {

		print("yo whats up")
		//for _,z := range s.events {
		//z.execute(s.entity)
		//}
	}
}
