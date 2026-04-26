package leetcode_daily

import (
	"fmt"
	"testing"
)

func Test_decodeCiphertext(t *testing.T) {
	testCase := []struct {
		encodedText string
		rows        int
		expected    string
	}{
		{"ch   ie   pr", 3, "cipher"},
		{"iveo    eed   l te   olc", 4, "i love leetcode"},
		{" b  ac", 2, "abc"},
	}

	for i, tc := range testCase {
		t.Run(fmt.Sprintf("test case: %d", i), func(t *testing.T) {
			result := decodeCiphertext(tc.encodedText, tc.rows)
			if result != tc.expected {
				t.Errorf("decodeCiphertext(%v, %v) = %v, expected %v", tc.encodedText, tc.rows, result, tc.expected)
			}
		})
	}
}
