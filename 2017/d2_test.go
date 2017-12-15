package main

import "testing"

type D2DataSet struct {
	spreadsheet [][]int
	checksum    int
}

var d2DataSetA = []D2DataSet{
	{[][]int{
		[]int{5, 1, 9, 5},
		[]int{7, 5, 3},
		[]int{2, 4, 6, 8}}, 18}}

func TestChecsumCalculator(t *testing.T) {
	t.Log("Testing checksum calculator A with dataset: \n", d2DataSetA)

	for _, data := range d2DataSetA {
		if result := CalcChecksum(data.spreadsheet); data.checksum != result {
			t.Errorf("Expected score of %d, but it was %d instead.", data.checksum, result)
		}
	}
}
