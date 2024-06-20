package engine

type NodeEventFunc func(n *Node)

type Engine struct {
	scene *Node
}

func (e *Engine) Scene() *Node {
	return e.scene
}
func (e *Engine) SetScene(n *Node) {
	e.scene = n
	e.fireLoadEvents(n)
}

var _nid = 0

func (e *Engine) NewNode(name string) *Node {
	_nid = _nid + 1
	return &Node{
		engine: e,
		id:     NodeID(_nid),
		Name:   name,
		Scale:  1,
	}
}
