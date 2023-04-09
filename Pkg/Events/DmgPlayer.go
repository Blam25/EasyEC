package Events

import (
	C "github.com/Blam25/Test/Pkg/Components"
	E "github.com/Blam25/Test/Pkg/Entities"
)

type DmgPlayer struct {
	entity *E.Entity
	damage int
}

func NewDmgPlayer(entity *E.Entity, damage int, eventMap map[int][]InteractionEvent) {
	new := DmgPlayer{}
	new.damage = damage
	new.entity = entity
	eventMap[entity.GetId()] = append(eventMap[entity.GetId()], &new)
	//C.Components.EventCollision = append(C.Components.EventCollision, &new)
}

func (s *DmgPlayer) Trigger(player *E.Entity) {
	C.Components.HealthMap[player.GetId()].Health = C.Components.HealthMap[player.GetId()].Health - s.damage
	print(C.Components.HealthMap[player.GetId()].Health)
}
