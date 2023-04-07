package EasyEC

import (
	"image"
	"time"
)

type SysCollide struct {
}

func NewSysCollide() {
	new := &SysCollide{}
	Systems.Second = append(Systems.Second, new)
}

func (s *SysCollide) Execute() {
	for _, z := range Components.Collider {

		//print("yo")
		s.collide(z)

	}
}
func (s *SysCollide) collide(comp *CompCollider) {

	rectEnemy := image.Rect(
		int(Components.PositionMap[comp.entity.id].Xpos),
		int(Components.PositionMap[comp.entity.id].Ypos),
		int(Components.PositionMap[comp.entity.id].Xpos)+30,
		int(Components.PositionMap[comp.entity.id].Ypos)+30,
	)
	rectPlayer := image.Rect(
		int(Components.PositionMap[Components.Player.entity.GetId()].Xpos),
		int(Components.PositionMap[Components.Player.entity.GetId()].Ypos),
		int(Components.PositionMap[Components.Player.entity.GetId()].Xpos)+30,
		int(Components.PositionMap[Components.Player.entity.GetId()].Ypos)+30,
	)
	if rectEnemy.Overlaps(rectPlayer) {
		if comp.Collided == false {

			comp.Collided = true
			go timerTwoSec(comp)

			//comp.Collided = true
			NewCompCollision(comp.entity, Components.Player.entity)
		}

		//print("yo whats up")
		//for _,z := range s.events {
		//z.execute(s.entity)
		//}
	}
}

func timerTwoSec(comp *CompCollider) {
	time.Sleep(2 * time.Second)
	comp.Collided = false
}
