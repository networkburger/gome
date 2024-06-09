package game_ken

import (
	en "jamesraine/grl/engine"
	cm "jamesraine/grl/engine/component"
	pt "jamesraine/grl/engine/parts"
	ph "jamesraine/grl/engine/physics"
)

const (
	CollidableNothing = iota
	CollidableCoin
	CollidableSomethingElse
)

type SpawnFunc func(*pt.Assets, *ph.PhysicsSolver) *en.Node

var _directory = map[string]SpawnFunc{
	"coin": _spawnCoin,
}

func Spawn(kind string, assets *pt.Assets, solver *ph.PhysicsSolver) *en.Node {
	spawner, ok := _directory[kind]
	if ok {
		return spawner(assets, solver)
	} else {
		return en.NewNode(kind)
	}
}

func _spawnCoin(assets *pt.Assets, solver *ph.PhysicsSolver) *en.Node {
	n := en.NewNode("Coin")
	sheet := assets.SpriteSheet("coin.spritesheet")
	tex := assets.Texture(sheet.ImageRef)

	ssComp := cm.SpritesheetComponent{
		Spritesheet: sheet,
		Texture:     tex,
	}
	ssComp.SetSprite("idle")
	ssComp.FrameTimeMilliseconds = 100
	en.G.AddComponent(n, &ssComp)

	signal := cm.PhysicsSignalComponent{
		PhysicsManager: solver,
		Radius:         16,
		Kind:           CollidableCoin,
	}
	en.G.AddComponent(n, &signal)
	return n
}
