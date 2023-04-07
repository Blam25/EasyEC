package EasyEC

type CompPosition struct {
	entity *Entity
	Xpos   float64
	Ypos   float64
}

func NewCompPosition(entity *Entity, xpos float64, ypos float64) {
	new := CompPosition{}
	new.entity = entity
	new.Xpos = xpos
	new.Ypos = ypos
	Components.Position = append(Components.Position, &new)
	Components.PositionMap[entity.GetId()] = &new
}
