package leetcode_daily

import (
	"testing"
)

func Test_maxWalls(t *testing.T) {
	// Example 1:
	robots := []int{10, 2}
	distance := []int{5, 1}
	walls := []int{5, 2, 7}
	expected := 3
	if got := maxWalls(robots, distance, walls); got != expected {
		t.Errorf("maxWalls() = %v, want %v", got, expected)
	}

	// Example 2:
	robots = []int{63, 56, 40, 45, 4, 9, 44, 69, 55, 26, 73, 15, 12, 60, 43, 39, 37, 74, 36, 34, 13, 23, 66, 14, 11, 42, 72, 3, 57, 10, 53, 8, 70, 17, 58, 61, 30, 32}
	distance = []int{8, 7, 4, 8, 9, 5, 2, 4, 5, 2, 6, 9, 5, 9, 5, 3, 7, 6, 9, 2, 8, 7, 4, 3, 5, 1, 7, 5, 1, 3, 5, 3, 5, 4, 8, 7, 6, 4}
	walls = []int{6, 22, 50, 52, 20, 9, 23, 75, 26, 21, 60, 58, 41, 28, 30}
	expected = 15
	if got := maxWalls(robots, distance, walls); got != expected {
		t.Errorf("maxWalls() = %v, want %v", got, expected)
	}
}
