package eccomp

import E "github.com/Blam25/Test/Pkg/ecentity"

type MovePatrol struct {
	Entity           *E.Entity
	Velocity         float64
	TimePerDirection []int
	Moved            bool
	Dir              Direction
}

func NewMovePatrol(entity *E.Entity) {
	new := MovePatrol{}
	new.Entity = entity
	//new.TimePerDirection = timePerDirection
	Components.MovePatrol = append(Components.MovePatrol, &new)
}

func (s *MovePatrol) MovePatrol() {

}

type Direction int

const (
	Left Direction = iota
	Up
	Right
	Down
)
