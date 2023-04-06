package EasyEC

type SysCollision struct {
}

func NewSysCollision() {
	new := &SysCollision{}
	Systems.Third = append(Systems.Third, new)
}

func (s *SysCollision) Execute() {
	for _, z := range Components.Collision {
		for _, j := range Components.EventCollisionMap[z.Entity.GetId()] {
			j.ExecuteInteraction(z.Player)
		}
		//Components.EventCollisionMap[z.Entity.GetId()].ExecuteInteraction(z.Entity, z.Player)
	}
}
