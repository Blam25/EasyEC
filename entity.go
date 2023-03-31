package main

type Entity struct {
	id         int
	components map[string]Component
}

func NewEntity() *Entity {

	new := Entity{}
	new.id = g.entityId
	g.entityId++
	g.entities = append(g.entities, &new)
	return &new
}

func (s *Entity) with(f func(*Entity, *data), data *data) *Entity {
	f(s, data)
	return s
}

/*func toDataArg[data string | int | float64 | bool](input1 []data, input2 []data, input3 []data, input4 []data) *dataArgs {

	return input1
}
*/
