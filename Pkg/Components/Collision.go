package Components

import E "github.com/Blam25/Test/Pkg/Entities"

type Collision struct {
	Entity *E.Entity
	Player *E.Entity
	Ypos   int
}

func NewCollision(entity *E.Entity, player *E.Entity) {
	new := Collision{}
	new.Entity = entity
	new.Player = player
	Components.Collision = append(Components.Collision, &new)
	Components.CollisionMap[entity.GetId()] = &new
}
