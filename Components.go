package EasyEC

type Components struct {
	player        Sprite
	RendNPO       []*CompRendNPO
	RendNPOMap    map[int]*CompRendNPO
	Collider      []*CompCollider
	ColliderMap   map[int]CompCollider
	MoveRand      []*CompMoveRand
	MoveRandMap   map[int]CompMoveRand
	MovePatrol    []*CompMovePatrol
	MovePatrolMap map[int]CompMovePatrol
}
