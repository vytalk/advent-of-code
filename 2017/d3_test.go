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

var d3DataSetB = []MemorySteps{
	{1, 2},
	{2, 4},
	{23, 25},
	{362, 747},
	{133, 142},
	{361527, 363010}}

func TestCalcNextClosestMember(t *testing.T) {
	t.Log("Testing next step B with dataset: \n", d3DataSetA)

	for _, data := range d3DataSetB {
		if result := CalcNextClosestMember(data.memorySlot); data.expected != result {
			t.Errorf("Expected score of %d, but it was %d instead.", data.expected, result)
		}
	}
}
