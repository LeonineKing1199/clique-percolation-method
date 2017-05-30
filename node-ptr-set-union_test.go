package main

import "testing"

func TestNodePtrSetUnion(t *testing.T) {
	nodes := BuildTestGraph()

	if len(nodes) != 8 {
		t.Errorf("Incorrect size")
		return
	}

	setA := make(NodePtrs, 0, len(nodes)/2)
	setB := make(NodePtrs, 0, len(nodes)/2)

	for idx := range nodes {
		if idx < len(nodes)/2 {
			setA = append(setA, &nodes[idx])
		} else {
			setA = append(setA, &nodes[idx])
			setB = append(setB, &nodes[idx])
		}
	}

	if len(setA) != len(nodes) {
		t.Errorf("Not enough node pointers in setA")
		return
	}

	setUnion := NodePtrSetUnion(setA, setB)

	if len(setUnion) != len(nodes) {
		t.Errorf("incorrect set union")
		return
	}
}
