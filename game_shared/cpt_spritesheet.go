package game_shared

import (
	"jamesraine/grl/engine"
	"jamesraine/grl/engine/parts"
	"jamesraine/grl/engine/render"
	"jamesraine/grl/engine/v"
	"math"

	"golang.org/x/exp/maps"
)

type SpritesheetComponent struct {
	Spritesheet           parts.RTexPacker
	Texture               render.Texture2D
	FlipX                 bool
	FrameTimeMilliseconds int
	Animations            map[string]parts.SpriteAnimation

	curanim parts.SpriteAnimation
}

func NewSpritesheetComponent(sheet parts.RTexPacker, tex render.Texture2D, animations map[string]parts.SpriteAnimation) SpritesheetComponent {
	initialanim := parts.SpriteAnimation{}
	if len(animations) > 0 {
		initialanim = maps.Values(animations)[0]
	}
	return SpritesheetComponent{
		Spritesheet: sheet,
		Texture:     tex,
		Animations:  animations,
		curanim:     initialanim,
	}
}

func (s *SpritesheetComponent) Event(event engine.NodeEvent, gs *engine.Scene, n *engine.Node) {
	if event == engine.NodeEventDraw {
		if len(s.curanim.Frames) == 0 {
			return
		}
		pos := gs.Camera.Transform(n.AbsolutePosition())
		scale := n.AbsoluteScale()
		frametime := s.FrameTimeMilliseconds
		if frametime == 0 {
			frametime = 200
		}
		tms := gs.T * 1000
		framen := int(math.Floor(tms/float64(frametime))) % len(s.curanim.Frames)
		frame := s.curanim.Frames[framen]
		srcX := float32(frame.Position.X)
		srcY := float32(frame.Position.Y)
		srcW := float32(frame.SourceSize.W)
		srcH := float32(frame.SourceSize.H)
		dstX := pos.X - float32(frame.Origin.X)*scale
		dstY := pos.Y - float32(frame.Origin.Y)*scale
		dstW := float32(frame.SourceSize.W) * scale
		dstH := float32(frame.SourceSize.H) * scale
		if s.FlipX {
			srcW *= -1
		}
		render.DrawRect(s.Texture,
			srcX, srcY, srcW, srcH,
			dstX, dstY, dstW, dstH,
			v.White)
	}
}

func (s *SpritesheetComponent) SetAnimation(anim string) {
	s.curanim = s.Animations[anim]
}
