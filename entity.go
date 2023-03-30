package main

type Entity struct {
	id int
	comprend CompRendNPO
}

func NewEntity() *Entity{
	new := Entity{}
	new.id = g.entityId
	g.entityId++
	g.entities = append(g.entities, &new)
	return &new
}

func (s *Entity) with(f func(*Entity, *dataArgs), data *dataArgs) *Entity {
	f(s, data)
	return s
	 
}