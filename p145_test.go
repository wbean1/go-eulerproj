package main

import "testing"

type testcase struct {
	input    int
	expected bool
}

var oddTestCases = []testcase{
	{
		input:    1,
		expected: true,
	},
	{
		input:    2,
		expected: false,
	},
	{
		input:    3,
		expected: true,
	},
	{
		input:    10,
		expected: false,
	},
	{
		input:    11,
		expected: true,
	},
	{
		input:    13579,
		expected: true,
	},
	{
		input:    135279,
		expected: false,
	},
}

func TestIsAllOddDigits(t *testing.T) {
	for _, tc := range oddTestCases {
		if isAllOddDigits(tc.input) != tc.expected {
			t.Fatalf("FAIL: input: %d, expected: %t, got: %t\n", tc.input, tc.expected, isAllOddDigits(tc.input))
		}
	}
}
