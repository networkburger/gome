package engine

func (e *Engine) Tick(gs *GameState) {
	gs.Camera.cache()
	tick(gs, e.scene)
}

func (e *Engine) Draw(gs *GameState) {
	gs.Camera.cache()
	draw(gs, e.scene)
}

func tick(gs *GameState, n *Node) {
	for i := 0; i < len(n.Components); i++ {
		n.Components[i].Tick(gs, n)
	}
	for i := 0; i < len(n.Children); i++ {
		tick(gs, n.Children[i])
	}
}
func draw(gs *GameState, n *Node) {
	for i := 0; i < len(n.Components); i++ {
		n.Components[i].Draw(gs, n)
	}
	for i := 0; i < len(n.Children); i++ {
		draw(gs, n.Children[i])
	}
}
