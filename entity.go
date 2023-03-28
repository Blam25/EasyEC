package main

type Entity struct {
}

func NewEntity() *Entity{
	new := Entity{}
	g.entities = append(g.entities, &new)
	return &new
}

/*func (s *Entity) with(f func(*Entity), entity *Entity) *Entity {
	f(entity)
	return entity
	 
}*/