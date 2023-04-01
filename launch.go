package main

import (
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var g Game
var entityOrder int = 1
var test int = 1.0
var systems *Systems
var systemsMap map[string]System
var systemsRenderMap map[string]SystemDraw
var components *Components

//var entities []*Entity
//var renderedNonPlayers []*renderedNonPlayer

//var op *ebiten.DrawImageOptions

func init() {
	components = &Components{}
	systems = &Systems{}
	systemsMap = make(map[string]System)
	systemsRenderMap = make(map[string]SystemDraw)
	NewSystemRender()

	//toDataArg([]string{"hej"})

	g = NewGame()
	g.compRendNPOMap = make(map[int]*CompRendNPO)
	sprite1 := New_Sprite()

	entity1 := NewEntity()
	print(&entity1)
	entity1.
		with(NewCompRendNPO, NewDataArg(nil, nil, FArr(500, 500), nil)).
		with(NewCompCollider, &data{strings: SArr("hej", "yo"), ints: IArr(9, 2)})

	entity2 := NewEntity().with(
		NewCompRendNPO, NewDataArgFloat(100, 100)).with(
		NewCompCollider, NewDataArgEmpty()).with(
		NewCompMoveRand, &data{})

	entity4 := NewEntity().with(
		NewCompRendNPO, NewDataArgFloat(400, 400)).with(
		NewCompMoveRand, NewDataArgEmpty()).with(
		NewCompCollider, NewDataArgEmpty())
	print(&entity4)

	NewEntity().with(
		NewCompRendNPO, NewDataArgFloat(400, 400)).with(
		NewCompCollider, NewDataArgEmpty())
	//print(&entity5)

	entity6 := NewEntity().
		with(NewCompRendNPO, NewDataArgFloat(400, 400)).with(
		NewCompCollider, NewDataArgEmpty())

	print(entity6)

	//		with(NewCompRendNPO, NewDataArgFloat(1000, 1000)).
	//						with(NewCompCollider, &dataArgs{})
	print(&entity2)
	//entity2.(NewCompRendNPO, NewDataArgFloat(1000, 1000)).
	//with(NewCompCollider, &dataArgs{})

	entity3 := NewEntity()
	entity3.with(
		NewCompRendNPO, NewDataArgFloat(1000, 0)).with(
		NewCompCollider, &data{})

	//NewCompRenderedNonPlayer(entity2, 1000, 1000)
	g.player = sprite1
	/*
				//enemy1 := New_Enemy()
				enemy2 := New_Enemy() ; enemy2.xpos = 500 ; enemy2.ypos = 500
				enemy2.addCollider([]event{NewDmgEvent(60, "Im hit"), NewDmgEvent(60, "Yes really")})
				//NewCollider(&enemy2, []event{NewDmgEvent(60, "Im hit"), NewDmgEvent(60, "Yes really")})
				enemy3 := New_Enemy() ; enemy3.xpos = 1000 ; enemy3.ypos = 1000 ;
				//game.collideables = append(game.collideables, &enemy1, &enemy2, &enemy3)

		//op := &ebiten.DrawImageOptions{}
						//op = &ebiten.DrawImageOptions{}*/
}

func initComponents() {
	components = &Components{}
}

func main() {
	ebiten.SetWindowSize(640, 480)
	//ebiten.SetVsyncEnabled(false)
	//ebiten.SetTPS(ebiten.SyncWithFPS)
	//ebiten.SetFPSMode(ebiten.)
	ebiten.SetWindowTitle("Render an image")
	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}

type Component interface {
	getEntity() *Entity
	getData() *data
}

type System interface {
	execute()
}

type SystemDraw interface {
	execute(*ebiten.Image)
	//getArray() genericArray[any]
	appendArray(any)
}

type collide interface {
	collide(image.Rectangle)
}

type event interface {
	execute(enemy *Enemy)
}

type data struct {
	strings      []string
	ints         []int
	floats       []float64
	bools        []bool
	stringArrays [][]string
	intArrays    [][]int
	floatArrays  [][]float64
	boolArrays   [][]bool
}

func (s *data) setString(theStrings ...string) {
	s.strings = theStrings
}

func (s *data) setInt(theInts ...int) {
	s.ints = theInts
}

func (s *data) setFloat(theFloats ...float64) {
	s.floats = theFloats
}

func (s *data) setBool(theBools ...bool) {
	s.bools = theBools
}

func NewDataArg(strings []string, ints []int, floats []float64, bools []bool) *data {
	new := data{}
	new.strings = strings
	new.ints = ints
	new.floats = floats
	new.bools = bools
	return &new
}

func NewDataArgEmpty() *data {
	new := data{}
	return &new
}

func NewDataArgString(theStrings ...string) *data {
	new := data{}
	new.strings = theStrings
	return &new
}

func NewDataArgInt(theInts ...int) *data {
	new := data{}
	new.ints = theInts
	return &new
}

func NewDataArgFloat(theFloats ...float64) *data {
	new := data{}
	new.floats = theFloats
	return &new
}

func NewDataArgBool(theBools ...bool) *data {
	new := data{}
	new.bools = theBools
	return &new
}

func IArr(theInts ...int) []int {
	return theInts
}

func FArr(theFloats ...float64) []float64 {
	return theFloats
}

func SArr(theStrings ...string) []string {
	return theStrings
}

func BArr(theBools ...bool) []bool {
	return theBools
}

/*func GetId(entity *Enemy) int {
	entitys[entityOrder] = entity
	entityOrder++
	return entityOrder -1
}*/
