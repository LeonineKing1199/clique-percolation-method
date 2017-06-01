package main

import (
	"fmt"
	"sort"
)

func isEmpty(list NodePtrs) bool {
	return len(list) > 0
}

func printNodePtrs(message string, ptrs NodePtrs) {
	fmt.Println(message)
	for idx := range ptrs {
		fmt.Println(ptrs[idx].userData.name)
	}
}

// BronKerbosch traverse a graph for the maximal cliques
// https://en.wikipedia.org/wiki/Bron%E2%80%93Kerbosch_algorithm
func BronKerbosch(R NodePtrs, P NodePtrs, X NodePtrs) NodePtrs {
	sort.Sort(P)

	fmt.Println("New stack frame")

	printNodePtrs("R is:", R)

	printNodePtrs("P is:", P)

	printNodePtrs("X is:", X)

	fmt.Println("")

	if isEmpty(P) && isEmpty(X) {
		return R
	}

	numVertices := len(P)

	idx := 0

	for numVertices > 0 {
		v := P[idx]

		tmp := BronKerbosch(
			NodePtrSetUnion(R, NodePtrs{v}),
			NodePtrSetIntersection(P, v.neighbors),
			NodePtrSetIntersection(X, v.neighbors))

		if len(tmp) > 0 {
			return tmp
		}

		P = NodePtrSetComplement(P, NodePtrs{v})
		X = NodePtrSetUnion(X, NodePtrs{v})

		sort.Sort(P)

		numVertices = len(P)

		if idx < numVertices && P[idx] == v {
			idx++
		}
	}

	return NodePtrs{}
}
