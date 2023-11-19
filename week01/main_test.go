package main

import "testing"

func TestAdd(test *testing.T) {
	test_case := []struct {
		a    int
		b    int
		want int
	}{
		{1, 2, 3},
		{-10, -5, -15},
		{0, 0, 0},
	}
	for _, tc := range test_case {
		result := add(tc.a, tc.b)
		if result != tc.want {
			test.Errorf("adding %d and %d, expected %d, got %d", tc.a, tc.b, result, tc.want)
		}
	}
}
