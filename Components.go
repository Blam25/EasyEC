package EasyEC

type ECComponents struct {
	Player            *CompPlayer
	PlayerMap         map[int]*CompPlayer
	RendNPO           []*CompRendNPO
	RendNPOMap        map[int]*CompRendNPO
	Collider          []*CompCollider
	ColliderMap       map[int]*CompCollider
	MoveRand          []*CompMoveRand
	MoveRandMap       map[int]*CompMoveRand
	MovePatrol        []*CompMovePatrol
	MovePatrolMap     map[int]*CompMovePatrol
	Collision         []*CompCollision
	CollisionMap      map[int]*CompCollision
	Health            []*CompHealth
	HealthMap         map[int]*CompHealth
	EventCollision    []InteractionEvent
	EventCollisionMap map[int][]InteractionEvent
}

type InteractionEvent interface {
	ExecuteInteraction(*Entity)
}
