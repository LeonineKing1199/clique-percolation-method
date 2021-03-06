package main

import "fmt"

// UserData will function as our customization point
type UserData struct {
	name string
}

// Node is a basic node structure that we use for our graph
type Node struct {
	userData  *UserData
	neighbors []*Node
}

// NodePtrs is a simple slice of Node pointers
type NodePtrs []*Node

func (nodePtrs NodePtrs) Len() int { return len(nodePtrs) }

func (nodePtrs NodePtrs) Swap(a, b int) {
	nodePtrs[a], nodePtrs[b] = nodePtrs[b], nodePtrs[a]
}
func (nodePtrs NodePtrs) Less(a, b int) bool {
	return nodePtrs[a].userData.name < nodePtrs[b].userData.name
}

// PrintNodePtrs is a simple utility function that prints a message
// and then prints the name of each node in the slice
func PrintNodePtrs(message string, ptrs NodePtrs) {
	fmt.Println(message)
	for idx := range ptrs {
		fmt.Println(ptrs[idx].userData.name)
	}
	fmt.Println()
}
