package main

import "testing"

func TestNodePtrSetComplement(t *testing.T) {
	nodes := BuildTestGraph()

	setA := make(NodePtrs, 0, 3)
	setB := make(NodePtrs, 0, 3)

	for idx := 0; idx < 3; idx++ {
		setA = append(setA, &nodes[idx])
		setB = append(setB, &nodes[idx+1])
	}

	complementAB := NodePtrSetComplement(setA, setB)
	complementBA := NodePtrSetComplement(setB, setA)

	if len(complementAB) != 1 || len(complementBA) != 1 {
		t.Error("incorrect complement length")
		return
	}

	if complementAB[0].userData.name != "1" {
		t.Errorf("Expected %q, got %q", "1", complementAB[0].userData.name)
		return
	}

	if complementBA[0].userData.name != "4" {
		t.Errorf("Expected %q, got %q", "4", complementBA[0].userData.name)
		return
	}

}
