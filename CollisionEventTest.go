package EasyEC

type CollisionEventTest struct {
	entity *Entity
	damage int
}

func NewCollisionEventTest(entity *Entity, damage int) {
	new := CollisionEventTest{}
	new.damage = damage
	new.entity = entity
	Components.EventCollisionMap[entity.GetId()] = append(Components.EventCollisionMap[entity.GetId()], &new)
	//Components.EventCollision = append(Components.EventCollision, &new)
}

func (s *CollisionEventTest) ExecuteInteraction(player *Entity) {
	Components.HealthMap[player.GetId()].Health = Components.HealthMap[player.GetId()].Health - s.damage
	print(Components.HealthMap[player.GetId()].Health)
}
