package game_ken

import en "jamesraine/grl/engine"

type SpawnFunc func(*en.Assets) en.Node

var _directory = map[string]SpawnFunc{
	"coin": _spawnCoin,
}

func Spawn(kind string, assets *en.Assets) en.Node {
	spawner, ok := _directory[kind]
	if ok {
		return spawner(assets)
	} else {
		return en.NewNode(kind)
	}
}

func _spawnCoin(assets *en.Assets) en.Node {
	n := en.NewNode("Coin")
	sheet := assets.SpriteSheet("coin.spritesheet")
	tex := assets.Texture(sheet.ImageRef)

	ssComp := en.SpritesheetComponent{
		Spritesheet: sheet,
		Texture:     tex,
	}
	ssComp.SetSprite("idle")
	ssComp.FrameTimeMilliseconds = 100
	en.G.AddComponent(&n, &ssComp)
	return n
}
