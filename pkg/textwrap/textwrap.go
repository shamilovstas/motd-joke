package textwrap

import (
	"bytes"
)

type Wrapper interface {
	Wrap(src string, width int) string
}

type MinLines struct{}

func NewMinLines() *MinLines {
	return &MinLines{}
}

func (ml *MinLines) Wrap(src string, width int) string {
	spaceLeft := width
	wordWidth := 0
	var out bytes.Buffer
	for i, r := range src {
		if r == ' ' || i == len(src)-1 {
			if wordWidth+1 > spaceLeft {
				out.WriteRune('\n')
				out.WriteString(src[i-(wordWidth) : i])
				out.WriteRune(r)
				spaceLeft = width - wordWidth
			} else {
				out.WriteString(src[i-(wordWidth) : i])
				out.WriteRune(r)
				spaceLeft -= wordWidth + 1
			}
			wordWidth = 0
		} else {
			wordWidth++
		}
	}
	return out.String()
}
