package main

import "fmt"

func main() {
	// simple slice of graph nodes
	nodes := BuildTestGraph()

	P := NodePtrs{}

	for idx := range nodes {
		P = append(P, &nodes[idx])
	}

	clique := BronKerbosch(NodePtrs{}, P, NodePtrs{})

	for idx := range clique {
		fmt.Printf("%q\n", clique[idx].userData.name)
	}
}
