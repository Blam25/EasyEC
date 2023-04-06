package EasyEC

type CompCollision struct {
	Entity *Entity
	Player *Entity
	Ypos   int
}

func NewCompCollision(entity *Entity, player *Entity) {
	new := CompCollision{}
	new.Entity = entity
	new.Player = player
	Components.Collision = append(Components.Collision, &new)
	Components.CollisionMap[entity.GetId()] = &new
}
