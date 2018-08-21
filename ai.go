package tic

import "math"

type node struct {
	children []node
	value    float64
}

func (n node) isTerminal() bool {
	return n.children == nil || len(n.children) == 0
}

func minimax(node node, depth int, maximizingPlayer bool) float64 {
	if depth == 0 || node.isTerminal() {
		return node.value
	}
	if maximizingPlayer {
		value := math.Inf(-1)
		for _, child := range node.children {
			value = math.Max(value, minimax(child, depth-1, false))
		}
		return value
	}
	value := math.Inf(1)
	for _, child := range node.children {
		value = math.Min(value, minimax(child, depth-1, true))
	}
	return value
}
