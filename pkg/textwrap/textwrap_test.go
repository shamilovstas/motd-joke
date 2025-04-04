package textwrap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinLinesWrap(t *testing.T) {
	input := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean bibendum risus diam, eget faucibus nulla varius vitae.\nSed vel enim enim. Sed sapien dui, scelerisque at diam vehicula, rhoncus malesuada justo."

	minLinesWrapper := MinLines{}
	output := minLinesWrapper.Wrap(input, 40)
	expected := `Lorem ipsum dolor sit amet, consectetur 
adipiscing elit. Aenean bibendum risus 
diam, eget faucibus nulla varius vitae.
Sed vel enim enim. Sed sapien dui, 
scelerisque at diam vehicula, rhoncus 
malesuada justo.`

	assert.Equal(t, expected, output)

}
