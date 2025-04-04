package table

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"unicode"
)

func dedent(data string) string {
	lines := strings.Split(data, "\n")

	trimFunc := func(r rune) bool {
		return unicode.IsSpace(r)
	}
	for i, line := range lines {
		lines[i] = strings.TrimLeftFunc(line, trimFunc)
	}

	return strings.Join(lines, "\n")
}
func TestDrawTable(t *testing.T) {
	input := "Testing"

	expected :=
		`---------
		|Testing|
		---------
		`

	actual := Draw(input, Border{VerticalSymbol: '|', HorizontalSymbol: '-'}, 0)
	assert.Equal(t, dedent(expected), actual)
}

func TestDrawPaddedTable(t *testing.T) {
	input := "Testing"

	expected :=
		`-------------
		|  Testing  |
		-------------
		`

	actual := Draw(input, Border{VerticalSymbol: '|', HorizontalSymbol: '-'}, 2)
	assert.Equal(t, dedent(expected), actual)
}
