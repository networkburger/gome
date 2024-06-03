package engine

import (
	"log/slog"
	"slices"
)

func (e *Engine) AddChild(p *Node, c *Node) {
	if c.Parent != nil {
		slog.Warn("Node already has a parent")
		return
	}
	if IsDescendant(c, p) {
		slog.Warn("Node is an ancestor of parent")
		return
	}

	p.Children = append(p.Children, c)
	c.Parent = p

	// If we're adding to a node that's already in the scene, trigger load events
	if IsDescendant(e.scene, p) {
		e.fireLoadEvents(c)
	}
}

func (e *Engine) fireLoadEvents(n *Node) {
	tree := delve(n)
	for _, n := range tree {
		for _, comp := range n.Components {
			comp.Event(NodeEventLoad, n)
		}
	}
}

func (e *Engine) AddComponent(n *Node, c NodeComponent) {
	n.Components = append(n.Components, c)
	if IsDescendant(e.scene, n) {
		c.Event(NodeEventLoad, n)
	}
}

func (e *Engine) RemoveNodeFromParent(killnode *Node) {
	nodeTree := delve(killnode)

	if IsDescendant(e.scene, killnode) {
		// remove nodes from 'bottom to top' so that our notfications
		// fire for each node in a sensible order
		for i := len(nodeTree) - 1; i >= 0; i-- {
			n := nodeTree[i]
			for componentIndex, component := range n.Components {
				n.Components = SliceRemoveIndex(n.Components, componentIndex)
				component.Event(NodeEventUnload, n)
			}
			index := slices.Index(n.Parent.Children, n)
			n.Parent.Children = SliceRemoveIndex(n.Parent.Children, index)
			n.Parent = nil
		}
	}
}

func (e *Engine) RemoveComponentFromParent(n *Node, c NodeComponent) {
	index := slices.Index(n.Parent.Children, n)
	n.Components = SliceRemoveIndex(n.Components, index)

	if IsDescendant(e.scene, n) {
		c.Event(NodeEventUnload, n)
	}
}

// delve returns a list of all nodes in the tree rooted at n
func delve(n *Node) []*Node {
	nodes := make([]*Node, 0)
	nodes = append(nodes, n)
	for _, c := range n.Children {
		nodes = append(nodes, delve(c)...)
	}
	return nodes
}

func IsDescendant(parent, node *Node) bool {
	if parent == nil || node == nil {
		return false
	}
	if parent == node {
		return true
	}
	return IsDescendant(parent, node.Parent)
}

func FindComponent[T any](components []NodeComponent) *T {
	for _, c := range components {
		if c, ok := c.(T); ok {
			return &c
		}
	}
	return nil
}
