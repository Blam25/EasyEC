package EasyEC

type Entity struct {
	id         int
	components map[string]Component
}

func NewEntity() *Entity {

	new := Entity{}
	new.id = G.entityId
	G.entityId++
	G.entities = append(G.entities, &new)
	return &new
}

func (s *Entity) With(f func(*Entity, *Data), data *Data) *Entity {
	f(s, data)
	return s
}

func (s *Entity) GetId() int {
	return s.id
}

/*func toDataArg[data string | int | float64 | bool](input1 []data, input2 []data, input3 []data, input4 []data) *dataArgs {

	return input1
}
*/
