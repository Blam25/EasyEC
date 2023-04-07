package EasyEC

type SysClear struct {
}

func NewSysClear() {
	new := &SysClear{}
	Systems.Fourth = append(Systems.Fourth, new)
}

func (s *SysClear) Execute() {
	Components.Collision = nil
	Components.CollisionMap = make(map[int]*CompCollision)
}
