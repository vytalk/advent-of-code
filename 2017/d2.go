package main

/*
The spreadsheet consists of rows of apparently-random numbers. To make sure the
recovery process is on the right track, they need you to calculate the spreadsheet's
checksum. For each row, determine the difference between the largest value and the
smallest value; the checksum is the sum of all of these differences.
*/
func CalcChecksum(spreadsheet [][]int) int {
	return 0
}

func findRowDiff(row []int) {
	smallest := 0
	largest := 0
	for _, v := range row {
		if v > largest {
			largset = v
		}
		if v < smallest {
			smallest = v
		}
	}
	return largest - smallest
}
