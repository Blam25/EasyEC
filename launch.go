package main

import (
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var g Game
var entityOrder int = 1
//var entities []*Entity
//var renderedNonPlayers []*renderedNonPlayer

//var op *ebiten.DrawImageOptions

func init() {
	g = NewGame()
	g.compRendNPOMap = make(map[int]*CompRendNPO)
	sprite1 := New_Sprite()
	
	entity1 := NewEntity()
	print(&entity1)
	entity1.
	with(NewCompRendNPO, NewDataArgs(nil, nil, []float64{500, 500}, nil)).
	with(NewCompCollider, &dataArgs{})
	
	entity2 := NewEntity()
	print(&entity2)
	entity2.with(
		NewCompRendNPO, NewArgFloat(1000, 1000)).with(
		NewCompCollider, &dataArgs{})
		
	entity3 := NewEntity()
	entity3.with(
		NewCompRendNPO, NewArgFloat(1000, 0)).with(
		NewCompCollider, &dataArgs{})
		
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

type component interface {
	getEntity() *Entity
}

type collide interface {
	collide(image.Rectangle)
}

type event interface {
	execute(enemy *Enemy)
}

type dataArgs struct {
	strings []string
	ints []int
	floats []float64
	bools	[]bool
	
}

func (s *dataArgs) setString(theStrings ...string) {
	s.strings = theStrings
}

func (s *dataArgs) setInt(theInts ...int) {
	s.ints = theInts
}

func (s *dataArgs) setFloat(theFloats ...float64) {
	s.floats = theFloats
}

func (s *dataArgs) setBool(theBools ...bool) {
	s.bools = theBools
}

func NewDataArgs(strings []string, ints []int, floats []float64, bools []bool) *dataArgs {
	new := dataArgs{}
	new.strings = strings
	new.ints = ints
	new.floats = floats
	new. bools = bools
	return &new
}

func NewArgString(theStrings ...string) *dataArgs{
	new := dataArgs{}
	new.strings = theStrings
	return &new
}

func NewArgInt(theInts ...int) *dataArgs {
	new := dataArgs{}
	new.ints = theInts
	return &new
}

func NewArgFloat(theFloats ...float64) *dataArgs {
	new := dataArgs{}
	new.floats = theFloats
	return &new
}

func NewArgBool(theBools ...bool) *dataArgs {
	new := dataArgs{}
	new.bools = theBools
	return &new
}

/*func GetId(entity *Enemy) int {
	entitys[entityOrder] = entity
	entityOrder++
	return entityOrder -1
}*/
