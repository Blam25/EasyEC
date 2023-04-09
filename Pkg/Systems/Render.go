package Systems

import (
	"github.com/hajimehoshi/ebiten/v2"

	C "github.com/Blam25/Test/Pkg/Components"
	//E "github.com/Blam25/Test/Pkg/Entities"
)

var renderDistance float64 = 1000

type Render struct {
	compRendNPO []*C.Rend
}

func NewRender() {
	new := Render{}
	Systems.DrawFirst = append(Systems.DrawFirst, &new)
	SystemsRenderMap["s1"] = &new
}

func (s *Render) Execute(screen *ebiten.Image) {
	//print("hej")
	for _, z := range C.Components.RendNPO {
		z.Op.GeoM.Reset()
		z.Op.GeoM.Translate(
			C.Components.PositionMap[z.Entity.GetId()].Xpos,
			C.Components.PositionMap[z.Entity.GetId()].Ypos,
		)
		screen.DrawImage(z.Img, &z.Op)
		x := C.Components.PositionMap[z.Entity.GetId()].Xpos
		//print(x)
		y := C.Components.PositionMap[z.Entity.GetId()].Ypos
		//print(y)

		if (x-C.Components.PositionMap[C.Components.Player.Entity.GetId()].Xpos) > renderDistance &&
			(x-C.Components.PositionMap[C.Components.Player.Entity.GetId()].Xpos) < renderDistance &&
			(y-C.Components.PositionMap[C.Components.Player.Entity.GetId()].Ypos) > renderDistance &&
			(y-C.Components.PositionMap[C.Components.Player.Entity.GetId()].Ypos) < renderDistance {
			z.Op.GeoM.Reset()
			z.Op.GeoM.Translate(
				C.Components.PositionMap[z.Entity.GetId()].Xpos,
				C.Components.PositionMap[z.Entity.GetId()].Ypos,
			)
			screen.DrawImage(z.Img, &z.Op)
		}
	}
}

func (s *Render) draw(screen *ebiten.Image, z C.Rend) {
	z.Op.GeoM.Reset()
	z.Op.GeoM.Translate(z.Xpos, z.Ypos)
	screen.DrawImage(z.Img, &z.Op)
}

func (s *Render) appendArray(item any) {
	//s.compRendNPO.Add(item.(CompRendNPO))
	s.compRendNPO = append(s.compRendNPO, item.(*C.Rend))
}
