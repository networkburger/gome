package component

import (
	"jamesraine/grl/engine"
	"jamesraine/grl/engine/render"
	"jamesraine/grl/engine/v"
)

type Billboard struct {
	Texture  render.Texture2D
	SrcRect  v.Rect
	DstRect  v.Rect
	Origin   v.Vec2
	Rotation float32
	Tint     v.Color
}

func (s *Billboard) Event(e engine.NodeEvent, gs *engine.Scene, n *engine.Node) {
	if e == engine.NodeEventDraw {

		pos := gs.Camera.Transform(n.AbsolutePosition())
		dr := s.DstRect
		dr.X = pos.X
		dr.Y = pos.Y
		render.DrawRect(s.Texture,
			s.SrcRect.X, s.SrcRect.Y, s.SrcRect.W, s.SrcRect.H,
			dr.X, dr.Y, dr.W, dr.H,
			s.Tint)
		//or := rl.NewVector2(s.Origin.X, s.Origin.Y)
		//rl.DrawTexturePro(s.Texture, s.SrcRect, dr, or, s.Rotation, s.Tint)
	}
}

func NewBillboard(tex render.Texture2D) Billboard {
	return Billboard{
		Texture:  tex,
		SrcRect:  v.R(0, 0, float32(tex.Width), float32(tex.Height)),
		DstRect:  v.R(0, 0, float32(tex.Width), float32(tex.Height)),
		Origin:   v.V2(0, 0),
		Rotation: 0,
		Tint:     v.White,
	}
}
