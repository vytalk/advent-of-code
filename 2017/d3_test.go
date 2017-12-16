package main

import "testing"

type MemorySteps struct {
	memorySlot int
	expected   int
}

var d3DataSetA = []MemorySteps{
	{1, 0},
	{12, 3},
	{23, 2},
	{1024, 31},
	{361527, 326}}

func TestCalculateSteps(t *testing.T) {
	t.Log("Testing step calculator A with dataset: \n", d3DataSetA)

	for _, data := range d3DataSetA {
		if result := CountMemorySteps(data.memorySlot); data.expected != result {
			t.Errorf("Expected score of %d, but it was %d instead.", data.expected, result)
		}
	}
}
