package main

/*
Each square on the grid is allocated in a spiral pattern starting at a location
marked 1 and then counting up while spiraling outward. For example, the first
few squares are allocated like this:

17  16  15  14  13
18   5   4   3  12
19   6   1   2  11
20   7   8   9  10
21  22  23---> ...
While this is very space-efficient (no squares are skipped), requested data must
 be carried back to square 1 (the location of the only access port for this
 memory system) by programs that can only move up, down, left, or right. They
 always take the shortest path: the Manhattan Distance between the location of
 the data and square 1.
*/
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
