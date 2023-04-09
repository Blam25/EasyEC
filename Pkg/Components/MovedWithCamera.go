package Components

import E "github.com/Blam25/Test/Pkg/Entities"

type MovedWithCamera struct {
	Entity *E.Entity
}

func NewMoveWithCamera(entity *E.Entity) {
	new := MovedWithCamera{}
	new.Entity = entity
	Components.MovedWithCamera = append(Components.MovedWithCamera, &new)
	Components.MovedWithCameraMap[entity.GetId()] = &new
}
