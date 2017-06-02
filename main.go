package main

func main() {
	// simple slice of graph nodes
	nodes := BuildTestGraph()

	P := NodePtrs{}

	for idx := range nodes {
		P = append(P, &nodes[idx])
	}

	cliques := BronKerbosch(NodePtrs{}, P, NodePtrs{})

	for idx := range cliques {

		PrintNodePtrs("Clique", cliques[idx])
	}
}
