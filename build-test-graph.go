package main

type graphStruct struct {
	userData UserData
	indices  []uint
}

func makeUserData(userName string, ids ...uint) graphStruct {
	return graphStruct{userData: UserData{userName}, indices: ids}
}

// BuildTestGraph is a small helper function that we can use
// to auto-generate a dummy graph to be used for testing our
// CPM algorithm
func BuildTestGraph() []Node {
	graphData := []struct {
		userData UserData
		indices  []uint
	}{
		makeUserData("1", 2, 1, 7),
		makeUserData("2", 0, 2, 7, 3, 4, 5),
		makeUserData("3", 0, 1, 6),
		makeUserData("4", 1, 5, 4),
		makeUserData("5", 1, 3, 5),
		makeUserData("6", 1, 3, 4, 6),
		makeUserData("7", 2, 5),
		makeUserData("8", 0, 1),
	}

	nodes := make([]Node, len(graphData))

	for idx := range graphData {
		nodes[idx].userData = &(graphData[idx].userData)

		indices := graphData[idx].indices
		tmp := []*Node{}

		for _, index := range indices {
			tmp = append(tmp, &(nodes[index]))
		}

		nodes[idx].neighbors = tmp
	}

	return nodes
}
