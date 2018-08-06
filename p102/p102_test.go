package main

import "testing"

type triangleTestCases struct {
	input    string
	expected bool
}

var testCases = []triangleTestCases{
	{
		input:    "-340,495,-153,-910,835,-947",
		expected: true,
	},
	{
		input:    "-175,41,-421,-714,574,-645",
		expected: false,
	},
}

func TestP102(t *testing.T) {
	for _, testcase := range testCases {
		triangle := parseTriangle(testcase.input)
		origin := Coord{0, 0}
		x := triangleContains(triangle, origin)
		if x != testcase.expected {
			t.Fatalf("FAIL: %s, expected(%t), got(%t)",
				testcase.input, testcase.expected, x)
		}
	}
}
