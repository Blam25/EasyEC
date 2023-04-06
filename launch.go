package EasyEC

import (
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var G Game
var entityOrder int = 1
var test int = 1.0
var Systems *ECSystems
var SystemsMap map[string]System
var SystemsRenderMap map[string]SystemDraw
var Components *ECComponents

//var entities []*Entity
//var renderedNonPlayers []*renderedNonPlayer

//var op *ebiten.DrawImageOptions

func init() {
	Components = &ECComponents{}
	Systems = &ECSystems{}
	SystemsMap = make(map[string]System)
	SystemsRenderMap = make(map[string]SystemDraw)

	initSystems()
	initComponents()

	//toDataArg([]string{"hej"})

	G = NewGame()
	G.compRendNPOMap = make(map[int]*CompRendNPO)
	//sprite1 := New_Sprite()

	/*entity1 := NewEntity()
	print(&entity1)
	entity1.
		With(NewCompRendNPO, NewDataArg(nil, nil, FArr(500, 500), nil)).
		With(NewCompCollider, &Data{Strings: SArr("hej", "yo"), Ints: IArr(9, 2)})

	entity2 := NewEntity().With(
		NewCompRendNPO, NewDataArgFloat(100, 100)).With(
		NewCompCollider, NewDataArgEmpty()).With(
		NewCompMoveRand, &Data{})

	entity4 := NewEntity().With(
		NewCompRendNPO, NewDataArgFloat(400, 400)).With(
		NewCompMoveRand, NewDataArgEmpty()).With(
		NewCompCollider, NewDataArgEmpty())
	print(&entity4)

	NewEntity().With(
		NewCompRendNPO, NewDataArgFloat(400, 400)).With(
		NewCompCollider, NewDataArgEmpty())
	//print(&entity5)

	entity6 := NewEntity().
		With(NewCompRendNPO, NewDataArgFloat(400, 400)).With(
		NewCompCollider, NewDataArgEmpty())

	print(entity6)

	//		with(NewCompRendNPO, NewDataArgFloat(1000, 1000)).
	//						with(NewCompCollider, &dataArgs{})
	print(&entity2)
	//entity2.(NewCompRendNPO, NewDataArgFloat(1000, 1000)).
	//with(NewCompCollider, &dataArgs{})

	entity3 := NewEntity()
	entity3.With(
		NewCompRendNPO, NewDataArgFloat(1000, 0)).With(
		NewCompCollider, &Data{})

	//NewCompRenderedNonPlayer(entity2, 1000, 1000)
	Components.Player = sprite1
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
	Components = &ECComponents{}
	Components.CollisionMap = make(map[int]*CompCollision)
	Components.EventCollisionMap = make(map[int][]InteractionEvent)
}

func initSystems() {
	NewSystemRender()
	NewSysMoveWithCam()
	NewSysCollide()
	NewSysCollision()
	NewSysClear()
}

func main() {
	ebiten.SetWindowSize(640, 480)
	//ebiten.SetVsyncEnabled(false)
	//ebiten.SetTPS(ebiten.SyncWithFPS)
	//ebiten.SetFPSMode(ebiten.)
	ebiten.SetWindowTitle("Render an image")
	if err := ebiten.RunGame(&G); err != nil {
		log.Fatal(err)
	}
}

type Component interface {
	getEntity() *Entity
	getData() *Data
}

type System interface {
	Execute()
}

type SystemDraw interface {
	execute(*ebiten.Image)
	//getArray() genericArray[any]
	//appendArray(any)
}

type collide interface {
	collide(image.Rectangle)
}

type event interface {
	execute(enemy *Enemy)
}

type Data struct {
	Strings      []string
	Ints         []int
	Floats       []float64
	Bools        []bool
	StringArrays [][]string
	IntArrays    [][]int
	FloatArrays  [][]float64
	BoolArrays   [][]bool
}

func (s *Data) setString(theStrings ...string) {
	s.Strings = theStrings
}

func (s *Data) setInt(theInts ...int) {
	s.Ints = theInts
}

func (s *Data) setFloat(theFloats ...float64) {
	s.Floats = theFloats
}

func (s *Data) setBool(theBools ...bool) {
	s.Bools = theBools
}

func NewDataArg(strings []string, ints []int, floats []float64, bools []bool) *Data {
	new := Data{}
	new.Strings = strings
	new.Ints = ints
	new.Floats = floats
	new.Bools = bools
	return &new
}

func NewDataArgEmpty() *Data {
	new := Data{}
	return &new
}

func NewDataArgString(theStrings ...string) *Data {
	new := Data{}
	new.Strings = theStrings
	return &new
}

func NewDataArgInt(theInts ...int) *Data {
	new := Data{}
	new.Ints = theInts
	return &new
}

func NewDataArgFloat(theFloats ...float64) *Data {
	new := Data{}
	new.Floats = theFloats
	return &new
}

func NewDataArgBool(theBools ...bool) *Data {
	new := Data{}
	new.Bools = theBools
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
