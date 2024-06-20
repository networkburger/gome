package engine

func (e *Engine) LoopEvent(event NodeEvent, gs *GameState) {
	if event == NodeEventTick {
		gs.Camera.cache()
	}
	send(event, gs, e.scene)
}

func send(e NodeEvent, gs *GameState, n *Node) {
	for i := 0; i < len(n.Components); i++ {
		n.Components[i].Event(e, gs, n)
	}
	for i := 0; i < len(n.Children); i++ {
		send(e, gs, n.Children[i])
	}
}
