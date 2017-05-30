package main

// UserData will function as our customization point
type UserData struct {
	name string
}

// Node is a basic node structure that we use for our graph
type Node struct {
	userData  *UserData
	neighbors []*Node
}
