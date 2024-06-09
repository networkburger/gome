package component

import (
	"fmt"
	en "jamesraine/grl/engine"
	pt "jamesraine/grl/engine/parts"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type SpritesheetComponent struct {
	Spritesheet           *pt.Spritesheet
	Texture               rl.Texture2D
	FlipX                 bool
	spritename            string
	nframes               int
	FrameTimeMilliseconds int
}

func (s *SpritesheetComponent) Event(e en.NodeEvent, n *en.Node) {}

func (s *SpritesheetComponent) Tick(gs *en.GameState, n *en.Node) {
	if len(s.spritename) < 1 {
		return
	}
	if n.Name == "Player" {
		fmt.Println("Player")
	}
	pos := gs.Camera.Transform(n.AbsolutePosition())
	a := n.AbsoluteRotation()
	scale := n.AbsoluteScale()
	frametime := s.FrameTimeMilliseconds
	if frametime == 0 {
		frametime = 200
	}
	tms := gs.T * 1000
	framen := int(math.Floor(tms/float64(frametime))) % s.nframes
	frame := s.Spritesheet.GetFrame(s.spritename, framen)
	destRect := rl.NewRectangle(
		pos.X-frame.Origin.X*scale,
		pos.Y-frame.Origin.Y*scale,
		frame.Source.Width*scale,
		frame.Source.Height*scale)
	srcRect := frame.Source
	if s.FlipX {
		srcRect.Width *= -1
	}
	rl.DrawTexturePro(s.Texture, srcRect, destRect, rl.Vector2{}, float32(a.Rad()), rl.White)
}

func (s *SpritesheetComponent) SetSprite(spritename string) {
	nframes := s.Spritesheet.NumberOfFrames(spritename)
	if nframes > 0 {
		s.spritename = spritename
		s.nframes = nframes
	}
}