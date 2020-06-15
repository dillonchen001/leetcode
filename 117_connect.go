package main

/**
 * Definition for a Node.
 */
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	var arrNodes []*Node
	arrNodes = append(arrNodes, root)

	start := 0
	for start < len(arrNodes) {
		la := len(arrNodes)
		m := 0
		for i := start; i < la; i++ {
			start++
			if arrNodes[i].Left != nil {
				arrNodes = append(arrNodes, arrNodes[i].Left)
				if m > 0 {
					arrNodes[len(arrNodes)-2].Next = arrNodes[i].Left
				}
				m++
			}

			if arrNodes[i].Right != nil {
				arrNodes = append(arrNodes, arrNodes[i].Right)
				if m > 0 {
					arrNodes[len(arrNodes)-2].Next = arrNodes[i].Right
				}
				m++
			}

		}
	}
	return root
}
