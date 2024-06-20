package engine

import (
	"jamesraine/grl/engine/v"
)

type NodeID int64
type NodeEvent int

const (
	NodeEventLoad NodeEvent = iota
	NodeEventUnload
	NodeEventTick
	NodeEventDraw
	NodeEventLateTick
	NodeEventLateDraw
	NodeEventSceneActivate
	NodeEventSceneDeativate
)

type NodeComponent interface {
	Event(NodeEvent, *Scene, *Node)
}

type Node struct {
	Name     string
	Position v.Vec2
	Rotation AngleD
	Scale    float32

	id         NodeID
	Children   []*Node
	Components []NodeComponent
	Parent     *Node
	engine     *Engine
}

func (n *Node) AddChild(c *Node) {
	n.engine.AddChild(n, c)
}

func (n *Node) AddComponent(c NodeComponent) {
	n.engine.AddComponent(n, c)
}

func (n *Node) RemoveFromParent() {
	n.engine.RemoveNodeFromParent(n)
}

func (n *Node) ID() NodeID {
	return n.id
}

var _up = v.V2(0, -1)
var _right = v.V2(1, 0)

func (n *Node) Forward() v.Vec2 {
	r := float32(n.AbsoluteRotation().Rad())
	return _up.Rot(r)
}

func (n *Node) Right() v.Vec2 {
	r := float32(n.AbsoluteRotation().Rad())
	return _right.Rot(r)
}

func (n *Node) AbsoluteRotation() AngleD {
	a := float32(0)
	for tree := n; tree != nil; tree = tree.Parent {
		a += float32(tree.Rotation)
	}
	return AngleD(a)
}

func (n *Node) AbsoluteScale() float32 {
	s := float32(1)
	for tree := n; tree != nil; tree = tree.Parent {
		s *= tree.Scale
	}
	return s
}

func (n *Node) AbsolutePosition() v.Vec2 {
	if n.Parent == nil {
		return n.Position
	} else {
		return n.Position.Xfm(n.Parent.Transform())
	}
}

func (n *Node) Transform() v.Mat {
	m := v.MatrixIdentity()
	for tree := n; tree != nil; tree = tree.Parent {
		m = tree.applyTransform(m)
	}
	return m
}
func (n *Node) applyTransform(m v.Mat) v.Mat {
	m = v.MatrixMultiply(m, v.MatrixScale(n.Scale, n.Scale, 1))
	m = v.MatrixMultiply(m, v.MatrixRotateZ(float32(-n.Rotation.Rad())))
	m = v.MatrixMultiply(m, v.MatrixTranslate(n.Position.X, n.Position.Y, 0))
	return m
}
