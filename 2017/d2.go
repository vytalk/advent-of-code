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

/*
It sounds like the goal is to find the only two numbers in each row where one
evenly divides the other - that is, where the result of the division operation
is a whole number. They would like you to find those numbers on each line,
divide them, and add up each line's result
*/
func SumDivNumbers(spreadsheet [][]int) int {
	c := make(chan int)

	for _, v := range spreadsheet {
		go findDivNumbers(v, c)
	}

	sum := 0
	for range spreadsheet {
		sum += <-c
	}
	return sum
}

func findDivNumbers(row []int, c chan int) {
	if len(row) <= 1 {
		c <- 0
		return
	}

	for i := 0; i < len(row)-1; i++ {
		for j := i + 1; j < len(row); j++ {
			if row[i] > row[j] {
				if row[i]%row[j] == 0 {
					c <- row[i] / row[j]
					return
				}
			} else {
				if row[j]%row[i] == 0 {
					c <- row[j] / row[i]
					return
				}
			}
		}
	}
	c <- 0
	return
}
