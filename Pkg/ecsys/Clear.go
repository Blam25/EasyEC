package ecsys

import (
	C "github.com/Blam25/Test/Pkg/eccomp"
)

type Clear struct {
}

func NewClear() {
	new := &Clear{}
	Systems.Fourth = append(Systems.Fourth, new)
}

func (s *Clear) Execute() {
	C.Components.Collision = nil
	C.Components.CollisionMap = make(map[int]*C.Collision)
}
