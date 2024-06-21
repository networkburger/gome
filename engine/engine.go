package engine

type NodeEventFunc func(n *Node)
type DeferredAction func()

type Engine struct {
	sceneStack        []*Scene
	scene             *Scene
	nodelock          bool
	queue             []DeferredAction
	WindowPixelHeight int
	WindowPixelWidth  int
}

func NewEngine(screenW, screenH int) *Engine {
	e := Engine{
		WindowPixelHeight: screenH,
		WindowPixelWidth:  screenW,
		scene: &Scene{
			Node: &Node{},
		},
	}
	e.scene.Node = &Node{
		engine: &e,
	}
	return &e
}

func (e *Engine) Scene() *Scene {
	return e.scene
}
func (e *Engine) PushScene(scene *Scene) {
	if e.scene != nil {
		e.fireDeepEvent(e.scene.Node, NodeEventSceneDeativate)
	}
	scene.Engine = e
	e.sceneStack = append(e.sceneStack, e.scene)
	e.scene = scene

	e.fireDeepEvent(scene.Node, NodeEventSceneActivate)
	e.fireDeepEvent(scene.Node, NodeEventLoad)
}

func (c *Engine) PopScene() {
	if c.scene != nil {
		c.fireDeepEvent(c.scene.Node, NodeEventUnload)
		c.fireDeepEvent(c.scene.Node, NodeEventSceneDeativate)
	}
	if len(c.sceneStack) > 0 {
		c.scene = c.sceneStack[len(c.sceneStack)-1]
		c.sceneStack = c.sceneStack[:len(c.sceneStack)-1]
		c.fireDeepEvent(c.scene.Node, NodeEventSceneActivate)
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

func (e *Engine) Enqueue(action DeferredAction) {
	e.queue = append(e.queue, action)
}

func (e *Engine) Lock() {
	e.nodelock = true
}

func (e *Engine) Unlock() {
	e.nodelock = false

	if len(e.queue) > 0 {
		for _, action := range e.queue {
			action()
		}
		e.queue = e.queue[:0]
	}
}

func (e *Engine) LoopEvent(event NodeEvent) {
	if event == NodeEventTick {
		e.scene.Camera.cache()
	}
	send(event, e.scene, e.scene.Node)
}

func send(e NodeEvent, s *Scene, n *Node) {
	for i := 0; i < len(n.Components); i++ {
		n.Components[i].Event(e, s, n)
	}
	for i := 0; i < len(n.Children); i++ {
		send(e, s, n.Children[i])
	}
}
