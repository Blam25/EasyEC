package EasyEC

type EventDmgPlayer struct {
	entity *Entity
	damage int
}

func NewEventDmgPlayer(entity *Entity, damage int, eventMap map[int][]InteractionEvent) {
	new := EventDmgPlayer{}
	new.damage = damage
	new.entity = entity
	eventMap[entity.GetId()] = append(eventMap[entity.GetId()], &new)
	//Components.EventCollision = append(Components.EventCollision, &new)
}

func (s *EventDmgPlayer) Trigger(player *Entity) {
	Components.HealthMap[player.GetId()].Health = Components.HealthMap[player.GetId()].Health - s.damage
	print(Components.HealthMap[player.GetId()].Health)
}
