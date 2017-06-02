package main

import (
	"fmt"
	"sort"
)

func isEmpty(list NodePtrs) bool {
	return len(list) == 0
}

const verbose = true

// BronKerbosch traverse a graph for the maximal cliques
// https://en.wikipedia.org/wiki/Bron%E2%80%93Kerbosch_algorithm
func BronKerbosch(R NodePtrs, P NodePtrs, X NodePtrs) []NodePtrs {
	sort.Sort(P)

	fmt.Println("New stack frame")

	// PrintNodePtrs("R is:", R)

	// PrintNodePtrs("P is:", P)

	// PrintNodePtrs("X is:", X)

	if isEmpty(P) && isEmpty(X) {
		// PrintNodePtrs("found maximal clique!", R)
		// fmt.Println()
		return []NodePtrs{R}
	}

	numVertices := len(P)

	cliques := []NodePtrs{}

	for numVertices > 0 {
		v := P[0]

		fmt.Printf("Vertex is: %q\n", v.userData.name)

		newR := NodePtrSetUnion(R, NodePtrs{v})
		newP := NodePtrSetIntersection(P, v.neighbors)
		newX := NodePtrSetIntersection(X, v.neighbors)

		// PrintNodePtrs("newR:", newR)
		// PrintNodePtrs("newP:", newP)
		// PrintNodePtrs("newX:", newX)

		tmp := BronKerbosch(newR, newP, newX)

		if len(tmp) > 0 {
			for idx := range tmp {
				cliques = append(cliques, tmp[idx])
			}
		}

		P = NodePtrSetComplement(P, NodePtrs{v})
		X = NodePtrSetUnion(X, NodePtrs{v})

		sort.Sort(P)

		numVertices = len(P)
	}

	fmt.Println("")
	return cliques
}
