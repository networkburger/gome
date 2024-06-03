package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Sprite struct {
	Texture rl.Texture2D
	Origin  rl.Vector2
}

func NewSprite(path string) Sprite {
	tex := rl.LoadTexture(path)
	return Sprite{
		Texture: tex,
		Origin:  rl.NewVector2(float32(tex.Width)/2, float32(tex.Height)/2),
	}
}

func (s Sprite) Draw(position rl.Vector2, rotation AngleD, scale float32) {
	rl.DrawTextureEx(s.Texture, rl.Vector2Subtract(position, s.Origin), float32(rotation), scale, rl.White)
}

type SpriteNode struct {
	Sprite Sprite
}

func (s *SpriteNode) Tick(gs *GameState, n *Node) {
	pos := rl.Vector2Transform(rl.NewVector2(0, 0), n.Transform())
	a := n.AbsoluteRotation()
	s.Sprite.Draw(pos, a, n.Scale)
}

type Billboard struct {
	Texture  rl.Texture2D
	SrcRect  rl.Rectangle
	DstRect  rl.Rectangle
	Origin   rl.Vector2
	Rotation float32
	Tint     rl.Color
}

func (s *Billboard) Event(e NodeEvent, n *Node) {}

func (s *Billboard) Tick(gs *GameState, n *Node) {
	pos := n.AbsolutePosition()
	dr := s.DstRect
	dr.X = pos.X
	dr.Y = pos.Y
	rl.DrawTexturePro(s.Texture, s.SrcRect, dr, s.Origin, s.Rotation, s.Tint)
}

func NewBillboard(path string) Billboard {
	tex := rl.LoadTexture(path)
	return Billboard{
		Texture:  tex,
		SrcRect:  rl.NewRectangle(0, 0, float32(tex.Width), float32(tex.Height)),
		DstRect:  rl.NewRectangle(0, 0, float32(tex.Width), float32(tex.Height)),
		Origin:   rl.NewVector2(0, 0),
		Rotation: 0,
		Tint:     rl.White,
	}
}
