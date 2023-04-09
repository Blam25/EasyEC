package Events

import (
	E "github.com/Blam25/Test/Pkg/Entities"
)

var Events *ECEvents = &ECEvents{}

func init() {
	Events.CollisionMap = make(map[int][]InteractionEvent)
}

type InteractionEvent interface {
	Trigger(*E.Entity)
}

type ECEvents struct {
	Collision    []InteractionEvent
	CollisionMap map[int][]InteractionEvent
}
