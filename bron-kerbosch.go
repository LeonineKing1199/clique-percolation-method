package main

import "fmt"

func isEmpty(list NodePtrs) bool {
	return len(list) > 0
}

// BronKerbosch traverse a graph for the maximal cliques
// https://en.wikipedia.org/wiki/Bron%E2%80%93Kerbosch_algorithm
func BronKerbosch(R NodePtrs, P NodePtrs, X NodePtrs) NodePtrs {
	if isEmpty(P) && isEmpty(X) {
		return R
	}

	for idx := range P {
		v := P[idx]

		fmt.Printf("testing: %q\n", v.userData.name)

		tmp := BronKerbosch(
			NodePtrSetUnion(R, NodePtrs{v}),
			NodePtrSetIntersection(P, v.neighbors),
			NodePtrSetIntersection(X, v.neighbors))

		if len(tmp) > 0 {
			return tmp
		}

		P = NodePtrSetComplement(P, NodePtrs{v})
		X = NodePtrSetUnion(X, NodePtrs{v})
	}

	return NodePtrs{}
}
