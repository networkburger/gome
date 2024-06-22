package component

import (
	"jamesraine/grl/engine"
	"jamesraine/grl/engine/parts"
	"jamesraine/grl/engine/render"
	"jamesraine/grl/engine/v"
	"math"
)

type SpritesheetComponent struct {
	Spritesheet           *parts.Spritesheet
	Texture               render.Texture2D
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
		scale := n.AbsoluteScale()
		frametime := s.FrameTimeMilliseconds
		if frametime == 0 {
			frametime = 200
		}
		tms := gs.T * 1000
		framen := int(math.Floor(tms/float64(frametime))) % s.nframes
		frame := s.Spritesheet.GetFrame(s.spritename, framen)

		srcRect := frame.Source
		if s.FlipX {
			srcRect.W *= -1
		}
		render.DrawRect(s.Texture,
			srcRect.X, srcRect.Y, srcRect.W, srcRect.H,
			pos.X-frame.Origin.X*scale,
			pos.Y-frame.Origin.Y*scale,
			frame.Source.W*scale,
			frame.Source.H*scale,
			v.White)
	}
}

func (s *SpritesheetComponent) SetSprite(spritename string) {
	nframes := s.Spritesheet.NumberOfFrames(spritename)
	if nframes > 0 {
		s.spritename = spritename
		s.nframes = nframes
	}
}
