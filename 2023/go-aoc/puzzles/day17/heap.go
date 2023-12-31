package day17

import "container/heap"

type NodeQueue []*node

func (nq NodeQueue) Len() int { return len(nq) }
func (nq NodeQueue) Less(i, j int) bool {
	return nq[i].cost < nq[j].cost
}

func (nq NodeQueue) Swap(i, j int) {
	nq[i], nq[j] = nq[j], nq[i]
	nq[i].index = i
	nq[j].index = j
}

func (nq *NodeQueue) Push(x any) {
	n := len(*nq)
	item := x.(*node)
	item.index = n
	*nq = append(*nq, item)
}

func (nq *NodeQueue) Pop() any {
	old := *nq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*nq = old[0 : n-1]
	return item
}

func (nq *NodeQueue) update(node *node, cost int) {
	node.cost = cost
	heap.Fix(nq, node.index)
}
