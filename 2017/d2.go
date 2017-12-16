package main

/*
The spreadsheet consists of rows of apparently-random numbers. To make sure the
recovery process is on the right track, they need you to calculate the spreadsheet's
checksum. For each row, determine the difference between the largest value and the
smallest value; the checksum is the sum of all of these differences.
*/
func CalcChecksum(spreadsheet [][]int) int {
	c := make(chan int)

	for _, v := range spreadsheet {
		go findRowDiff(v, c)
	}

	sum := 0
	for range spreadsheet {
		sum += <-c
	}
	return sum
}

func findRowDiff(row []int, c chan int) {
	if len(row) <= 1 {
		c <- 0
		return
	}

	largest := row[0]
	smallest := row[0]
	for i := 1; i < len(row); i++ {
		if largest < row[i] {
			largest = row[i]
		}
		if smallest > row[i] {
			smallest = row[i]
		}
	}
	c <- largest - smallest
}
