package EasyEC

type ECComponents struct {
	Player             *CompPlayer
	PlayerMap          map[int]*CompPlayer
	RendNPO            []*CompRend
	RendNPOMap         map[int]*CompRend
	Collider           []*CompCollider
	ColliderMap        map[int]*CompCollider
	MoveRand           []*CompMoveRand
	MoveRandMap        map[int]*CompMoveRand
	MovePatrol         []*CompMovePatrol
	MovePatrolMap      map[int]*CompMovePatrol
	Collision          []*CompCollision
	CollisionMap       map[int]*CompCollision
	Health             []*CompHealth
	HealthMap          map[int]*CompHealth
	Position           []*CompPosition
	PositionMap        map[int]*CompPosition
	MovedWithCamera    []*CompMovedWithCamera
	MovedWithCameraMap map[int]*CompMovedWithCamera

	EventCollision    []InteractionEvent
	EventCollisionMap map[int][]InteractionEvent
}

type InteractionEvent interface {
	Trigger(*Entity)
}
