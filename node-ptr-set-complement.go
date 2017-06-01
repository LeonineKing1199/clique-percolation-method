package main

// NodePtrSetComplement gathers up all the elements in setA that aren't
// in setB
func NodePtrSetComplement(setA NodePtrs, setB NodePtrs) NodePtrs {
	setBMap := map[*Node]bool{}

	for idx := range setB {
		setBMap[setB[idx]] = true
	}

	nodePtrs := make(NodePtrs, 0, len(setA))

	for idx := range setA {
		if setBMap[setA[idx]] {
			continue
		}

		nodePtrs = append(nodePtrs, setA[idx])
	}

	return nodePtrs
}
