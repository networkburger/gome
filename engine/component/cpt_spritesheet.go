package component

import (
	"jamesraine/grl/engine"
	"jamesraine/grl/engine/parts"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type SpritesheetComponent struct {
	Spritesheet           *parts.Spritesheet
	Texture               rl.Texture2D
	FlipX                 bool
	spritename            string
	nframes               int
	FrameTimeMilliseconds int
}

func (s *SpritesheetComponent) Event(event engine.NodeEvent, gs *engine.Scene, n *engine.Node) {
	if event == engine.NodeEventDraw {
		if len(s.spritename) < 1 {
			return
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
}

func (s *SpritesheetComponent) SetSprite(spritename string) {
	nframes := s.Spritesheet.NumberOfFrames(spritename)
	if nframes > 0 {
		s.spritename = spritename
		s.nframes = nframes
	}
}
