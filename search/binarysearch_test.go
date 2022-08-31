package main

import (
	"testing"
)

func TestSearch(t *testing.T) {

	testCases := []struct {
		name   string
		arr    []int
		target int
		expect int
	}{
		{
			name:   "example 1",
			arr:    []int{-1, 0, 3, 5, 9, 12},
			target: 9,
			expect: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			index := Search(tc.arr, tc.target)

			if index != tc.expect {
				t.Errorf("got %d, expect %d", index, tc.expect)
			}

			t.Logf("%d exist, index is %d", tc.target, index)
		})
	}
}
