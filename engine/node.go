package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type NodeID int64
type NodeEvent int

const (
	NodeEventLoad NodeEvent = iota
	NodeEventUnload
)

type NodeComponent interface {
	Tick(*GameState, *Node)
	Event(NodeEvent, *Node)
}

type Node struct {
	Name     string
	Position rl.Vector2
	Rotation AngleD
	Scale    float32

	id         NodeID
	Children   []*Node
	Components []NodeComponent
	Parent     *Node
}

var _nid = 0

func NewNode(name string) Node {
	_nid = _nid + 1
	return Node{
		id:    NodeID(_nid),
		Name:  name,
		Scale: 1,
	}
}

func (n *Node) ID() NodeID {
	return n.id
}

var _up = rl.NewVector2(0, -1)
var _right = rl.NewVector2(1, 0)

func (n Node) Forward() rl.Vector2 {
	r := float32(n.AbsoluteRotation().Rad())
	return rl.Vector2Rotate(_up, r)
}

func (n Node) Right() rl.Vector2 {
	r := float32(n.AbsoluteRotation().Rad())
	return rl.Vector2Rotate(_right, r)
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

func (n *Node) AbsolutePosition() rl.Vector2 {
	if n.Parent == nil {
		return n.Position
	} else {
		return rl.Vector2Transform(rl.NewVector2(0, 0), n.Parent.Transform())
	}
}

func (n *Node) Transform() rl.Matrix {
	m := rl.MatrixIdentity()
	for tree := n; tree != nil; tree = tree.Parent {
		m = tree.applyTransform(m)
	}
	return m
}
func (n Node) applyTransform(m rl.Matrix) rl.Matrix {
	m = rl.MatrixMultiply(m, rl.MatrixScale(n.Scale, n.Scale, 1))
	m = rl.MatrixMultiply(m, rl.MatrixRotateZ(float32(-n.Rotation.Rad())))
	m = rl.MatrixMultiply(m, rl.MatrixTranslate(n.Position.X, n.Position.Y, 0))
	return m
}
