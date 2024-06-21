package game_dig

import (
	"jamesraine/grl/engine"
	"jamesraine/grl/engine/parts"
	"jamesraine/grl/engine/physics"
	"jamesraine/grl/engine/v"
	"jamesraine/grl/game_shared"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type digScene struct {
	parts.Assets
	*engine.Engine
	player *engine.Node
}

func (s *digScene) Event(event engine.NodeEvent, gs *engine.Scene, n *engine.Node) {
	switch event {
	case engine.NodeEventSceneActivate:
		rl.SetTargetFPS(90)
	case engine.NodeEventDraw:
		rl.ClearBackground(rl.NewColor(18, 65, 68, 255))
	case engine.NodeEventTick:
		gs.Camera.Position.X = s.player.Position.X - float32(gs.Engine.WindowPixelWidth)/2
		gs.Camera.Position.Y = s.player.Position.Y - float32(gs.Engine.WindowPixelHeight)/2

		if rl.IsKeyPressed(rl.KeyEscape) {
			game_shared.ShowPauseMenu(gs, &s.Assets)
		}
	}
}

func DigScene(e *engine.Engine) *engine.Scene {
	k := digScene{}
	k.Assets = parts.NewAssets("ass")
	solver := physics.NewPhysicsSolver(func(b *engine.Node, s *engine.Node) {
		// something hit something
	})

	rootNode := e.NewNode("RootNode - Dig")
	rootNode.AddComponent(&k)

	mapNode := NewDigMap(e, &k.Assets)

	k.player = StandardPlayerNode(e)
	k.player.Position = v.V2(985*20, 35*20)
	k.player.Rotation = 90

	rootNode.AddChild(mapNode)
	rootNode.AddChild(k.player)

	return &engine.Scene{
		Node:    rootNode,
		Physics: &solver,
	}
}
