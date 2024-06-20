package engine

import (
	"jamesraine/grl/engine/util"
	"log/slog"
	"slices"
)

type nodeActionAdd struct {
	Parent *Node
	Child  *Node
}
type nodeActionRemove struct {
	*Node
}
type nodeActionAddComponent struct {
	Parent    *Node
	Component NodeComponent
}
type nodeActionRemoveComponent struct {
	Parent    *Node
	Component NodeComponent
}

func (e *Engine) AddChild(p *Node, c *Node) {
	if c.Parent != nil {
		slog.Warn("Node already has a parent")
		return
	}
	if IsDescendant(c, p) {
		slog.Warn("Node is an ancestor of parent")
		return
	}

	inScene := IsDescendant(e.scene.RootNode, p)
	if inScene && e.nodelock {
		e.queue = append(e.queue, nodeActionAdd{Parent: p, Child: c})
		return
	}

	p.Children = append(p.Children, c)
	c.Parent = p

	// defer firing events until the node is fully attached to the tree
	if inScene {
		e.fireDeepEvent(c, NodeEventLoad)
	}
}

func (e *Engine) fireDeepEvent(n *Node, event NodeEvent) {
	tree := delve(n)
	for _, n := range tree {
		for _, comp := range n.Components {
			comp.Event(event, e.scene, n)
		}
	}
}

func (e *Engine) AddComponent(n *Node, c NodeComponent) {
	// if the parent node is NOT part of the scene, we don't need to
	// pay attention to the lock state - events won't be fired anyway
	if IsDescendant(e.scene.RootNode, n) {
		if e.nodelock {
			e.queue = append(e.queue, nodeActionAddComponent{Parent: n, Component: c})
			return
		}
		n.Components = append(n.Components, c)
		c.Event(NodeEventLoad, e.scene, n)
	} else {
		n.Components = append(n.Components, c)
	}
}

func (e *Engine) RemoveNodeFromParent(killnode *Node) {
	if e.nodelock {
		e.queue = append(e.queue, nodeActionRemove{Node: killnode})
		return
	}

	if IsDescendant(e.scene.RootNode, killnode) {
		e.fireDeepEvent(killnode, NodeEventUnload)
	}

	if killnode.Parent != nil {
		index := util.SliceIndexOf(killnode.Parent.Children, killnode)
		if index == -1 {
			slog.Warn("Node not found in parent's children")
		} else {
			killnode.Parent.Children = util.SliceRemoveIndex(killnode.Parent.Children, index)
		}
		killnode.Parent = nil
	}
}

func (e *Engine) RemoveComponentFromNode(n *Node, c NodeComponent) {
	if e.nodelock {
		e.queue = append(e.queue, nodeActionRemoveComponent{Parent: n, Component: c})
		return
	}

	index := slices.Index(n.Components, c)
	n.Components = util.SliceRemoveIndex(n.Components, index)

	if IsDescendant(e.scene.RootNode, n) {
		c.Event(NodeEventUnload, e.scene, n)
	}
}

// delve returns a list of all nodes in the tree rooted at n, depth-first
func delve(n *Node) []*Node {
	nodes := make([]*Node, 0)
	for _, c := range n.Children {
		nodes = append(nodes, delve(c)...)
	}
	nodes = append(nodes, n)
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

func FindComponent[T NodeComponent](components []NodeComponent) (T, bool) {
	for _, c := range components {
		if c, ok := c.(T); ok {
			return c, true
		}
	}
	var zeroValue T
	return zeroValue, false
}
