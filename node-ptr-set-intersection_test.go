package main

import (
	"sort"
	"testing"
)

func TestNodePtrSetIntersection(t *testing.T) {
	nodes := BuildTestGraph()

	setA := NodePtrs{}
	setB := NodePtrs{}

	for idx := 0; idx < len(nodes)/2; idx++ {
		setA = append(setA, &nodes[idx])
		setB = append(setB, &nodes[idx+1])
	}

	intersection := NodePtrSetIntersection(setA, setB)

	sort.Sort(intersection)

	expectedNames := []string{
		"2", "3", "4",
	}

	for idx := range intersection {
		if intersection[idx].userData.name != expectedNames[idx] {
			t.Errorf("Expected %q, got %q\n", expectedNames[idx], intersection[idx].userData.name)
			return
		}
	}
}
