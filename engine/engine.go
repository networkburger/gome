package engine

import (
	"jamesraine/grl/engine/parts"
	"jamesraine/grl/engine/window"
)

type NodeEventFunc func(n *Node)
type DeferredAction func()

type Engine struct {
	sceneStack        []*Scene
	scene             *Scene
	nodelock          bool
	queue             []DeferredAction
	WindowPixelHeight int32
	WindowPixelWidth  int32
	parts.Assets
}

func NewEngine(screenW, screenH int32) *Engine {
	e := Engine{
		WindowPixelHeight: screenH,
		WindowPixelWidth:  screenW,
		Assets:            parts.NewAssets("ass"),
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
func (e *Engine) SetScene(scene *Scene) {
	if e.scene != nil {
		e.fireDeepEvent(e.scene.Node, NodeEventUnload)
	}
	if scene.TargetFramerate == 0 {
		scene.TargetFramerate = 30
	}
	scene.Engine = e
	e.scene = scene
	window.SetTargetFPS(scene.TargetFramerate)
	e.fireDeepEvent(scene.Node, NodeEventLoad)
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
