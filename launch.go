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
	sprite1 := New_Sprite()
	entity1 := NewEntity()
	NewCompRenderedNonPlayer(entity1, 500, 500)
	entity2 := NewEntity()
	NewCompRenderedNonPlayer(entity2, 1000, 1000)
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

/*func GetId(entity *Enemy) int {
	entitys[entityOrder] = entity
	entityOrder++
	return entityOrder -1
}*/
