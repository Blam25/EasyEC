package ecentity

type Entity struct {
	id int
}

func NewEntity() *Entity {

	new := Entity{}
	new.id = entityId
	entityId++
	entities = append(entities, &new)
	return &new
}

func (s *Entity) GetId() int {
	return s.id
}
