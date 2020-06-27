package main

/**
 * Definition for a Node.
 */
type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	lookup := make(map[*Node]*Node)
	return dfs(node, lookup)
}

func dfs(node *Node, lookup map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}

	if _, ok := lookup[node]; ok {
		return lookup[node]
	}

	clone := &Node{node.Val, nil}
	lookup[node] = clone

	for _, n := range node.Neighbors {
		clone.Neighbors = append(clone.Neighbors, dfs(n, lookup))
	}
	return clone
}
