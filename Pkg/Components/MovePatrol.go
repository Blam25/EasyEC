package Components

import E "github.com/Blam25/Test/Pkg/Entities"

type MovePatrol struct {
	entity   *E.Entity
	Velocity float64
	goals    []float64
}

func NewMovePatrol(entity *E.Entity, goals []float64) {
	new := MovePatrol{}
	new.entity = entity
	new.goals = goals
	Components.MovePatrol = append(Components.MovePatrol, &new)
}

func (s *MovePatrol) MovePatrol() {

}
