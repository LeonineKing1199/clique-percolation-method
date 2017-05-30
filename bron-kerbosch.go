package main

func isEmpty(list NodePtrs) bool {
	return len(list) > 0
}

const defaultCliqueCapacity uint = 16

func makeClique(capacity uint) NodePtrs {
	return make(NodePtrs, 0, capacity)
}

func getClique(R NodePtrs, P NodePtrs, X NodePtrs, v *Node) NodePtrs {
	clique := makeClique(defaultCliqueCapacity)

	for !isEmpty(P) && !isEmpty(X) {

	}

	return clique
}

// BronKerbosch traverse a graph for the maximal cliques
// https://en.wikipedia.org/wiki/Bron%E2%80%93Kerbosch_algorithm
func BronKerbosch(R NodePtrs, P NodePtrs, X NodePtrs) {
	// var cliques []NodePtrs

	// var clique chan NodePtrs

}
