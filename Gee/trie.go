package Gee

import (
	"strings"
)

type node struct {
	pattern  string // Route to be matched
	part     string // part of route
	children []*node
	isWild   bool // whether is a wildcard
}

// get first matching children node to insert
func (n *node) getFirstMatchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part {
			return child
		}
		if child.isWild {
			panic("router conflicts")
		}
	}
	return nil
}

// get all matching children node to search
func (n *node) getAllMatchChild(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		// fmt.Println(n)
		return
	}

	part := parts[height]
	child := n.getFirstMatchChild(part)
	if child == nil {
		child = &node{
			part:   part,
			isWild: part[0] == ':' || part[0] == '*',
		}
		if child.isWild && len(n.children) > 0 {
			panic("new path '" + part + "' conflicts with existing path '" + n.children[0].part + "'")
		}

		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	childs := n.getAllMatchChild(part)

	for _, child := range childs {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}
