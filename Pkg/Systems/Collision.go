package Systems

import (
	C "github.com/Blam25/Test/Pkg/Components"
	EV "github.com/Blam25/Test/Pkg/Events"
)

type Collision struct {
}

func NewCollision() {
	new := &Collision{}
	Systems.Third = append(Systems.Third, new)
}

func (s *Collision) Execute() {
	for _, z := range C.Components.Collision {
		for _, j := range EV.Events.CollisionMap[z.Entity.GetId()] {
			j.Trigger(z.Player)
		}
		//C.Components.EventCollisionMap[z.Entity.GetId()].ExecuteInteraction(z.Entity, z.Player)
	}
}
