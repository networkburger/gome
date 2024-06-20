package game_dig

import (
	"jamesraine/grl/engine"
	"jamesraine/grl/engine/parts"
	"jamesraine/grl/engine/physics"
	"jamesraine/grl/engine/v"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type digScene struct {
	parts.Assets
	*engine.Engine
	physics.PhysicsSolver
	player *engine.Node
}

func (s *digScene) Event(event engine.NodeEvent, gs *engine.GameState, n *engine.Node) {
	switch event {
	case engine.NodeEventSceneActivate:
		rl.SetTargetFPS(90)
	case engine.NodeEventDraw:
		rl.ClearBackground(rl.NewColor(18, 65, 68, 255))
	case engine.NodeEventTick:
		gs.Camera.Position.X = s.player.Position.X - 500 // screenw / 2
		gs.Camera.Position.Y = s.player.Position.Y - 300 // screenh / 2
	case engine.NodeEventLateTick:
		s.PhysicsSolver.Solve(gs)
	}
}

func DigScene(e *engine.Engine) *engine.Node {
	k := digScene{}
	k.Assets = parts.NewAssets("ass")
	k.PhysicsSolver = physics.NewPhysicsSolver(func(b *engine.Node, s *engine.Node) {
		// something hit something
	})

	rootNode := e.NewNode("RootNode")
	rootNode.AddComponent(&k)

	mapNode := NewDigMap(e, &k.PhysicsSolver, &k.Assets)

	k.player = StandardPlayerNode(e, &k.PhysicsSolver)
	k.player.Position = v.V2(985*20, 35*20)
	k.player.Rotation = 90

	rootNode.AddChild(mapNode)
	rootNode.AddChild(k.player)
	return rootNode
}
