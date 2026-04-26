package leetcode_daily

import (
	"strings"
)

// date: 2026-04-03
// https://leetcode.com/problems/decode-the-slanted-ciphertext/description
func decodeCiphertext(encodedText string, rows int) string {
	l := len(encodedText)

	//add := (rows * (rows - 1)) / 2
	cols := (l + rows - 1) / rows

	var result []byte
	for j := range cols {
		i, k := 0, j

		for i < rows && k < cols {
			idx := i*cols + k
			result = append(result, encodedText[idx])

			i++
			k++
		}
	}

	// Remove trailing spaces
	return strings.TrimRight(string(result), " ")
}
