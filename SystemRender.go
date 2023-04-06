package EasyEC

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type SystemRender struct {
	compRendNPO []*CompRendNPO
}

func NewSystemRender() {
	new := SystemRender{}
	Systems.DrawFirst = append(Systems.DrawFirst, &new)
	SystemsRenderMap["s1"] = &new
}

func (s *SystemRender) execute(screen *ebiten.Image) {
	for _, z := range Components.RendNPO {
		x, y := z.getPosition()
		if (x-G.player.Xpos) > -G.renderDistance && (x-G.player.Xpos) < G.renderDistance && (y-G.player.Ypos) > -G.renderDistance && (y-G.player.Ypos) < G.renderDistance {
			z.Op.GeoM.Reset()
			z.Op.GeoM.Translate(z.Xpos, z.Ypos)
			screen.DrawImage(z.Img, &z.Op)
		}
	}
}

func (s *SystemRender) draw(screen *ebiten.Image, z CompRendNPO) {
	z.Op.GeoM.Reset()
	z.Op.GeoM.Translate(z.Xpos, z.Ypos)
	screen.DrawImage(z.Img, &z.Op)
}

func (s *SystemRender) appendArray(item any) {
	//s.compRendNPO.Add(item.(CompRendNPO))
	s.compRendNPO = append(s.compRendNPO, item.(*CompRendNPO))
}
