package main

import (
	"sort"
	"testing"
)

func TestNodeSortingInterface(t *testing.T) {
	nodes := BuildTestGraph()

	nodePtrs := make(NodePtrs, 0, 8)

	for idx := range nodes {
		nodePtrs = append(nodePtrs, &nodes[idx])
	}

	if len(nodePtrs) != len(nodes) {
		t.Errorf("could not copy node addresses")
	}

	nodePtrs.Swap(0, 3)

	if nodePtrs[0].userData.name != "4" && nodePtrs[3].userData.name != "1" {
		t.Errorf("failed to swap correctly")
	}

	nodePtrs.Swap(1, 5)
	nodePtrs.Swap(2, 6)

	sort.Sort(nodePtrs)

	size := len(nodePtrs)
	for idx := range nodePtrs {
		if idx == size-1 {
			break
		}

		if !nodePtrs.Less(idx, idx+1) {
			t.Errorf("incorrectly sorted array of node pointers")
		}
	}
}
