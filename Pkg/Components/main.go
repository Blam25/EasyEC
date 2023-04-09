package Components

var Components *ECComponents = &ECComponents{}

func init() {
	Components.CollisionMap = make(map[int]*Collision)
	Components.HealthMap = make(map[int]*Health)
	Components.PositionMap = make(map[int]*Position)
	Components.MovedWithCameraMap = make(map[int]*MovedWithCamera)
}

type ECComponents struct {
	Player             *Player
	PlayerMap          map[int]*Player
	RendNPO            []*Rend
	RendNPOMap         map[int]*Rend
	Collider           []*Collider
	ColliderMap        map[int]*Collider
	MoveRand           []*MoveRand
	MoveRandMap        map[int]*MoveRand
	MovePatrol         []*MovePatrol
	MovePatrolMap      map[int]*MovePatrol
	Collision          []*Collision
	CollisionMap       map[int]*Collision
	Health             []*Health
	HealthMap          map[int]*Health
	Position           []*Position
	PositionMap        map[int]*Position
	MovedWithCamera    []*MovedWithCamera
	MovedWithCameraMap map[int]*MovedWithCamera
}

type Component interface {
}
