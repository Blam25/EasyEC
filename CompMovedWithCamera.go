package EasyEC

type CompMovedWithCamera struct {
	Entity *Entity
}

func NewCompMoveWithCamera(entity *Entity) {
	new := CompMovedWithCamera{}
	new.Entity = entity
	Components.MovedWithCamera = append(Components.MovedWithCamera, &new)
	Components.MovedWithCameraMap[entity.GetId()] = &new
}
