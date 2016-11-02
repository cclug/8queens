package main

import "testing"

func TestIsValidHorizontal(t *testing.T) {
	var tests = []struct {
		input []int
		want  bool
	}{
		{[]int{0, 2, 4, 4}, false},
		{[]int{4, 1, 3, 4}, false},
		{[]int{3, 3}, false},

		{[]int{1, 2, 3}, true},
		{[]int{1, 2, 3, 0}, true},
	}

	for _, test := range tests {
		if got := isValidHorizontal(test.input); got != test.want {
			t.Errorf("fail: %v\n", test.input)
		}
	}
}

func TestIsValidDiagonal(t *testing.T) {
	var tests = []struct {
		input []int
		want  bool
	}{
		{[]int{0, 1, 2, 3, 4}, false},
		{[]int{0, 2, 1}, false},

		{[]int{0, 2, 4, 1, 3}, true},
	}

	for _, test := range tests {
		if got := isValidDiagonal(test.input); got != test.want {
			t.Errorf("fail: %v\n", test.input)
		}
	}
}
