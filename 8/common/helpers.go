package common

func GetNode(nodes *map[string]*Node, nodeLabel string) *Node {
	n, ok := (*nodes)[nodeLabel]
	if !ok {
		n := &Node{}
		(*nodes)[nodeLabel] = n
		return n
	}
	return n
}
