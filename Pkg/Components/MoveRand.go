package Components

import (
	E "github.com/Blam25/Test/Pkg/Entities"
)

type MoveRand struct {
	entity   *E.Entity
	Velocity float64
}

func NewMoveRand(entity *E.Entity) {
	new := MoveRand{}
	new.entity = entity
	Components.MoveRand = append(Components.MoveRand, &new)
}
