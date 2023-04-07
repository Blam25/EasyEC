package EasyEC

type CompHealth struct {
	entity *Entity
	Health int
}

func NewCompHealth(entity *Entity, health int) {
	new := CompHealth{}
	new.entity = entity
	new.Health = health
	Components.Health = append(Components.Health, &new)
	Components.HealthMap[entity.GetId()] = &new
}
