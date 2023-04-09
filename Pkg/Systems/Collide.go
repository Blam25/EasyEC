package Systems

import (
	"image"
	"time"

	C "github.com/Blam25/Test/Pkg/Components"
)

type Collide struct {
}

func NewCollide() {
	new := &Collide{}
	Systems.Second = append(Systems.Second, new)
}

func (s *Collide) Execute() {
	for _, z := range C.Components.Collider {

		//print("yo")
		s.collide(z)

	}
}
func (s *Collide) collide(comp *C.Collider) {

	rectEnemy := image.Rect(
		int(C.Components.PositionMap[comp.Entity.GetId()].Xpos),
		int(C.Components.PositionMap[comp.Entity.GetId()].Ypos),
		int(C.Components.PositionMap[comp.Entity.GetId()].Xpos)+30,
		int(C.Components.PositionMap[comp.Entity.GetId()].Ypos)+30,
	)
	rectPlayer := image.Rect(
		int(C.Components.PositionMap[C.Components.Player.Entity.GetId()].Xpos),
		int(C.Components.PositionMap[C.Components.Player.Entity.GetId()].Ypos),
		int(C.Components.PositionMap[C.Components.Player.Entity.GetId()].Xpos)+30,
		int(C.Components.PositionMap[C.Components.Player.Entity.GetId()].Ypos)+30,
	)
	if rectEnemy.Overlaps(rectPlayer) {
		if !comp.Collided {

			comp.Collided = true
			go timerTwoSec(comp)
			C.NewCollision(comp.Entity, C.Components.Player.Entity)
		}
	}
}

func timerTwoSec(comp *C.Collider) {
	time.Sleep(2 * time.Second)
	comp.Collided = false
}
