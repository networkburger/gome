package engine

import "jamesraine/grl/engine/v"

type NodeEventFunc func(n *Node)

type Engine struct {
	sceneStack                          []*Scene
	scene                               *Scene
	nodelock                            bool
	queue                               []any
	WindowPixelHeight, WindowPixelWidth int
}

func NewEngine(screenW, screenH int) *Engine {
	e := Engine{
		WindowPixelHeight: screenH,
		WindowPixelWidth:  screenW,
		scene: &Scene{
			RootNode: &Node{},
		},
	}
	e.scene.RootNode = &Node{
		engine: &e,
	}
	return &e
}

func (e *Engine) Scene() *Scene {
	return e.scene
}
func (e *Engine) PushScene(sceneNode *Node) {
	scene := Scene{
		G:        e,
		RootNode: sceneNode,
		Camera: &Camera{
			Position: v.R(0, 0, float32(e.WindowPixelWidth), float32(e.WindowPixelHeight)),
		},
	}

	if e.scene != nil {
		e.fireDeepEvent(e.scene.RootNode, NodeEventSceneDeativate)
	}
	e.sceneStack = append(e.sceneStack, e.scene)
	e.scene = &scene

	e.fireDeepEvent(scene.RootNode, NodeEventSceneActivate)
	e.fireDeepEvent(scene.RootNode, NodeEventLoad)
}

func (c *Engine) PopScene() {
	if c.scene != nil {
		c.fireDeepEvent(c.scene.RootNode, NodeEventUnload)
		c.fireDeepEvent(c.scene.RootNode, NodeEventSceneDeativate)
	}
	if len(c.sceneStack) > 0 {
		c.scene = c.sceneStack[len(c.sceneStack)-1]
		c.sceneStack = c.sceneStack[:len(c.sceneStack)-1]
		c.fireDeepEvent(c.scene.RootNode, NodeEventSceneActivate)
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

func (e *Engine) Lock() {
	e.nodelock = true
}
func (e *Engine) Unlock() {
	e.nodelock = false

	if len(e.queue) > 0 {
		for _, action := range e.queue {
			addChild, ok := action.(nodeActionAdd)
			if ok {
				e.AddChild(addChild.Parent, addChild.Child)
				continue
			}
			removeChild, ok := action.(nodeActionRemove)
			if ok {
				e.RemoveNodeFromParent(removeChild.Node)
				continue
			}
			addComponent, ok := action.(nodeActionAddComponent)
			if ok {
				e.AddComponent(addComponent.Parent, addComponent.Component)
				continue
			}
			removeComponent, ok := action.(nodeActionRemoveComponent)
			if ok {
				e.RemoveComponentFromNode(removeComponent.Parent, removeComponent.Component)
				continue
			}
		}
		e.queue = e.queue[:0]
	}
}

func (e *Engine) LoopEvent(event NodeEvent) {
	if event == NodeEventTick {
		e.scene.Camera.cache()
	}
	send(event, e.scene, e.scene.RootNode)
}

func send(e NodeEvent, s *Scene, n *Node) {
	for i := 0; i < len(n.Components); i++ {
		n.Components[i].Event(e, s, n)
	}
	for i := 0; i < len(n.Children); i++ {
		send(e, s, n.Children[i])
	}
}
