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

type SpawnFunc func(*parts.Assets, *physics.PhysicsSolver) *engine.Node

var _directory = map[string]SpawnFunc{
	"coin": _spawnCoin,
}

func Spawn(kind string, assets *parts.Assets, solver *physics.PhysicsSolver) *engine.Node {
	spawner, ok := _directory[kind]
	if ok {
		return spawner(assets, solver)
	} else {
		return engine.NewNode(kind)
	}
}

func _spawnCoin(assets *parts.Assets, solver *physics.PhysicsSolver) *engine.Node {
	n := engine.NewNode("Coin")
	sheet := assets.SpriteSheet("coin.spritesheet")
	tex := assets.Texture(sheet.ImageRef)

	ssComp := component.SpritesheetComponent{
		Spritesheet: sheet,
		Texture:     tex,
	}
	ssComp.SetSprite("idle")
	ssComp.FrameTimeMilliseconds = 100
	engine.G.AddComponent(n, &ssComp)

	signal := component.PhysicsSignalComponent{
		PhysicsManager: solver,
		Radius:         16,
		Kind:           CollidableCoin,
	}
	engine.G.AddComponent(n, &signal)
	return n
}
