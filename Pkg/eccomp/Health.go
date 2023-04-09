package eccomp

import E "github.com/Blam25/Test/Pkg/ecentity"

type Health struct {
	entity *E.Entity
	Health int
}

func NewHealth(entity *E.Entity, health int) {
	new := Health{}
	new.entity = entity
	new.Health = health
	Components.Health = append(Components.Health, &new)
	Components.HealthMap[entity.GetId()] = &new
}
