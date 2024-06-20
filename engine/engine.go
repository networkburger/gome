package engine

type NodeEventFunc func(n *Node)

type Engine struct {
	stack []*Node
	scene *Node
}

func (e *Engine) Scene() *Node {
	return e.scene
}
func (e *Engine) PushScene(n *Node) {
	if e.scene != nil {
		e.fireDeepEvent(e.scene, NodeEventSceneDeativate)
	}
	e.stack = append(e.stack, e.scene)
	e.scene = n

	e.fireDeepEvent(n, NodeEventSceneActivate)
	e.fireDeepEvent(n, NodeEventLoad)
}

func (c *Engine) PopScene() {
	if c.scene != nil {
		c.fireDeepEvent(c.scene, NodeEventUnload)
		c.fireDeepEvent(c.scene, NodeEventSceneDeativate)
	}
	if len(c.stack) > 0 {
		c.scene = c.stack[len(c.stack)-1]
		c.stack = c.stack[:len(c.stack)-1]
		c.fireDeepEvent(c.scene, NodeEventSceneActivate)
	} else {
		c.scene = nil
	}
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
