package game_ken

import (
	"jamesraine/grl/engine"
	"jamesraine/grl/engine/component"
	"jamesraine/grl/engine/parts"
	"jamesraine/grl/engine/physics"
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
		return e.NewNode(kind)
	}
}

func _spawnCoin(e *engine.Engine, assets *parts.Assets) *engine.Node {
	n := e.NewNode("Coin")
	sheet, _ := assets.SpriteSheet("coin.spritesheet")
	tex := assets.Texture(sheet.ImagePath)

	ssComp := component.SpritesheetComponent{
		Spritesheet: sheet,
		Texture:     tex,
	}
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
