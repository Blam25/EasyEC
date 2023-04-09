package Components

import E "github.com/Blam25/Test/Pkg/Entities"

type Position struct {
	entity *E.Entity
	Xpos   float64
	Ypos   float64
}

func NewPosition(entity *E.Entity, xpos float64, ypos float64) {
	new := Position{}
	new.entity = entity
	new.Xpos = xpos
	new.Ypos = ypos
	Components.Position = append(Components.Position, &new)
	Components.PositionMap[entity.GetId()] = &new
}
