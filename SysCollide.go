package EasyEC

import "image"

type SysCollide struct {
}

func NewSysCollide() {
	new := &SysCollide{}
	Systems.Second = append(Systems.Second, new)
}

func (s *SysCollide) Execute() {
	for _, z := range Components.Collider {
		z.Collided = false
		s.collide(z)
	}
}
func (s *SysCollide) collide(comp *CompCollider) {

	rectEnemy := image.Rect(int(G.compRendNPOMap[comp.entity.id].Xpos), int(G.compRendNPOMap[comp.entity.id].Ypos), int(G.compRendNPOMap[comp.entity.id].Xpos)+30, int(G.compRendNPOMap[comp.entity.id].Ypos)+30)
	rectPlayer := image.Rect(int(Components.Player.Xpos), int(Components.Player.Ypos), int(Components.Player.Xpos)+30, int(Components.Player.Ypos)+30)
	if rectEnemy.Overlaps(rectPlayer) {

		comp.Collided = true
		NewCompCollision(comp.entity, Components.Player.entity)

		//print("yo whats up")
		//for _,z := range s.events {
		//z.execute(s.entity)
		//}
	}
}
