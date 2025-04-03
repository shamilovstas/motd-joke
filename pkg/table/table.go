package table

import (
	"bytes"
	"fmt"
	"strings"
)

type Border struct {
	HorizontalSymbol rune
	VerticalSymbol   rune
}

func Draw(input string, b Border, pad int) string {
	lines := strings.Split(input, "\n")
	maxLength := maxLine(lines)

	var out bytes.Buffer

	leftBorder := fmt.Sprintf("%-*c", pad+1, b.VerticalSymbol)
	rightBorder := fmt.Sprintf("%*c", pad+1, b.VerticalSymbol)

	count := maxLength + pad*2 + 2 //max line length + padding from both sized + 2 symbols for vertical border
	horizontal := strings.Repeat(string(b.HorizontalSymbol), count)
	out.WriteString(horizontal)
	out.WriteRune('\n')
	for _, line := range lines {
		padded := fmt.Sprintf("%-*s", maxLength, line)
		out.WriteString(leftBorder)
		out.WriteString(padded)
		out.WriteString(rightBorder)
		out.WriteRune('\n')
	}
	out.WriteString(horizontal)
	out.WriteRune('\n')
	return out.String()
}

func maxLine(lines []string) int {
	maxLength := 0

	for _, line := range lines {
		l := len(line)
		if l > maxLength {
			maxLength = l
		}
	}
	return maxLength
}
