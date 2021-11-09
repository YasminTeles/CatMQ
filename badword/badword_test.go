package badword

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBadWords(t *testing.T) {
	t.Parallel()

	badWords, err := NewBadWords()

	assert.NoError(t, err)
	assert.NotNil(t, badWords)
}

func TestGetBadWordFormatter(t *testing.T) {
	t.Parallel()

	content, err := getBadWordFormatter()

	assert.NoError(t, err)
	assert.Equal(t, "paralelepipedo|feira|feio", content)
}

func TestGetFileContent(t *testing.T) {
	t.Parallel()

	content, err := getFileContent()

	expected := `paralelepipedo
feira
feio
`

	assert.NoError(t, err)
	assert.Equal(t, expected, content)
}

func TestIsBadWord(t *testing.T) {
	t.Parallel()

	badWords, err := NewBadWords()

	assert.NoError(t, err)

	cases := []struct {
		in   string
		want bool
	}{
		{"O paralelepipedo foi pintado de verde e amarelo para copa.", true},
		{"O pato feio foi a feira.", true},
		{"Arturo y Lucho son mis patas.", false},
	}

	var isBadWord bool

	for _, c := range cases {
		isBadWord = badWords.Check(c.in)

		assert.Equal(t, c.want, isBadWord)
	}
}
