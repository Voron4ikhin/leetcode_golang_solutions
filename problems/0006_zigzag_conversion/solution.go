package _006_zigzag_conversion

import "strings"

func Convert(s string, numRows int) string {
	if len(s) < numRows || numRows == 1 {
		return s
	}
	var builder strings.Builder
	matrix := make([][]rune, numRows)
	direction := 1
	matrixNumber := 0
	for i := 0; i < len(s); i++ {
		matrix[matrixNumber] = append(matrix[matrixNumber], rune(s[i]))
		matrixNumber += direction
		if matrixNumber == numRows-1 || matrixNumber == 0 {
			direction *= -1
		}
	}

	for i := 0; i < numRows; i++ {
		for j := 0; j < len(matrix[i]); j++ {
			builder.WriteRune(matrix[i][j])
		}
	}

	result := builder.String()

	return result
}
