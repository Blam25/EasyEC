package EasyEC

type SysClear struct {
}

func NewSysClear() {
	new := &SysClear{}
	Systems.Third = append(Systems.Third, new)
}

func (s *SysClear) Execute() {
	Components.Collision = nil
	Components.CollisionMap = make(map[int]*CompCollision)
}
