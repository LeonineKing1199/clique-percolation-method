package main

// NodePtrSetUnion takes two slices of node pointers and then
// retrieves the union of them, preserving only unique elements
func NodePtrSetUnion(setA NodePtrs, setB NodePtrs) NodePtrs {
	nodePtrsMap := map[*Node]bool{}

	for idx := range setA {
		nodePtrsMap[setA[idx]] = true
	}

	for idx := range setB {
		nodePtrsMap[setB[idx]] = true
	}

	nodePtrs := make(NodePtrs, 0, len(nodePtrsMap))

	for key := range nodePtrsMap {
		nodePtrs = append(nodePtrs, key)
	}

	return nodePtrs
}
