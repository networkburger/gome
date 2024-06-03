package engine

func (e *Engine) Run(gs *GameState) {
	run(gs, e.scene)
}

func run(gs *GameState, n *Node) {
	for i := 0; i < len(n.Components); i++ {
		n.Components[i].Tick(gs, n)
	}
	for i := 0; i < len(n.Children); i++ {
		run(gs, n.Children[i])
	}
}
