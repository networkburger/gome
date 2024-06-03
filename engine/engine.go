package engine

// Global ENGINE instance
var G *Engine

type NodeEventFunc func(n *Node)

type Engine struct {
	scene *Node
}

func NewEngine() {
	G = &Engine{}
}

func (e *Engine) Scene() *Node {
	return e.scene
}
func (e *Engine) SetScene(n *Node) {
	e.scene = n
	e.fireLoadEvents(n)
}
