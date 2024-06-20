package component

import (
	"jamesraine/grl/engine"
	"jamesraine/grl/engine/v"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Sprite struct {
	Texture rl.Texture2D
	Origin  v.Vec2
}

func NewSprite(tex rl.Texture2D) Sprite {
	return Sprite{
		Texture: tex,
		Origin:  v.V2(float32(tex.Width)/2, float32(tex.Height)/2),
	}
}

func (s Sprite) Draw(position v.Vec2, rotation engine.AngleD, scale float32) {
	p := position.Sub(s.Origin)
	rl.DrawTextureEx(s.Texture, rl.NewVector2(p.X, p.Y), float32(rotation), scale, rl.White)
}

type SpriteNode struct {
	Sprite Sprite
}

func (s *SpriteNode) Tick(gs *engine.Scene, n *engine.Node) {
	pos := v.V2(0, 0).Xfm(n.Transform())
	a := n.AbsoluteRotation()
	s.Sprite.Draw(pos, a, n.Scale)
}

type Billboard struct {
	Texture  rl.Texture2D
	SrcRect  rl.Rectangle
	DstRect  rl.Rectangle
	Origin   v.Vec2
	Rotation float32
	Tint     rl.Color
}

func (s *Billboard) Event(e engine.NodeEvent, gs *engine.Scene, n *engine.Node) {
	if e == engine.NodeEventDraw {

		pos := gs.Camera.Transform(n.AbsolutePosition())
		dr := s.DstRect
		dr.X = pos.X
		dr.Y = pos.Y
		or := rl.NewVector2(s.Origin.X, s.Origin.Y)
		rl.DrawTexturePro(s.Texture, s.SrcRect, dr, or, s.Rotation, s.Tint)
	}
}

func NewBillboard(tex rl.Texture2D) Billboard {
	return Billboard{
		Texture:  tex,
		SrcRect:  rl.NewRectangle(0, 0, float32(tex.Width), float32(tex.Height)),
		DstRect:  rl.NewRectangle(0, 0, float32(tex.Width), float32(tex.Height)),
		Origin:   v.V2(0, 0),
		Rotation: 0,
		Tint:     rl.White,
	}
}
