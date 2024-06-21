package game_dig

import (
	"jamesraine/grl/engine"
	"jamesraine/grl/engine/component"
	"jamesraine/grl/engine/parts"
	"jamesraine/grl/engine/physics"
	"jamesraine/grl/engine/v"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewDigMap(e *engine.Engine) *engine.Node {
	bgSprite := component.NewBillboard(e.Assets.Texture("bg.png"))
	mapSprite := component.NewBillboard(e.Assets.Texture("map.png"))
	mapPixels := e.Assets.Pixels("map.png")

	baseSize := v.V2(float32(mapSprite.Texture.Width), float32(mapSprite.Texture.Height))
	worldScale := float32(20)
	worldSize := baseSize.Scl(worldScale)
	worldRect := rl.NewRectangle(0, 0, worldSize.X, worldSize.Y)

	bgSprite.DstRect = worldRect
	mapSprite.DstRect = worldRect

	obstacle := physics.PhysicsObstacleComponent{
		CollisionSurfaceProvider: &PixelObstacleProvider{
			PixelBuffer: mapPixels,
		},
	}

	mapNode := e.NewNode("Map")
	mapNode.Scale = worldScale
	mapNode.AddComponent(&bgSprite)
	mapNode.AddComponent(&mapSprite)
	mapNode.AddComponent(&obstacle)

	return mapNode
}

type PixelObstacleProvider struct {
	parts.PixelBuffer
}

func (p *PixelObstacleProvider) Surfaces(n *engine.Node, pos v.Vec2, radius float32, hits []physics.CollisionSurface, nhits *int) {
	sc := n.AbsoluteScale()
	sx := int32(pos.X / sc)
	sy := int32(pos.Y / sc)
	sradius := int32(radius / sc)
	left := int32(sx - sradius)
	top := int32(sy - sradius)
	right := int32(sx + sradius)
	bottom := int32(sy + sradius)
	w := p.PixelBuffer.Image.Width
	h := p.PixelBuffer.Image.Height

	if right < 0 || left >= w || bottom < 0 || top >= h {
		return
	}

	if left < 0 {
		left = 0
	}
	if top < 0 {
		top = 0
	}
	if right >= w {
		right = w - 1
	}
	if bottom >= h {
		bottom = h - 1
	}

	for y := top; y <= bottom; y++ {
		for x := left; x <= right; x++ {
			if p.PixelBuffer.Pixels[y*w+x].A > 0 {
				blockRect := rl.NewRectangle(float32(x)*sc, float32(y)*sc, sc, sc)
				physics.GenHitsForSquare(pos, radius, blockRect, physics.SurfaceProperties{
					Friction:    0,
					Restitution: 0.5,
				}, hits, nhits)
			}
		}
	}
}
