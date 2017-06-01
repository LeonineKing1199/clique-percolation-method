package main

// NodePtrSetIntersection finds the intersection between two
// slices of node pointers
// intersection = elements common to both sets
func NodePtrSetIntersection(setA NodePtrs, setB NodePtrs) NodePtrs {
	setAMap := map[*Node]bool{}
	setBMap := map[*Node]bool{}

	for idx := range setA {
		setAMap[setA[idx]] = true
	}

	for idx := range setB {
		setBMap[setB[idx]] = true
	}

	intersectingNodePtrs := map[*Node]bool{}

	for idx := range setA {
		if setBMap[setA[idx]] {
			intersectingNodePtrs[setA[idx]] = true
		}
	}

	for idx := range setB {
		if setAMap[setB[idx]] {
			intersectingNodePtrs[setB[idx]] = true
		}
	}

	nodePtrs := make(NodePtrs, 0, len(setA)+len(setB))

	for key := range intersectingNodePtrs {
		nodePtrs = append(nodePtrs, key)
	}

	return nodePtrs
}
