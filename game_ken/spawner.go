package game_ken

import (
	"jamesraine/grl/engine"
	"jamesraine/grl/engine/parts"
	"jamesraine/grl/engine/physics"
	"jamesraine/grl/game_shared"
)

const (
	CollidableNothing = iota
	CollidableCoin
	CollidableSomethingElse
)

type SpawnFunc func(*engine.Engine, *parts.Assets) *engine.Node

var _directory = map[string]SpawnFunc{
	"coin": _spawnCoin,
}

func Spawn(e *engine.Engine, kind string, assets *parts.Assets) *engine.Node {
	spawner, ok := _directory[kind]
	if ok {
		return spawner(e, assets)
	} else {
		return nil
	}
}

func _spawnCoin(e *engine.Engine, assets *parts.Assets) *engine.Node {
	n := e.NewNode("Coin")
	sheet, _ := assets.SpriteSheet("coin.json")
	tex := assets.Texture(sheet.ImagePath)

	ssComp := game_shared.NewSpritesheetComponent(sheet, tex, map[string]parts.SpriteAnimation{
		"idle": parts.NewSpriteAnimation(sheet, "coin"),
	})
	ssComp.SetAnimation("idle")
	ssComp.FrameTimeMilliseconds = 100
	n.AddComponent(&ssComp)

	signal := physics.PhysicsSignalComponent{
		Radius: 16,
		Kind:   CollidableCoin,
	}
	n.AddComponent(&signal)
	return n
}
