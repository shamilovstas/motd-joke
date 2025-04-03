package textwrap

import (
	"testing"
)

func TestMinLinesWrap(t *testing.T) {
	input := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean bibendum risus diam, eget faucibus nulla varius vitae. Sed vel enim enim. Sed sapien dui, scelerisque at diam vehicula, rhoncus malesuada justo."

	minLinesWrapper := MinLines{}
	output := minLinesWrapper.Wrap(input, 40)

	expected := `Lorem ipsum dolor sit amet, consectetur 
adipiscing elit. Aenean bibendum risus 
diam, eget faucibus nulla varius vitae. 
Sed vel enim enim. Sed sapien dui, 
scelerisque at diam vehicula, rhoncus 
malesuada justo.`

	if expected != output {
		t.Fatalf("output differs from expected")
	}
}
