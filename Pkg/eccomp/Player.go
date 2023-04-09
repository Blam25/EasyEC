package eccomp

import (
	E "github.com/Blam25/Test/Pkg/ecentity"

	_ "image/png"
)

type Player struct {
	Entity    *E.Entity
	MoveSpeed float64
}

func NewPlayer(entity *E.Entity) Player {
	s := Player{}
	s.Entity = entity
	s.MoveSpeed = 7
	Components.Player = &s
	return s
}
