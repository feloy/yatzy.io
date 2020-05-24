package yatzy

import (
	"reflect"
	"testing"
)

func TestCountByValue(t *testing.T) {
	var tests = []struct {
		Die    []int
		Result []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 1, 1, 1, 1, 0}},
		{[]int{1, 1, 1, 1, 1}, []int{5, 0, 0, 0, 0, 0}},
		{[]int{6, 6, 6, 6, 6}, []int{0, 0, 0, 0, 0, 5}},
	}

	for _, test := range tests {
		res := countByValue(test.Die)
		if !reflect.DeepEqual(res, test.Result) {
			t.Fatalf("countByValue of %v should be %v but is %v\n", test.Die, test.Result, res)
		}
	}
}

func TestComputePoints(t *testing.T) {
	var tests = []struct {
		Formula int
		Die     []int
		Result  int
	}{
		{ACES, []int{1, 3, 4, 1, 2}, 2},
		{SIXES, []int{6, 3, 6, 1, 6}, 18},
		{SIXES, []int{6, 6, 6, 6, 6}, 30},
		// three in a row
		{THREE_IN_A_ROW, []int{1, 3, 4, 5, 1}, 0},
		{THREE_IN_A_ROW, []int{2, 2, 4, 5, 2}, 6},
		{THREE_IN_A_ROW, []int{2, 2, 2, 5, 2}, 8},
		// Four in a row
		{FOUR_IN_A_ROW, []int{1, 3, 1, 5, 1}, 0},
		{FOUR_IN_A_ROW, []int{2, 2, 2, 5, 2}, 8},
		{FOUR_IN_A_ROW, []int{3, 3, 3, 3, 3}, 15},
		// Full house
		{FULL_HOUSE, []int{1, 2, 2, 1, 2}, 25},
		{FULL_HOUSE, []int{1, 2, 3, 1, 2}, 0},
		{FULL_HOUSE, []int{3, 3, 3, 3, 3}, 25},
		// Small straight
		{SMALL_STRAIGHT, []int{5, 3, 1, 2, 2}, 0},
		{SMALL_STRAIGHT, []int{5, 3, 4, 2, 2}, 30},
		{SMALL_STRAIGHT, []int{5, 3, 4, 2, 6}, 30},
		// Large straight
		{LARGE_STRAIGHT, []int{5, 3, 1, 2, 2}, 0},
		{LARGE_STRAIGHT, []int{5, 3, 4, 2, 2}, 0},
		{LARGE_STRAIGHT, []int{5, 3, 4, 2, 6}, 40},
		{LARGE_STRAIGHT, []int{2, 3, 1, 5, 4}, 40},
		// Yatzy
		{YATZY, []int{2, 1, 2, 2, 2}, 0},
		{YATZY, []int{2, 2, 2, 2, 2}, 50},
		// Chance
		{CHANCE, []int{2, 4, 6, 1, 2}, 15},
	}

	for _, test := range tests {
		res := ComputePoints(test.Formula, test.Die)
		if res != test.Result {
			t.Fatalf("computePoints of %d for %v should be %d but is %d\n", test.Formula, test.Die, test.Result, res)
		}
	}
}
