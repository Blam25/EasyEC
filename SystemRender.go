package EasyEC

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type SystemRender struct {
	compRendNPO []*CompRend
}

func NewSystemRender() {
	new := SystemRender{}
	Systems.DrawFirst = append(Systems.DrawFirst, &new)
	SystemsRenderMap["s1"] = &new
}

func (s *SystemRender) execute(screen *ebiten.Image) {
	for _, z := range Components.RendNPO {
		x := Components.PositionMap[z.Entity.GetId()].Xpos
		y := Components.PositionMap[z.Entity.GetId()].Ypos

		if (x-Components.PositionMap[Components.Player.entity.GetId()].Xpos) > -G.renderDistance &&
			(x-Components.PositionMap[Components.Player.entity.GetId()].Xpos) < G.renderDistance &&
			(y-Components.PositionMap[Components.Player.entity.GetId()].Ypos) > -G.renderDistance &&
			(y-Components.PositionMap[Components.Player.entity.GetId()].Ypos) < G.renderDistance {
			z.Op.GeoM.Reset()
			z.Op.GeoM.Translate(
				Components.PositionMap[z.Entity.GetId()].Xpos,
				Components.PositionMap[z.Entity.GetId()].Ypos,
			)
			screen.DrawImage(z.Img, &z.Op)
		}
	}
}

func (s *SystemRender) draw(screen *ebiten.Image, z CompRend) {
	z.Op.GeoM.Reset()
	z.Op.GeoM.Translate(z.Xpos, z.Ypos)
	screen.DrawImage(z.Img, &z.Op)
}

func (s *SystemRender) appendArray(item any) {
	//s.compRendNPO.Add(item.(CompRendNPO))
	s.compRendNPO = append(s.compRendNPO, item.(*CompRend))
}
