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

/*
	So, the first few squares' values are chosen as follows:

    Square 1 starts with the value 1.
    Square 2 has only one adjacent filled square (with value 1), so it also stores 1.
    Square 3 has both of the above squares as neighbors and stores the sum of their values, 2.
    Square 4 has all three of the aforementioned squares as neighbors and stores the sum of their values, 4.
    Square 5 only has the first and fourth squares as neighbors, so it gets the value 5.

	Once a square is written, its value does not change. Therefore, the first few squares would receive the following values:

	147  142  133  122   59
	304    5    4    2   57
	330   10    1    1   54
	351   11   23   25   26
	362  747  806--->   ...

	What is the first value written that is larger than your puzzle input?
*/

const initSize int = 15

var calcMatrix = [][]int{[]int{-1, -1}, []int{-1, 0}, []int{-1, 1}, []int{0, -1}, []int{0, 0}, []int{0, 1}, []int{1, -1}, []int{1, 0}, []int{1, 1}}

func CalcNextClosestMember(search int) int {
	if initSize%2 == 0 {
		return 0
	}

	currPoss := []int{initSize / 2, initSize / 2}

	var matrix [initSize][initSize]int
	matrix[currPoss[0]][currPoss[1]] = 1

	for stop := 0; stop < 1000; stop++ {

		nextPossition(currPoss, matrix)

		sum := 0
		for i := 0; i < len(calcMatrix); i++ {
			sum += matrix[currPoss[0]+calcMatrix[i][0]][currPoss[1]+calcMatrix[i][1]]
		}

		if sum > search {
			return sum
		}

		matrix[currPoss[0]][currPoss[1]] = sum
	}
	return 0
}

func nextPossition(currPoss []int, matrix [initSize][initSize]int) {
	right := matrix[currPoss[0]+calcMatrix[5][0]][currPoss[1]+calcMatrix[5][1]]
	up := matrix[currPoss[0]+calcMatrix[1][0]][currPoss[1]+calcMatrix[1][1]]
	left := matrix[currPoss[0]+calcMatrix[3][0]][currPoss[1]+calcMatrix[3][1]]
	down := matrix[currPoss[0]+calcMatrix[7][0]][currPoss[1]+calcMatrix[7][1]]

	if right == 0 && up != 0 {
		currPoss[0] += calcMatrix[5][0]
		currPoss[1] += calcMatrix[5][1]
	} else if up == 0 && left != 0 {
		currPoss[0] += calcMatrix[1][0]
		currPoss[1] += calcMatrix[1][1]
	} else if left == 0 && down != 0 {
		currPoss[0] += calcMatrix[3][0]
		currPoss[1] += calcMatrix[3][1]
	} else if down == 0 && right != 0 {
		currPoss[0] += calcMatrix[7][0]
		currPoss[1] += calcMatrix[7][1]
	} else {
		currPoss[0] += calcMatrix[5][0]
		currPoss[1] += calcMatrix[5][1]
	}
}
