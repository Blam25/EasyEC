package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type SystemRender struct {
	compRendNPO []*CompRendNPO
}

func NewSystemRender() {
	new := SystemRender{}
	systems.drawFirst = append(systems.drawFirst, &new)
	systemsRenderMap["s1"] = &new
}

func (s *SystemRender) execute(screen *ebiten.Image) {
	for _, z := range components.RendNPO {
		x, y := z.getPosition()
		if (x-g.player.xpos) > -g.renderDistance && (x-g.player.xpos) < g.renderDistance && (y-g.player.ypos) > -g.renderDistance && (y-g.player.ypos) < g.renderDistance {
			z.draw(screen)
		}
	}
}

func (s *SystemRender) appendArray(item any) {
	//s.compRendNPO.Add(item.(CompRendNPO))
	s.compRendNPO = append(s.compRendNPO, item.(*CompRendNPO))

}
