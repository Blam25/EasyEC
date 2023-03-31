package main

/*
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"*/
import	"image"

type CompCollider struct {
	entity *Entity
	events []event
	Velocity float64
}
/*
func NewCompCollider(enemy *Enemy, events []event) CompCollider{
	new := CompCollider{}
	new.enemy = enemy
	new.events = events
	return new
}*/

func NewCompCollider(entity *Entity, data *data) {
	new := CompCollider{}
	new.entity = entity
	g.compCollider = append(g.compCollider, &new)
}

func (s *CompCollider) collide(rectPlayer image.Rectangle) {
	
	rectEnemy := image.Rect(int(g.compRendNPOMap[s.entity.id].xpos), int(g.compRendNPOMap[s.entity.id].ypos), int(g.compRendNPOMap[s.entity.id].xpos) + 30, int(g.compRendNPOMap[s.entity.id].ypos) + 30)
	//rectPlayer := image.Rect(int(player.xpos), int(player.ypos), int(player.xpos) + 30, int(player.ypos) + 30)
	if rectEnemy.Overlaps(rectPlayer) {
		
		print("yo whats up")
		//for _,z := range s.events {
			//z.execute(s.entity)
		//}
	}
}


