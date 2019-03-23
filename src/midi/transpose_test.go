package midi

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranspose(t *testing.T) {
	var notes = []string{"C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B"}
	fnames, _ := filepath.Glob("phrases/phrase*.mid")
	for _, note := range notes {
		for _, fname := range fnames {
			if strings.Contains(fname, "tr") {
				continue
			}
			assert.Nil(t, Transpose(fname, note))
		}

	}
}
