package main

func CountMemorySteps(slot int) int {
	lvl := 0
	curr := 1
	edge := 1

	for ; curr < slot; lvl++ {
		edge += 2
		curr = edge * edge
	}

	// we need last element of edge not 2 in the other edge
	edge -= 1
	for curr > slot {
		curr -= edge
	}

	return slot - (curr + edge/2) + lvl
}
