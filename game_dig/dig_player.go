package game_dig

import (
	en "jamesraine/grl/engine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type PlayerStats struct {
	Speed        float32
	ReverseSpeed float32
	TurnSpeed    float32
	MaxHealth    int
}

func StartingPlayerStats() PlayerStats {
	return PlayerStats{
		Speed:     30,
		TurnSpeed: 50,
		MaxHealth: 100,
	}
}

type Player struct {
	Stats      PlayerStats
	Health     int
	Ballistics *en.BallisticComponent
}

func (s *Player) Event(e en.NodeEvent, n *en.Node) {}

func (p *Player) Tick(gs *en.GameState, node *en.Node) {
	en.ProcessInputs(InputOverworld, func(action en.ActionID, power float32) {
		switch action {
		case Accelerate:
			p.Ballistics.Impulse = rl.Vector2Scale(node.Forward(), power*p.Stats.Speed)
		case TurnLeft:
			p.Ballistics.Torque = p.Stats.TurnSpeed * -power
		case TurnRight:
			p.Ballistics.Torque = p.Stats.TurnSpeed * power
		case Turn:
			p.Ballistics.Torque = p.Stats.TurnSpeed * power
		}
	})
}

func StandardPlayerNode() *en.Node {
	player := Player{
		Stats:  StartingPlayerStats(),
		Health: 100,
	}
	playerNode := en.NewNode("Player")
	en.G.AddComponent(&playerNode, &player)

	polyNode := NewLineStripComponent(rl.White, []rl.Vector2{
		rl.NewVector2(0, -5),
		rl.NewVector2(10, 40),
		rl.NewVector2(-10, 40),
		rl.NewVector2(0, -5),
	})
	hull := en.NewNode("Hull")
	en.G.AddComponent(&hull, &polyNode)
	en.G.AddChild(&playerNode, &hull)

	ballistics := en.BallisticComponent{
		VelocityDamping: 0.4,
		AngularDamping:  0.8,
	}
	ballisticsNode := en.NewNode("Ballistics")
	en.G.AddComponent(&ballisticsNode, &ballistics)
	player.Ballistics = &ballistics
	en.G.AddChild(&playerNode, &ballisticsNode)

	return &playerNode
}
