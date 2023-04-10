package ecsys

import (
	"time"

	C "github.com/Blam25/Test/Pkg/eccomp"
)

type MovePatrol struct {
}

func NewMovePatrol() {
	new := &MovePatrol{}
	Systems.Second = append(Systems.Second, new)
}

func (s *MovePatrol) Execute() {
	for _, z := range C.Components.MovePatrol {

		if !z.Moved {
			switch z.Dir {
			case C.Left:
				z.Dir = C.Up
			case C.Up:
				z.Dir = C.Right

			case C.Right:
				z.Dir = C.Down

			case C.Down:
				z.Dir = C.Left
			}
			z.Moved = true
			go timerTwoSec2(z)

			//C.NewCollision(comp.Entity, C.Components.Player.Entity)
		}
		switch z.Dir {
		case C.Left:
			C.Components.PositionMap[z.Entity.GetId()].Xpos -= 5

		case C.Up:
			C.Components.PositionMap[z.Entity.GetId()].Ypos -= 5

		case C.Right:
			C.Components.PositionMap[z.Entity.GetId()].Xpos += 5

		case C.Down:
			C.Components.PositionMap[z.Entity.GetId()].Ypos += 5
		}

	}
}

func timerTwoSec2(comp *C.MovePatrol) {
	time.Sleep(2 * time.Second)
	comp.Moved = false
}
